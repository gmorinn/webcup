package api

import (
	"context"
	"fmt"
	"log"
	publicusers "webcup/back/gen/public_users"
	db "webcup/back/internal"
	"webcup/back/utils"

	"goa.design/goa/v3/security"
)

// publicUsers service example implementation.
// The example methods log the requests and return zero values.
type publicUserssrvc struct {
	logger *log.Logger
	server *Server
}

// NewPublicUsers returns the publicUsers service implementation.
func NewPublicUsers(logger *log.Logger, server *Server) publicusers.Service {
	return &publicUserssrvc{logger, server}
}

func (s *publicUserssrvc) errorResponse(msg string, err error) *publicusers.UnknownError {
	return &publicusers.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "publicUsers" for
// the "OAuth2" security scheme.
func (s *publicUserssrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// List user for search bar
func (s *publicUserssrvc) ListUsers(ctx context.Context, p *publicusers.ListUsersPayload) (res *publicusers.ListUsersResult, err error) {
	if p == nil {
		return nil, ErrNullPayload
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		var result []*publicusers.ResUser
		users, err := q.ListUsers(ctx, utils.NullS("%"+p.Key+"%"))
		if err != nil {
			return fmt.Errorf("error list users by key: %v", err)
		}
		for _, v := range users {
			result = append(result, &publicusers.ResUser{
				ID:        v.ID.String(),
				Firstname: v.Firstname.String,
				Lastname:  v.Lastname.String,
				Username:  v.Username,
				Avatar:    v.Avatar.String,
				Role:      string(v.Role),
				Email:     v.Email,
			})
		}
		res = &publicusers.ListUsersResult{
			Success: true,
			Users:   result,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_LIST_USERS", err)
	}
	return res, nil
}

// List users the most recent
func (s *publicUserssrvc) ListUsersMostRecent(ctx context.Context, p *publicusers.ListUsersMostRecentPayload) (res *publicusers.ListUsersMostRecentResult, err error) {
	if p == nil {
		return nil, ErrNullPayload
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		var result []*publicusers.ResUser
		arg := db.ListUsersMostRecentParams{
			Limit:  p.Limit,
			Offset: p.Offset,
		}
		users, err := q.ListUsersMostRecent(ctx, arg)
		if err != nil {
			return fmt.Errorf("error list recent users: %v", err)
		}
		for _, v := range users {
			result = append(result, &publicusers.ResUser{
				ID:        v.ID.String(),
				Firstname: v.Firstname.String,
				Lastname:  v.Lastname.String,
				Username:  v.Username,
				Email:     v.Email,
				Avatar:    v.Avatar.String,
				Role:      string(v.Role),
			})
		}
		total, err := q.CountUser(ctx)
		if err != nil {
			return fmt.Errorf("error count users: %v", err)
		}
		res = &publicusers.ListUsersMostRecentResult{
			Users:   result,
			Success: true,
			Count:   total,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_ALL_USERS", err)
	}
	return res, nil
}

// Get one User by username
func (s *publicUserssrvc) GetUserByUsername(ctx context.Context, p *publicusers.GetUserByUsernamePayload) (res *publicusers.GetUserByUsernameResult, err error) {
	if p == nil {
		return nil, ErrNullPayload
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		u, err := q.GetUserByUsername(ctx, p.Username)
		if err != nil {
			return fmt.Errorf("error get user by username: %v", err)
		}
		res = &publicusers.GetUserByUsernameResult{
			User: &publicusers.ResUser{
				ID:        u.ID.String(),
				Firstname: u.Firstname.String,
				Lastname:  u.Lastname.String,
				Email:     u.Email,
				Username:  u.Username,
				Avatar:    u.Avatar.String,
				Role:      string(u.Role),
			},
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_USER_BY_USERNAME", err)
	}
	return res, nil
}
