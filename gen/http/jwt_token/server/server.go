// Code generated by goa v3.7.3, DO NOT EDIT.
//
// jwtToken HTTP server
//
// Command:
// $ goa gen webcup/design

package server

import (
	"context"
	"net/http"
	"regexp"
	jwttoken "webcup/gen/jwt_token"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the jwtToken service endpoint HTTP handlers.
type Server struct {
	Mounts   []*MountPoint
	Signup   http.Handler
	Signin   http.Handler
	Refresh  http.Handler
	SigninBo http.Handler
	CORS     http.Handler
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

// New instantiates HTTP handlers for all the jwtToken service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *jwttoken.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Signup", "POST", "/signup"},
			{"Signin", "POST", "/signin"},
			{"Refresh", "POST", "/resfresh"},
			{"SigninBo", "POST", "/bo/signin"},
			{"CORS", "OPTIONS", "/signup"},
			{"CORS", "OPTIONS", "/signin"},
			{"CORS", "OPTIONS", "/resfresh"},
			{"CORS", "OPTIONS", "/bo/signin"},
		},
		Signup:   NewSignupHandler(e.Signup, mux, decoder, encoder, errhandler, formatter),
		Signin:   NewSigninHandler(e.Signin, mux, decoder, encoder, errhandler, formatter),
		Refresh:  NewRefreshHandler(e.Refresh, mux, decoder, encoder, errhandler, formatter),
		SigninBo: NewSigninBoHandler(e.SigninBo, mux, decoder, encoder, errhandler, formatter),
		CORS:     NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "jwtToken" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Signup = m(s.Signup)
	s.Signin = m(s.Signin)
	s.Refresh = m(s.Refresh)
	s.SigninBo = m(s.SigninBo)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the jwtToken endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountSignupHandler(mux, h.Signup)
	MountSigninHandler(mux, h.Signin)
	MountRefreshHandler(mux, h.Refresh)
	MountSigninBoHandler(mux, h.SigninBo)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the jwtToken endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountSignupHandler configures the mux to serve the "jwtToken" service
// "signup" endpoint.
func MountSignupHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleJWTTokenOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/signup", f)
}

// NewSignupHandler creates a HTTP handler which loads the HTTP request and
// calls the "jwtToken" service "signup" endpoint.
func NewSignupHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSignupRequest(mux, decoder)
		encodeResponse = EncodeSignupResponse(encoder)
		encodeError    = EncodeSignupError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "signup")
		ctx = context.WithValue(ctx, goa.ServiceKey, "jwtToken")
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

// MountSigninHandler configures the mux to serve the "jwtToken" service
// "signin" endpoint.
func MountSigninHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleJWTTokenOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/signin", f)
}

// NewSigninHandler creates a HTTP handler which loads the HTTP request and
// calls the "jwtToken" service "signin" endpoint.
func NewSigninHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSigninRequest(mux, decoder)
		encodeResponse = EncodeSigninResponse(encoder)
		encodeError    = EncodeSigninError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "signin")
		ctx = context.WithValue(ctx, goa.ServiceKey, "jwtToken")
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

// MountRefreshHandler configures the mux to serve the "jwtToken" service
// "refresh" endpoint.
func MountRefreshHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleJWTTokenOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/resfresh", f)
}

// NewRefreshHandler creates a HTTP handler which loads the HTTP request and
// calls the "jwtToken" service "refresh" endpoint.
func NewRefreshHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRefreshRequest(mux, decoder)
		encodeResponse = EncodeRefreshResponse(encoder)
		encodeError    = EncodeRefreshError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "refresh")
		ctx = context.WithValue(ctx, goa.ServiceKey, "jwtToken")
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

// MountSigninBoHandler configures the mux to serve the "jwtToken" service
// "signinBo" endpoint.
func MountSigninBoHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleJWTTokenOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/bo/signin", f)
}

// NewSigninBoHandler creates a HTTP handler which loads the HTTP request and
// calls the "jwtToken" service "signinBo" endpoint.
func NewSigninBoHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSigninBoRequest(mux, decoder)
		encodeResponse = EncodeSigninBoResponse(encoder)
		encodeError    = EncodeSigninBoError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "signinBo")
		ctx = context.WithValue(ctx, goa.ServiceKey, "jwtToken")
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
// service jwtToken.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleJWTTokenOrigin(h)
	mux.Handle("OPTIONS", "/signup", h.ServeHTTP)
	mux.Handle("OPTIONS", "/signin", h.ServeHTTP)
	mux.Handle("OPTIONS", "/resfresh", h.ServeHTTP)
	mux.Handle("OPTIONS", "/bo/signin", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleJWTTokenOrigin applies the CORS response headers corresponding to the
// origin for the service jwtToken.
func HandleJWTTokenOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile(".*team-gm.re.*")
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
