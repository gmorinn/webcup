package api

import (
	"context"
	"fmt"
	"log"
	"strings"
	users "webcup/back/gen/users"
	db "webcup/back/internal"
	"webcup/back/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"goa.design/goa/v3/security"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	logger *log.Logger
	server *Server
}

type RealisationArray struct {
	Img  string `json:"img"`
	Text string `json:"text"`
}

type LinkArray struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}

type CompetenceArray struct {
	Value string `json:"value"`
}

// newUsers returns the users service implementation.
func NewUsers(logger *log.Logger, server *Server) users.Service {
	return &userssrvc{logger, server}
}

func (s *userssrvc) errorResponse(msg string, err error) *users.UnknownError {
	return &users.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "users" for the
// "OAuth2" security scheme.
func (s *userssrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// JWTAuth implements the authorization logic for service "users" for the "jwt"
// security scheme.
func (s *userssrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return s.server.CheckJWT(ctx, token, scheme)
}

// Delete one User by ID
func (s *userssrvc) DeleteUser(ctx context.Context, p *users.DeleteUserPayload) (res *users.DeleteUserResult, err error) {
	if p == nil {
		return nil, ErrNullPayload
	}
	userID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, ErrWrongIdFormat
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		isExist, _ := q.CheckIDExist(ctx, userID)
		if err != nil || !isExist {
			return ErrUserNotExist
		}
		if p.JWTToken == nil {
			return ErrInvalidToken
		}
		claims := make(jwt.MapClaims)
		// authorize request
		// 1. parse JWT token, token key is hardcoded to "secret" in this example
		_, err := jwt.ParseWithClaims(*p.JWTToken, claims, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(s.server.Config.Security.Secret))
			return b, nil
		})
		if err != nil || claims["id"] == nil {
			return ErrInvalidToken
		}
		if claims["id"].(string) != p.ID {
			return ErrInvalidRequest
		}
		if err := q.DeleteUserByID(ctx, userID); err != nil {
			return fmt.Errorf("ERROR_DELETE_USER_BY_ID %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_DELETE_USER", err)
	}
	return &users.DeleteUserResult{Success: true}, nil
}

// Get one User
func (s *userssrvc) GetUserByID(ctx context.Context, p *users.GetUserByIDPayload) (res *users.GetUserByIDResult, err error) {
	if p == nil {
		return nil, ErrNullPayload
	}
	userID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, ErrWrongIdFormat
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		u, err := q.GetUserByID(ctx, userID)
		if err != nil {
			return fmt.Errorf("error get user by id: %v", err)
		}
		res = &users.GetUserByIDResult{
			User: &users.ResUser{
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
		return nil, s.errorResponse("TX_GET_USER_BY_ID", err)
	}
	return res, nil
}

// Update main info like email, username, firstname, lastname
func (s *userssrvc) UpdateDescription(ctx context.Context, p *users.UpdateDescriptionPayload) (res *users.UpdateDescriptionResult, err error) {
	if p == nil {
		return nil, ErrNullPayload
	}
	userID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, ErrWrongIdFormat
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.UpdateDescriptionUserParams{
			ID:        userID,
			Firstname: utils.NullS(*p.Firstname),
			Lastname:  utils.NullS(*p.Lastname),
			Username:  strings.ToLower(p.Username),
			Email:     p.Email,
		}
		if err := q.UpdateDescriptionUser(ctx, arg); err != nil {
			return fmt.Errorf("ERROR_UPDATE_USER_DESCRIPTION %v", err)
		}
		newUser, err := q.GetUserByID(ctx, userID)
		if err != nil {
			return fmt.Errorf("error get user by id: %v", err)
		}
		res = &users.UpdateDescriptionResult{
			Success: true,
			User: &users.ResUser{
				ID:        newUser.ID.String(),
				Firstname: newUser.Firstname.String,
				Lastname:  newUser.Lastname.String,
				Username:  newUser.Username,
				Email:     newUser.Email,
				Avatar:    newUser.Avatar.String,
				Role:      string(newUser.Role),
			},
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_UPDATE_USER_DESCRIPTION", err)
	}
	return res, nil
}
