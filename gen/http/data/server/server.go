// Code generated by goa v3.7.3, DO NOT EDIT.
//
// data HTTP server
//
// Command:
// $ goa gen webcup/design

package server

import (
	"context"
	"net/http"
	"regexp"
	data "webcup/gen/data"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the data service endpoint HTTP handlers.
type Server struct {
	Mounts             []*MountPoint
	ListDataMostRecent http.Handler
	CreateData         http.Handler
	UpdateData         http.Handler
	GetDataByUserID    http.Handler
	GetDataByID        http.Handler
	CORS               http.Handler
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

// New instantiates HTTP handlers for all the data service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *data.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ListDataMostRecent", "GET", "/web/data/recents/{offset}/{limit}"},
			{"CreateData", "POST", "/web/data/add"},
			{"UpdateData", "PUT", "/web/data/{id}"},
			{"GetDataByUserID", "GET", "/web/data/user/{user_id}"},
			{"GetDataByID", "GET", "/web/data/{id}"},
			{"CORS", "OPTIONS", "/web/data/recents/{offset}/{limit}"},
			{"CORS", "OPTIONS", "/web/data/add"},
			{"CORS", "OPTIONS", "/web/data/{id}"},
			{"CORS", "OPTIONS", "/web/data/user/{user_id}"},
		},
		ListDataMostRecent: NewListDataMostRecentHandler(e.ListDataMostRecent, mux, decoder, encoder, errhandler, formatter),
		CreateData:         NewCreateDataHandler(e.CreateData, mux, decoder, encoder, errhandler, formatter),
		UpdateData:         NewUpdateDataHandler(e.UpdateData, mux, decoder, encoder, errhandler, formatter),
		GetDataByUserID:    NewGetDataByUserIDHandler(e.GetDataByUserID, mux, decoder, encoder, errhandler, formatter),
		GetDataByID:        NewGetDataByIDHandler(e.GetDataByID, mux, decoder, encoder, errhandler, formatter),
		CORS:               NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "data" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ListDataMostRecent = m(s.ListDataMostRecent)
	s.CreateData = m(s.CreateData)
	s.UpdateData = m(s.UpdateData)
	s.GetDataByUserID = m(s.GetDataByUserID)
	s.GetDataByID = m(s.GetDataByID)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the data endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListDataMostRecentHandler(mux, h.ListDataMostRecent)
	MountCreateDataHandler(mux, h.CreateData)
	MountUpdateDataHandler(mux, h.UpdateData)
	MountGetDataByUserIDHandler(mux, h.GetDataByUserID)
	MountGetDataByIDHandler(mux, h.GetDataByID)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the data endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountListDataMostRecentHandler configures the mux to serve the "data"
// service "listDataMostRecent" endpoint.
func MountListDataMostRecentHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleDataOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/web/data/recents/{offset}/{limit}", f)
}

// NewListDataMostRecentHandler creates a HTTP handler which loads the HTTP
// request and calls the "data" service "listDataMostRecent" endpoint.
func NewListDataMostRecentHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListDataMostRecentRequest(mux, decoder)
		encodeResponse = EncodeListDataMostRecentResponse(encoder)
		encodeError    = EncodeListDataMostRecentError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "listDataMostRecent")
		ctx = context.WithValue(ctx, goa.ServiceKey, "data")
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

// MountCreateDataHandler configures the mux to serve the "data" service
// "createData" endpoint.
func MountCreateDataHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleDataOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/web/data/add", f)
}

// NewCreateDataHandler creates a HTTP handler which loads the HTTP request and
// calls the "data" service "createData" endpoint.
func NewCreateDataHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateDataRequest(mux, decoder)
		encodeResponse = EncodeCreateDataResponse(encoder)
		encodeError    = EncodeCreateDataError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "createData")
		ctx = context.WithValue(ctx, goa.ServiceKey, "data")
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

// MountUpdateDataHandler configures the mux to serve the "data" service
// "updateData" endpoint.
func MountUpdateDataHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleDataOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/web/data/{id}", f)
}

// NewUpdateDataHandler creates a HTTP handler which loads the HTTP request and
// calls the "data" service "updateData" endpoint.
func NewUpdateDataHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateDataRequest(mux, decoder)
		encodeResponse = EncodeUpdateDataResponse(encoder)
		encodeError    = EncodeUpdateDataError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "updateData")
		ctx = context.WithValue(ctx, goa.ServiceKey, "data")
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

// MountGetDataByUserIDHandler configures the mux to serve the "data" service
// "getDataByUserID" endpoint.
func MountGetDataByUserIDHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleDataOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/web/data/user/{user_id}", f)
}

// NewGetDataByUserIDHandler creates a HTTP handler which loads the HTTP
// request and calls the "data" service "getDataByUserID" endpoint.
func NewGetDataByUserIDHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetDataByUserIDRequest(mux, decoder)
		encodeResponse = EncodeGetDataByUserIDResponse(encoder)
		encodeError    = EncodeGetDataByUserIDError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getDataByUserID")
		ctx = context.WithValue(ctx, goa.ServiceKey, "data")
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

// MountGetDataByIDHandler configures the mux to serve the "data" service
// "getDataByID" endpoint.
func MountGetDataByIDHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleDataOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/web/data/{id}", f)
}

// NewGetDataByIDHandler creates a HTTP handler which loads the HTTP request
// and calls the "data" service "getDataByID" endpoint.
func NewGetDataByIDHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetDataByIDRequest(mux, decoder)
		encodeResponse = EncodeGetDataByIDResponse(encoder)
		encodeError    = EncodeGetDataByIDError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getDataByID")
		ctx = context.WithValue(ctx, goa.ServiceKey, "data")
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
// service data.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleDataOrigin(h)
	mux.Handle("OPTIONS", "/web/data/recents/{offset}/{limit}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/web/data/add", h.ServeHTTP)
	mux.Handle("OPTIONS", "/web/data/{id}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/web/data/user/{user_id}", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleDataOrigin applies the CORS response headers corresponding to the
// origin for the service data.
func HandleDataOrigin(h http.Handler) http.Handler {
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
