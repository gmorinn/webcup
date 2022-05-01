package api

import (
	"context"
	"fmt"
	"log"
	contacts "webcup/gen/contacts"
	db "webcup/internal"

	"github.com/google/uuid"
	"goa.design/goa/v3/security"
)

// contacts service example implementation.
// The example methods log the requests and return zero values.
type contactssrvc struct {
	logger *log.Logger
	server *Server
}

// NewContacts returns the contacts service implementation.
func NewContacts(logger *log.Logger, server *Server) contacts.Service {
	return &contactssrvc{logger, server}
}

func (s *contactssrvc) errorResponse(msg string, err error) *contacts.UnknownError {
	return &contacts.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "contacts" for the
// "OAuth2" security scheme.
func (s *contactssrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// JWTAuth implements the authorization logic for service "contacts" for the
// "jwt" security scheme.
func (s *contactssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return s.server.CheckJWT(ctx, token, scheme)
}

// user ask for something
func (s *contactssrvc) AddMessage(ctx context.Context, p *contacts.AddMessagePayload) (res *contacts.AddMessageResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		email, err := q.GetEmailByUserID(ctx, uuid.MustParse(p.UserID))
		if err != nil {
			return fmt.Errorf("error get info user: %v", err)
		}
		arg := db.InsertMessageParams{
			Msg:    p.Msg,
			UserID: uuid.MustParse(p.UserID),
			Email:  email,
		}
		if err := q.InsertMessage(ctx, arg); err != nil {
			return fmt.Errorf("error insert message: %v", err)
		}
		res = &contacts.AddMessageResult{
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_CREATE_MESSAGE", err)
	}
	return res, nil
}
