// Code generated by goa v3.7.2, DO NOT EDIT.
//
// auth HTTP server
//
// Command:
// $ goa gen webcup/back/design

package server

import (
	"context"
	"net/http"
	"regexp"
	auth "webcup/back/gen/auth"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the auth service endpoint HTTP handlers.
type Server struct {
	Mounts           []*MountPoint
	EmailExist       http.Handler
	SendConfirmation http.Handler
	ResetPassword    http.Handler
	CORS             http.Handler
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

// New instantiates HTTP handlers for all the auth service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *auth.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"EmailExist", "POST", "/v1/email-exist"},
			{"SendConfirmation", "POST", "/v1/lost"},
			{"ResetPassword", "PUT", "/v1/reset-password"},
			{"CORS", "OPTIONS", "/v1/email-exist"},
			{"CORS", "OPTIONS", "/v1/lost"},
			{"CORS", "OPTIONS", "/v1/reset-password"},
		},
		EmailExist:       NewEmailExistHandler(e.EmailExist, mux, decoder, encoder, errhandler, formatter),
		SendConfirmation: NewSendConfirmationHandler(e.SendConfirmation, mux, decoder, encoder, errhandler, formatter),
		ResetPassword:    NewResetPasswordHandler(e.ResetPassword, mux, decoder, encoder, errhandler, formatter),
		CORS:             NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "auth" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.EmailExist = m(s.EmailExist)
	s.SendConfirmation = m(s.SendConfirmation)
	s.ResetPassword = m(s.ResetPassword)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the auth endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountEmailExistHandler(mux, h.EmailExist)
	MountSendConfirmationHandler(mux, h.SendConfirmation)
	MountResetPasswordHandler(mux, h.ResetPassword)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the auth endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountEmailExistHandler configures the mux to serve the "auth" service
// "email-exist" endpoint.
func MountEmailExistHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/email-exist", f)
}

// NewEmailExistHandler creates a HTTP handler which loads the HTTP request and
// calls the "auth" service "email-exist" endpoint.
func NewEmailExistHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeEmailExistRequest(mux, decoder)
		encodeResponse = EncodeEmailExistResponse(encoder)
		encodeError    = EncodeEmailExistError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "email-exist")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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

// MountSendConfirmationHandler configures the mux to serve the "auth" service
// "send-confirmation" endpoint.
func MountSendConfirmationHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/lost", f)
}

// NewSendConfirmationHandler creates a HTTP handler which loads the HTTP
// request and calls the "auth" service "send-confirmation" endpoint.
func NewSendConfirmationHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSendConfirmationRequest(mux, decoder)
		encodeResponse = EncodeSendConfirmationResponse(encoder)
		encodeError    = EncodeSendConfirmationError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "send-confirmation")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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

// MountResetPasswordHandler configures the mux to serve the "auth" service
// "reset-password" endpoint.
func MountResetPasswordHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/v1/reset-password", f)
}

// NewResetPasswordHandler creates a HTTP handler which loads the HTTP request
// and calls the "auth" service "reset-password" endpoint.
func NewResetPasswordHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeResetPasswordRequest(mux, decoder)
		encodeResponse = EncodeResetPasswordResponse(encoder)
		encodeError    = EncodeResetPasswordError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "reset-password")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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
// service auth.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleAuthOrigin(h)
	mux.Handle("OPTIONS", "/v1/email-exist", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/lost", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/reset-password", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleAuthOrigin applies the CORS response headers corresponding to the
// origin for the service auth.
func HandleAuthOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile(".*gm-influence.re.*")
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
