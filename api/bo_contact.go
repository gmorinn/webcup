package api

import (
	"context"
	"fmt"
	"log"
	bocontact "webcup/gen/bo_contact"
	db "webcup/internal"
	"webcup/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"goa.design/goa/v3/security"
)

// boContact service example implementation.
// The example methods log the requests and return zero values.
type boContactsrvc struct {
	logger *log.Logger
	server *Server
}

// NewBoContact returns the boContact service implementation.
func NewBoContact(logger *log.Logger, server *Server) bocontact.Service {
	return &boContactsrvc{logger, server}
}

func (s *boContactsrvc) errorResponse(msg string, err error) *bocontact.UnknownError {
	return &bocontact.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "boContact" for
// the "OAuth2" security scheme.
func (s *boContactsrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// JWTAuth implements the authorization logic for service "boContact" for the
// "jwt" security scheme.
func (s *boContactsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	if scheme == nil {
		return ctx, ErrNullPayload
	}

	claims := make(jwt.MapClaims)
	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		b := ([]byte(s.server.Config.Security.Secret))
		return b, nil
	})
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["id"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["exp"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	if claims["role"] != "admin" {
		return ctx, ErrInvalidTokenScopes
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		return ctx, ErrInvalidTokenScopes
	}

	// 3. add authInfo to context
	ctx = contextWithAuthInfo(ctx, authInfo{
		jwtToken: claims,
	})
	return ctx, nil
}

// Get All messages
func (s *boContactsrvc) GetBoContact(ctx context.Context, p *bocontact.GetBoContactPayload) (res *bocontact.GetBoContactResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.GetMessagesFilteredParams{
			Limit:     p.Limit,
			Offset:    p.Offset,
			EmailAsc:  utils.FilterOrderBy(p.Field, p.Direction, "EmailAsc"),
			EmailDesc: utils.FilterOrderBy(p.Field, p.Direction, "EmailDesc"),
		}
		uS, err := q.GetMessagesFiltered(ctx, arg)
		if err != nil {
			return fmt.Errorf("error get all messages: %v", err)
		}
		var allMessages []*bocontact.ResContact
		for _, v := range uS {
			allMessages = append(allMessages, &bocontact.ResContact{
				ID:      v.ID.String(),
				Email:   v.Email,
				Message: v.Msg,
				UserID:  v.UserID.String(),
			})
		}
		total, err := q.CountMessage(ctx)
		if err != nil {
			return fmt.Errorf("error count message %v", err)
		}
		res = &bocontact.GetBoContactResult{
			Contacts: allMessages,
			Count:    total,
			Success:  true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_ALL_MESSAGE", err)
	}
	return res, nil
}

// Delete one contact by ID
func (s *boContactsrvc) DeleteBoContact(ctx context.Context, p *bocontact.DeleteBoContactPayload) (res *bocontact.DeleteBoContactResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if err := q.DeleteMessageByID(ctx, uuid.MustParse(p.ID)); err != nil {
			return fmt.Errorf("error delete message: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_DELETE_MESSAGE", err)
	}
	return &bocontact.DeleteBoContactResult{Success: true}, nil
}

// get one contact by ID
func (s *boContactsrvc) GetBoContactByID(ctx context.Context, p *bocontact.GetBoContactByIDPayload) (res *bocontact.GetBoContactByIDResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		u, err := q.GetMessageByID(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("error get message by id	: %v", err)
		}
		res = &bocontact.GetBoContactByIDResult{
			Contact: &bocontact.ResContact{
				ID:      u.ID.String(),
				Email:   u.Email,
				Message: u.Msg,
				UserID:  u.UserID.String(),
			},
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_CONTACT_BY_ID", err)
	}
	return res, nil
}

// Delete many contact with IDs send in body
func (s *boContactsrvc) DeleteBoManyContact(ctx context.Context, p *bocontact.DeleteBoManyContactPayload) (res *bocontact.DeleteBoManyContactResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		for _, v := range p.Tab {
			_, err = uuid.Parse(v)
			if err != nil {
				return fmt.Errorf("error parse id: %v %v", v, err)
			}
			if err := q.DeleteMessageByID(ctx, uuid.MustParse(v)); err != nil {
				return fmt.Errorf("error delete message: %v %v", v, err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_DELETE_MESSAGE", err)
	}
	return &bocontact.DeleteBoManyContactResult{Success: true}, nil

}
