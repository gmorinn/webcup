// Code generated by goa v3.7.3, DO NOT EDIT.
//
// bo HTTP server
//
// Command:
// $ goa gen webcup/design

package server

import (
	"context"
	"net/http"
	"regexp"
	bo "webcup/gen/bo"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the bo service endpoint HTTP handlers.
type Server struct {
	Mounts            []*MountPoint
	GetBoUsers        http.Handler
	GetBoData         http.Handler
	DeleteBoUser      http.Handler
	DeleteBoManyUsers http.Handler
	UpdateBoUser      http.Handler
	GetBoUser         http.Handler
	CORS              http.Handler
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

// New instantiates HTTP handlers for all the bo service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *bo.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"GetBoUsers", "GET", "/v1/bo/users/{offset}/{limit}"},
			{"GetBoData", "GET", "/v1/bo/datas/{offset}/{limit}"},
			{"DeleteBoUser", "DELETE", "/v1/bo/user/remove/{id}"},
			{"DeleteBoManyUsers", "PATCH", "/v1/bo/users/remove"},
			{"UpdateBoUser", "PUT", "/v1/bo/user/{id}"},
			{"GetBoUser", "GET", "/v1/bo/user/{id}"},
			{"CORS", "OPTIONS", "/v1/bo/users/{offset}/{limit}"},
			{"CORS", "OPTIONS", "/v1/bo/datas/{offset}/{limit}"},
			{"CORS", "OPTIONS", "/v1/bo/user/remove/{id}"},
			{"CORS", "OPTIONS", "/v1/bo/users/remove"},
			{"CORS", "OPTIONS", "/v1/bo/user/{id}"},
		},
		GetBoUsers:        NewGetBoUsersHandler(e.GetBoUsers, mux, decoder, encoder, errhandler, formatter),
		GetBoData:         NewGetBoDataHandler(e.GetBoData, mux, decoder, encoder, errhandler, formatter),
		DeleteBoUser:      NewDeleteBoUserHandler(e.DeleteBoUser, mux, decoder, encoder, errhandler, formatter),
		DeleteBoManyUsers: NewDeleteBoManyUsersHandler(e.DeleteBoManyUsers, mux, decoder, encoder, errhandler, formatter),
		UpdateBoUser:      NewUpdateBoUserHandler(e.UpdateBoUser, mux, decoder, encoder, errhandler, formatter),
		GetBoUser:         NewGetBoUserHandler(e.GetBoUser, mux, decoder, encoder, errhandler, formatter),
		CORS:              NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "bo" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.GetBoUsers = m(s.GetBoUsers)
	s.GetBoData = m(s.GetBoData)
	s.DeleteBoUser = m(s.DeleteBoUser)
	s.DeleteBoManyUsers = m(s.DeleteBoManyUsers)
	s.UpdateBoUser = m(s.UpdateBoUser)
	s.GetBoUser = m(s.GetBoUser)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the bo endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetBoUsersHandler(mux, h.GetBoUsers)
	MountGetBoDataHandler(mux, h.GetBoData)
	MountDeleteBoUserHandler(mux, h.DeleteBoUser)
	MountDeleteBoManyUsersHandler(mux, h.DeleteBoManyUsers)
	MountUpdateBoUserHandler(mux, h.UpdateBoUser)
	MountGetBoUserHandler(mux, h.GetBoUser)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the bo endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountGetBoUsersHandler configures the mux to serve the "bo" service
// "getBoUsers" endpoint.
func MountGetBoUsersHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/bo/users/{offset}/{limit}", f)
}

// NewGetBoUsersHandler creates a HTTP handler which loads the HTTP request and
// calls the "bo" service "getBoUsers" endpoint.
func NewGetBoUsersHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetBoUsersRequest(mux, decoder)
		encodeResponse = EncodeGetBoUsersResponse(encoder)
		encodeError    = EncodeGetBoUsersError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getBoUsers")
		ctx = context.WithValue(ctx, goa.ServiceKey, "bo")
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

// MountGetBoDataHandler configures the mux to serve the "bo" service
// "getBoData" endpoint.
func MountGetBoDataHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/bo/datas/{offset}/{limit}", f)
}

// NewGetBoDataHandler creates a HTTP handler which loads the HTTP request and
// calls the "bo" service "getBoData" endpoint.
func NewGetBoDataHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetBoDataRequest(mux, decoder)
		encodeResponse = EncodeGetBoDataResponse(encoder)
		encodeError    = EncodeGetBoDataError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getBoData")
		ctx = context.WithValue(ctx, goa.ServiceKey, "bo")
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

// MountDeleteBoUserHandler configures the mux to serve the "bo" service
// "deleteBoUser" endpoint.
func MountDeleteBoUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/v1/bo/user/remove/{id}", f)
}

// NewDeleteBoUserHandler creates a HTTP handler which loads the HTTP request
// and calls the "bo" service "deleteBoUser" endpoint.
func NewDeleteBoUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteBoUserRequest(mux, decoder)
		encodeResponse = EncodeDeleteBoUserResponse(encoder)
		encodeError    = EncodeDeleteBoUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "deleteBoUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "bo")
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

// MountDeleteBoManyUsersHandler configures the mux to serve the "bo" service
// "deleteBoManyUsers" endpoint.
func MountDeleteBoManyUsersHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PATCH", "/v1/bo/users/remove", f)
}

// NewDeleteBoManyUsersHandler creates a HTTP handler which loads the HTTP
// request and calls the "bo" service "deleteBoManyUsers" endpoint.
func NewDeleteBoManyUsersHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteBoManyUsersRequest(mux, decoder)
		encodeResponse = EncodeDeleteBoManyUsersResponse(encoder)
		encodeError    = EncodeDeleteBoManyUsersError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "deleteBoManyUsers")
		ctx = context.WithValue(ctx, goa.ServiceKey, "bo")
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

// MountUpdateBoUserHandler configures the mux to serve the "bo" service
// "updateBoUser" endpoint.
func MountUpdateBoUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/v1/bo/user/{id}", f)
}

// NewUpdateBoUserHandler creates a HTTP handler which loads the HTTP request
// and calls the "bo" service "updateBoUser" endpoint.
func NewUpdateBoUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateBoUserRequest(mux, decoder)
		encodeResponse = EncodeUpdateBoUserResponse(encoder)
		encodeError    = EncodeUpdateBoUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "updateBoUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "bo")
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

// MountGetBoUserHandler configures the mux to serve the "bo" service
// "getBoUser" endpoint.
func MountGetBoUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleBoOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/bo/user/{id}", f)
}

// NewGetBoUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "bo" service "getBoUser" endpoint.
func NewGetBoUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetBoUserRequest(mux, decoder)
		encodeResponse = EncodeGetBoUserResponse(encoder)
		encodeError    = EncodeGetBoUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getBoUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "bo")
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
// service bo.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleBoOrigin(h)
	mux.Handle("OPTIONS", "/v1/bo/users/{offset}/{limit}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/bo/datas/{offset}/{limit}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/bo/user/remove/{id}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/bo/users/remove", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/bo/user/{id}", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleBoOrigin applies the CORS response headers corresponding to the origin
// for the service bo.
func HandleBoOrigin(h http.Handler) http.Handler {
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
