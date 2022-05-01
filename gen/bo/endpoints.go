// Code generated by goa v3.7.3, DO NOT EDIT.
//
// bo endpoints
//
// Command:
// $ goa gen webcup/design

package bo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "bo" service endpoints.
type Endpoints struct {
	GetBoUsers        goa.Endpoint
	DeleteBoUser      goa.Endpoint
	DeleteBoManyUsers goa.Endpoint
	UpdateBoUser      goa.Endpoint
	GetBoUser         goa.Endpoint
}

// NewEndpoints wraps the methods of the "bo" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		GetBoUsers:        NewGetBoUsersEndpoint(s, a.OAuth2Auth, a.JWTAuth),
		DeleteBoUser:      NewDeleteBoUserEndpoint(s, a.OAuth2Auth, a.JWTAuth),
		DeleteBoManyUsers: NewDeleteBoManyUsersEndpoint(s, a.OAuth2Auth, a.JWTAuth),
		UpdateBoUser:      NewUpdateBoUserEndpoint(s, a.OAuth2Auth, a.JWTAuth),
		GetBoUser:         NewGetBoUserEndpoint(s, a.OAuth2Auth, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "bo" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetBoUsers = m(e.GetBoUsers)
	e.DeleteBoUser = m(e.DeleteBoUser)
	e.DeleteBoManyUsers = m(e.DeleteBoManyUsers)
	e.UpdateBoUser = m(e.UpdateBoUser)
	e.GetBoUser = m(e.GetBoUser)
}

// NewGetBoUsersEndpoint returns an endpoint function that calls the method
// "getBoUsers" of service "bo".
func NewGetBoUsersEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetBoUsersPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "OAuth2",
			Scopes:         []string{"api:read"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "client_credentials",
					TokenURL:   "/authorization",
					RefreshURL: "/refresh",
				},
			},
		}
		var token string
		if p.Oauth != nil {
			token = *p.Oauth
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWTToken != nil {
				token = *p.JWTToken
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.GetBoUsers(ctx, p)
	}
}

// NewDeleteBoUserEndpoint returns an endpoint function that calls the method
// "deleteBoUser" of service "bo".
func NewDeleteBoUserEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteBoUserPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "OAuth2",
			Scopes:         []string{"api:read"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "client_credentials",
					TokenURL:   "/authorization",
					RefreshURL: "/refresh",
				},
			},
		}
		var token string
		if p.Oauth != nil {
			token = *p.Oauth
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWTToken != nil {
				token = *p.JWTToken
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.DeleteBoUser(ctx, p)
	}
}

// NewDeleteBoManyUsersEndpoint returns an endpoint function that calls the
// method "deleteBoManyUsers" of service "bo".
func NewDeleteBoManyUsersEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteBoManyUsersPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "OAuth2",
			Scopes:         []string{"api:read"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "client_credentials",
					TokenURL:   "/authorization",
					RefreshURL: "/refresh",
				},
			},
		}
		var token string
		if p.Oauth != nil {
			token = *p.Oauth
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWTToken != nil {
				token = *p.JWTToken
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.DeleteBoManyUsers(ctx, p)
	}
}

// NewUpdateBoUserEndpoint returns an endpoint function that calls the method
// "updateBoUser" of service "bo".
func NewUpdateBoUserEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateBoUserPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "OAuth2",
			Scopes:         []string{"api:read"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "client_credentials",
					TokenURL:   "/authorization",
					RefreshURL: "/refresh",
				},
			},
		}
		var token string
		if p.Oauth != nil {
			token = *p.Oauth
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWTToken != nil {
				token = *p.JWTToken
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.UpdateBoUser(ctx, p)
	}
}

// NewGetBoUserEndpoint returns an endpoint function that calls the method
// "getBoUser" of service "bo".
func NewGetBoUserEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetBoUserPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "OAuth2",
			Scopes:         []string{"api:read"},
			RequiredScopes: []string{},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "client_credentials",
					TokenURL:   "/authorization",
					RefreshURL: "/refresh",
				},
			},
		}
		var token string
		if p.Oauth != nil {
			token = *p.Oauth
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err == nil {
			sc := security.JWTScheme{
				Name:           "jwt",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{},
			}
			var token string
			if p.JWTToken != nil {
				token = *p.JWTToken
			}
			ctx, err = authJWTFn(ctx, token, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.GetBoUser(ctx, p)
	}
}