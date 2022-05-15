// Code generated by goa v3.7.3, DO NOT EDIT.
//
// publicUsers HTTP server
//
// Command:
// $ goa gen webcup/design

package server

import (
	"context"
	"net/http"
	"regexp"
	publicusers "webcup/gen/public_users"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the publicUsers service endpoint HTTP handlers.
type Server struct {
	Mounts              []*MountPoint
	GetUserByUsername   http.Handler
	ListUsers           http.Handler
	ListUsersMostRecent http.Handler
	CORS                http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the publicUsers service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *publicusers.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"GetUserByUsername", "GET", "/v1/web/public/user/{username}"},
			{"ListUsers", "PATCH", "/v1/web/public/user/search"},
			{"ListUsersMostRecent", "GET", "/v1/web/public/user/recents/{offset}/{limit}"},
			{"CORS", "OPTIONS", "/v1/web/public/user/{username}"},
			{"CORS", "OPTIONS", "/v1/web/public/user/search"},
			{"CORS", "OPTIONS", "/v1/web/public/user/recents/{offset}/{limit}"},
		},
		GetUserByUsername:   NewGetUserByUsernameHandler(e.GetUserByUsername, mux, decoder, encoder, errhandler, formatter),
		ListUsers:           NewListUsersHandler(e.ListUsers, mux, decoder, encoder, errhandler, formatter),
		ListUsersMostRecent: NewListUsersMostRecentHandler(e.ListUsersMostRecent, mux, decoder, encoder, errhandler, formatter),
		CORS:                NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "publicUsers" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.GetUserByUsername = m(s.GetUserByUsername)
	s.ListUsers = m(s.ListUsers)
	s.ListUsersMostRecent = m(s.ListUsersMostRecent)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the publicUsers endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetUserByUsernameHandler(mux, h.GetUserByUsername)
	MountListUsersHandler(mux, h.ListUsers)
	MountListUsersMostRecentHandler(mux, h.ListUsersMostRecent)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the publicUsers endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountGetUserByUsernameHandler configures the mux to serve the "publicUsers"
// service "getUserByUsername" endpoint.
func MountGetUserByUsernameHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePublicUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/web/public/user/{username}", f)
}

// NewGetUserByUsernameHandler creates a HTTP handler which loads the HTTP
// request and calls the "publicUsers" service "getUserByUsername" endpoint.
func NewGetUserByUsernameHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetUserByUsernameRequest(mux, decoder)
		encodeResponse = EncodeGetUserByUsernameResponse(encoder)
		encodeError    = EncodeGetUserByUsernameError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getUserByUsername")
		ctx = context.WithValue(ctx, goa.ServiceKey, "publicUsers")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListUsersHandler configures the mux to serve the "publicUsers" service
// "listUsers" endpoint.
func MountListUsersHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePublicUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PATCH", "/v1/web/public/user/search", f)
}

// NewListUsersHandler creates a HTTP handler which loads the HTTP request and
// calls the "publicUsers" service "listUsers" endpoint.
func NewListUsersHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListUsersRequest(mux, decoder)
		encodeResponse = EncodeListUsersResponse(encoder)
		encodeError    = EncodeListUsersError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "listUsers")
		ctx = context.WithValue(ctx, goa.ServiceKey, "publicUsers")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListUsersMostRecentHandler configures the mux to serve the
// "publicUsers" service "listUsersMostRecent" endpoint.
func MountListUsersMostRecentHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePublicUsersOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/web/public/user/recents/{offset}/{limit}", f)
}

// NewListUsersMostRecentHandler creates a HTTP handler which loads the HTTP
// request and calls the "publicUsers" service "listUsersMostRecent" endpoint.
func NewListUsersMostRecentHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListUsersMostRecentRequest(mux, decoder)
		encodeResponse = EncodeListUsersMostRecentResponse(encoder)
		encodeError    = EncodeListUsersMostRecentError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "listUsersMostRecent")
		ctx = context.WithValue(ctx, goa.ServiceKey, "publicUsers")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service publicUsers.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandlePublicUsersOrigin(h)
	mux.Handle("OPTIONS", "/v1/web/public/user/{username}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/web/public/user/search", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/web/public/user/recents/{offset}/{limit}", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandlePublicUsersOrigin applies the CORS response headers corresponding to
// the origin for the service publicUsers.
func HandlePublicUsersOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile(".*localhost.*")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Content-Type, Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE, PATCH")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, jwtToken, Origin")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
