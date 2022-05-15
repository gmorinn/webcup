package api

import (
	"context"
	"fmt"
	"log"
	data "webcup/gen/data"
	db "webcup/internal"
	"webcup/utils"

	"github.com/google/uuid"
	"goa.design/goa/v3/security"
)

// data service example implementation.
// The example methods log the requests and return zero values.
type datasrvc struct {
	logger *log.Logger
	server *Server
}

// NewData returns the data service implementation.
func NewData(logger *log.Logger, server *Server) data.Service {
	return &datasrvc{logger, server}
}

func (s *datasrvc) errorResponse(msg string, err error) *data.UnknownError {
	return &data.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "data" for the
// "OAuth2" security scheme.
func (s *datasrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// JWTAuth implements the authorization logic for service "data" for the "jwt"
// security scheme.
func (s *datasrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return s.server.CheckJWT(ctx, token, scheme)
}

func (s *datasrvc) ListData(ctx context.Context, p *data.ListDataPayload) (res *data.ListDataResult, err error) {
	if p == nil {
		return nil, ErrNullPayload
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		var result []*data.ResData
		datas, err := q.ListData(ctx, "%"+p.Key+"%")
		if err != nil {
			return fmt.Errorf("error list datas by key: %v", err)
		}
		for _, v := range datas {
			result = append(result, &data.ResData{
				ID:          v.ID.String(),
				Title:       v.Title,
				Description: v.Description,
				Image:       v.Img.String,
				Category:    string(v.Category),
				UserID:      v.UserID.String(),
			})
		}
		res = &data.ListDataResult{
			Success: true,
			Data:    result,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_LIST_DATA", err)
	}
	return res, nil
}

// Create one data
func (s *datasrvc) CreateData(ctx context.Context, p *data.CreateDataPayload) (res *data.CreateDataResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		stock, err := q.GetStockByUserID(ctx, uuid.MustParse(p.Data.UserID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_STOCK_BY_ID %v", err)
		}
		count, err := q.CountDataByUserID(ctx, uuid.MustParse(p.Data.UserID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_STOCK_BY_ID %v", err)
		}
		if count >= stock || stock == 0 {
			return fmt.Errorf("ERROR_CREATE_DATA_MAX_DATA")
		}
		arg := db.CreateDataParams{
			Title:       p.Data.Title,
			Description: p.Data.Description,
			Img:         utils.NullS(*p.Data.Image),
			Category:    db.Futur(p.Data.Category),
			UserID:      uuid.MustParse(p.Data.UserID),
		}
		createData, err := q.CreateData(ctx, arg)
		if err != nil {
			return fmt.Errorf("ERROR_CREATE_DATA %v", err)
		}
		NewData, err := q.GetDataByID(ctx, createData.ID)
		if err != nil {
			return fmt.Errorf("ERROR_GET_DATA_BY_ID %v", err)
		}
		arg2 := db.UpdateStockParams{
			ID:    NewData.UserID,
			Stock: stock - 1,
		}
		if err := q.UpdateStock(ctx, arg2); err != nil {
			return fmt.Errorf("ERROR_UPDATE_USER_DESCRIPTION %v", err)
		}
		res = &data.CreateDataResult{
			Data: &data.ResData{
				ID:          NewData.ID.String(),
				Title:       NewData.Title,
				Description: NewData.Description,
				Image:       NewData.Img.String,
				Category:    string(NewData.Category),
				UserID:      NewData.UserID.String()},
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_CREATE_DATA", err)
	}
	return res, nil
}

// Update one data
func (s *datasrvc) UpdateData(ctx context.Context, p *data.UpdateDataPayload) (res *data.UpdateDataResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.UpdateDataParams{
			ID:          uuid.MustParse(p.ID),
			Title:       p.Data.Title,
			Description: p.Data.Description,
			Img:         utils.NullS(*p.Data.Image),
			Category:    db.Futur(p.Data.Category),
		}
		if err := q.UpdateData(ctx, arg); err != nil {
			return fmt.Errorf("ERROR_UPDATE_DATA %v", err)
		}
		newData, err := q.GetDataByID(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_DATA_BY_ID %v", err)
		}
		res = &data.UpdateDataResult{
			Success: true,
			Data: &data.ResData{
				ID:          newData.ID.String(),
				Title:       newData.Title,
				Description: newData.Description,
				Image:       newData.Img.String,
				Category:    string(newData.Category),
				UserID:      newData.UserID.String()},
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_UPDATE_DATA", err)
	}
	return res, nil
}

// List data the most recent
func (s *datasrvc) ListDataMostRecent(ctx context.Context, p *data.ListDataMostRecentPayload) (res *data.ListDataMostRecentResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		var result []*data.ResData
		arg := db.ListDataMostRecentParams{
			Limit:  p.Limit,
			Offset: p.Offset,
		}
		ds, err := q.ListDataMostRecent(ctx, arg)
		if err != nil {
			return fmt.Errorf("error list recent data: %v", err)
		}
		for _, v := range ds {
			result = append(result, &data.ResData{
				ID:          v.ID.String(),
				Title:       v.Title,
				Description: v.Description,
				Image:       v.Img.String,
				Category:    string(v.Category),
				UserID:      v.UserID.String(),
			})
		}
		total, err := q.CountData(ctx)
		if err != nil {
			return fmt.Errorf("error count data: %v", err)
		}
		res = &data.ListDataMostRecentResult{
			Data:    result,
			Success: true,
			Count:   total,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_ALL_DATA_RECENT", err)
	}
	return res, nil
}

// Get one data user id
func (s *datasrvc) GetDataByUserID(ctx context.Context, p *data.GetDataByUserIDPayload) (res *data.GetDataByUserIDResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.GetDataByUserIDParams{
			Limit:  p.Limit,
			Offset: p.Offset,
			UserID: uuid.MustParse(p.UserID),
		}
		tab, err := q.GetDataByUserID(ctx, arg)
		if err != nil {
			return fmt.Errorf("ERROR_GET_ALL_DATA_BY_USER_ID %v", err)
		}
		var ds []*data.ResData
		for _, v := range tab {
			ds = append(ds, &data.ResData{
				ID:          v.ID.String(),
				Title:       v.Title,
				Description: v.Description,
				Image:       v.Img.String,
				Category:    string(v.Category),
				UserID:      v.UserID.String(),
			})
		}
		total, err := q.CountData(ctx)
		if err != nil {
			return fmt.Errorf("error count data: %v", err)
		}
		res = &data.GetDataByUserIDResult{
			Data:    ds,
			Count:   total,
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_ALL_DATA_BY_USER_ID", err)
	}
	return res, nil
}

// Get one data by id
func (s *datasrvc) GetDataByID(ctx context.Context, p *data.GetDataByIDPayload) (res *data.GetDataByIDResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		d, err := q.GetDataByID(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_DATA_BY_ID %v", err)
		}
		res = &data.GetDataByIDResult{
			Data: &data.ResData{
				ID:          d.ID.String(),
				Title:       d.Title,
				Description: d.Description,
				Image:       d.Img.String,
				Category:    string(d.Category),
				UserID:      d.UserID.String(),
			},
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_DATA_BY_ID", err)
	}
	return res, nil
}
