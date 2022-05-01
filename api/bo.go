package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	bo "webcup/gen/bo"
	db "webcup/internal"
	"webcup/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"goa.design/goa/v3/security"
)

// bo service example implementation.
// The example methods log the requests and return zero values.
type bosrvc struct {
	logger *log.Logger
	server *Server
}

// NewBo returns the bo service implementation.
func NewBo(logger *log.Logger, server *Server) bo.Service {
	return &bosrvc{logger, server}
}

func (s *bosrvc) errorResponse(msg string, err error) *bo.UnknownError {
	return &bo.UnknownError{
		Err:       err.Error(),
		ErrorCode: msg,
	}
}

// OAuth2Auth implements the authorization logic for service "bo" for the
// "OAuth2" security scheme.
func (s *bosrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	return s.server.CheckAuth(ctx, token, scheme)
}

// JWTAuth implements the authorization logic for service "bo" for the "jwt"
// security scheme.
func (s *bosrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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

// Get All users
func (s *bosrvc) GetBoUsers(ctx context.Context, p *bo.GetBoUsersPayload) (res *bo.GetBoUsersResult, err error) {
	if p == nil {
		return nil, s.errorResponse("ERROR_UPDATE_USER_PAYLOAD_NIL", ErrNullPayload)
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.GetBoAllUsersParams{
			Limit:         p.Limit,
			Offset:        p.Offset,
			FirstnameAsc:  utils.FilterOrderBy(p.Field, p.Direction, "FirstnameAsc"),
			FirstnameDesc: utils.FilterOrderBy(p.Field, p.Direction, "FirstnameDesc"),
			EmailAsc:      utils.FilterOrderBy(p.Field, p.Direction, "EmailAsc"),
			EmailDesc:     utils.FilterOrderBy(p.Field, p.Direction, "EmailDesc"),
			LastnameAsc:   utils.FilterOrderBy(p.Field, p.Direction, "LastnameAsc"),
			LastnameDesc:  utils.FilterOrderBy(p.Field, p.Direction, "LastnameDesc"),
			UsernameAsc:   utils.FilterOrderBy(p.Field, p.Direction, "UsernameAsc"),
			UsernameDesc:  utils.FilterOrderBy(p.Field, p.Direction, "UsernameDesc"),
			RoleAsc:       utils.FilterOrderBy(p.Field, p.Direction, "RoleAsc"),
			RoleDesc:      utils.FilterOrderBy(p.Field, p.Direction, "RoleDesc"),
		}
		uS, err := q.GetBoAllUsers(ctx, arg)
		if err != nil {
			return fmt.Errorf("ERROR_GET_ALL_USERS %v", err)
		}
		var allUsers []*bo.ResUser
		for _, v := range uS {
			allUsers = append(allUsers, &bo.ResUser{
				ID:        v.ID.String(),
				Firstname: v.Firstname.String,
				Lastname:  v.Lastname.String,
				Email:     v.Email,
				Username:  v.Username,
				Avatar:    v.Avatar.String,
				Role:      string(v.Role),
			})
		}
		total, err := q.CountUser(ctx)
		if err != nil {
			return fmt.Errorf("ERROR_COUNT_USERS %v", err)
		}
		res = &bo.GetBoUsersResult{
			Users:   allUsers,
			Count:   total,
			Success: true,
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_GET_ALL_USERS", err)
	}
	return res, nil
}

// Delete one User by ID
func (s *bosrvc) DeleteBoUser(ctx context.Context, p *bo.DeleteBoUserPayload) (res *bo.DeleteBoUserResult, err error) {
	if p == nil {
		return nil, s.errorResponse("ERROR_UPDATE_USER_PAYLOAD_NIL", ErrNullPayload)
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		if err := q.DeleteUserByID(ctx, uuid.MustParse(p.ID)); err != nil {
			return fmt.Errorf("ERROR_DELETE_USER_BY_ID %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_DELETE_USER", err)
	}
	return &bo.DeleteBoUserResult{Success: true}, nil
}

// Delete many users with IDs send in body
func (s *bosrvc) DeleteBoManyUsers(ctx context.Context, p *bo.DeleteBoManyUsersPayload) (res *bo.DeleteBoManyUsersResult, err error) {
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		for _, v := range p.Tab {
			_, err = uuid.Parse(v)
			if err != nil {
				return fmt.Errorf("ERROR_DELETE_USER_BY_ID_%v %v", v, err)
			}
			if err := q.DeleteUserByID(ctx, uuid.MustParse(v)); err != nil {
				return fmt.Errorf("ERROR_DELETE_USER_BY_ID_%v %v", v, err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_DELETE_USERS", err)
	}
	return &bo.DeleteBoManyUsersResult{Success: true}, nil
}

// Update one User
func (s *bosrvc) UpdateBoUser(ctx context.Context, p *bo.UpdateBoUserPayload) (res *bo.UpdateBoUserResult, err error) {
	if p == nil {
		return nil, s.errorResponse("ERROR_UPDATE_USER_PAYLOAD_NIL", ErrNullPayload)
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.UpdateUserBoParams{
			ID:        uuid.MustParse(p.ID),
			Firstname: utils.NullS(p.User.Firstname),
			Lastname:  utils.NullS(p.User.Lastname),
			Email:     p.User.Email,
			Avatar:    utils.NullS(p.User.Avatar),
			Username:  p.User.Username,
			Role:      db.Role(p.User.Role),
		}
		if err := q.UpdateUserBo(ctx, arg); err != nil {
			return fmt.Errorf("ERROR_UPDATE_USER %v", err)
		}
		newUser, err := q.GetUserByID(ctx, uuid.MustParse(p.ID))
		if err != nil {
			return fmt.Errorf("ERROR_GET_USER_BY_ID %v", err)
		}
		res = &bo.UpdateBoUserResult{
			Success: true,
			User: &bo.ResUser{
				ID:        newUser.ID.String(),
				Firstname: newUser.Firstname.String,
				Lastname:  newUser.Lastname.String,
				Email:     newUser.Email,
				Username:  newUser.Username,
				Avatar:    newUser.Avatar.String,
				Role:      string(newUser.Role),
			},
		}
		return nil
	})
	if err != nil {
		return nil, s.errorResponse("TX_UPDATE_USER", err)
	}
	return res, nil
}

// Get one User
func (s *bosrvc) GetBoUser(ctx context.Context, p *bo.GetBoUserPayload) (res *bo.GetBoUserResult, err error) {
	if p == nil {
		return nil, s.errorResponse("ERROR_UPDATE_USER_PAYLOAD_NIL", ErrNullPayload)
	}
	userID, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, s.errorResponse("ERROR_UPDATE_USER_WRONG_ID", errors.New("wrong id"))
	}
	err = s.server.Store.ExecTx(ctx, func(q *db.Queries) error {
		u, err := q.GetUserByID(ctx, userID)
		if err != nil {
			return fmt.Errorf("ERROR_GET_USER_BY_ID %v", err)
		}
		res = &bo.GetBoUserResult{
			User: &bo.ResUser{
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
