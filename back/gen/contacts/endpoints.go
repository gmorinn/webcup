// Code generated by goa v3.7.2, DO NOT EDIT.
//
// contacts endpoints
//
// Command:
// $ goa gen webcup/back/design

package contacts

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "contacts" service endpoints.
type Endpoints struct {
	AddMessage goa.Endpoint
}

// NewEndpoints wraps the methods of the "contacts" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		AddMessage: NewAddMessageEndpoint(s, a.OAuth2Auth, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "contacts" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.AddMessage = m(e.AddMessage)
}

// NewAddMessageEndpoint returns an endpoint function that calls the method
// "addMessage" of service "contacts".
func NewAddMessageEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddMessagePayload)
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
		return s.AddMessage(ctx, p)
	}
}
