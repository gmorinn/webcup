// Code generated by goa v3.7.2, DO NOT EDIT.
//
// fileapi HTTP server
//
// Command:
// $ goa gen webcup/back/design

package server

import (
	"context"
	"net/http"
	"regexp"
	fileapi "webcup/back/gen/fileapi"

	goahttp "goa.design/goa/v3/http"
	"goa.design/plugins/v3/cors"
)

// Server lists the fileapi service endpoint HTTP handlers.
type Server struct {
	Mounts    []*MountPoint
	CORS      http.Handler
	BinPublic http.Handler
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

// New instantiates HTTP handlers for all the fileapi service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *fileapi.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	fileSystemBinPublic http.FileSystem,
) *Server {
	if fileSystemBinPublic == nil {
		fileSystemBinPublic = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"CORS", "OPTIONS", "/public/{*path}"},
			{"bin/public", "GET", "/public"},
		},
		CORS:      NewCORSHandler(),
		BinPublic: http.FileServer(fileSystemBinPublic),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "fileapi" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the fileapi endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCORSHandler(mux, h.CORS)
	MountBinPublic(mux, goahttp.Replace("/public", "/bin/public", h.BinPublic))
}

// Mount configures the mux to serve the fileapi endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountBinPublic configures the mux to serve GET request made to "/public".
func MountBinPublic(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/public/", HandleFileapiOrigin(h).ServeHTTP)
	mux.Handle("GET", "/public/*path", HandleFileapiOrigin(h).ServeHTTP)
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service fileapi.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleFileapiOrigin(h)
	mux.Handle("OPTIONS", "/public/{*path}", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleFileapiOrigin applies the CORS response headers corresponding to the
// origin for the service fileapi.
func HandleFileapiOrigin(h http.Handler) http.Handler {
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
