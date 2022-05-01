package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
	filesclient "webcup/api"
	"webcup/config"
	authsvr "webcup/gen/http/auth/server"
	bosvr "webcup/gen/http/bo/server"
	bocontactsvr "webcup/gen/http/bo_contact/server"
	contactssvr "webcup/gen/http/contacts/server"
	fileapisvr "webcup/gen/http/fileapi/server"
	filessvr "webcup/gen/http/files/server"
	jwttokensvr "webcup/gen/http/jwt_token/server"
	oauthsvr "webcup/gen/http/o_auth/server"
	openapisvr "webcup/gen/http/openapi/server"
	public_userssvr "webcup/gen/http/public_users/server"
	userssvr "webcup/gen/http/users/server"

	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, api *ApiEndpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Get .env
	cnf := config.New()

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.

	var filePath string
	if cnf.Mode == "PROD" {
		filePath = "/go/"
	} else {
		filePath = "./"
	}

	var (
		eh                                         = errorHandler(logger)
		openapiServer      *openapisvr.Server      = openapisvr.New(nil, mux, dec, enc, nil, nil, http.Dir("gen/http"))
		usersServer        *userssvr.Server        = userssvr.New(api.usersEndpoints, mux, dec, enc, eh, nil)
		filesServer        *filessvr.Server        = filessvr.New(api.filesEndpoints, mux, dec, enc, eh, nil, filesclient.FilesImportFileDecoderFunc)
		jwtTokenServer     *jwttokensvr.Server     = jwttokensvr.New(api.jwtTokenEndpoints, mux, dec, enc, eh, nil)
		contactsServer     *contactssvr.Server     = contactssvr.New(api.contactsEndpoints, mux, dec, enc, eh, nil)
		public_usersServer *public_userssvr.Server = public_userssvr.New(api.public_usersEndpoints, mux, dec, enc, eh, nil)
		fileapiServer      *fileapisvr.Server      = fileapisvr.New(nil, mux, dec, enc, nil, nil, http.Dir(filePath))
		oAuthServer        *oauthsvr.Server        = oauthsvr.New(api.oAuthEndpoints, mux, dec, enc, eh, nil)
		authServer         *authsvr.Server         = authsvr.New(api.authEndpoints, mux, dec, enc, eh, nil)
		boServer           *bosvr.Server           = bosvr.New(api.boEndpoints, mux, dec, enc, eh, nil)
		boContactServer    *bocontactsvr.Server    = bocontactsvr.New(api.boContactEndpoints, mux, dec, enc, eh, nil)
	)
	{
		if debug {
			servers := goahttp.Servers{
				contactsServer,
				public_usersServer,
				usersServer,
				fileapiServer,
				filesServer,
				openapiServer,
				jwtTokenServer,
				authServer,
				oAuthServer,
				boServer,
				boContactServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	openapisvr.Mount(mux, openapiServer)
	contactssvr.Mount(mux, contactsServer)
	public_userssvr.Mount(mux, public_usersServer)
	userssvr.Mount(mux, usersServer)
	fileapisvr.Mount(mux, fileapiServer)
	authsvr.Mount(mux, authServer)
	filessvr.Mount(mux, filesServer)
	jwttokensvr.Mount(mux, jwtTokenServer)
	oauthsvr.Mount(mux, oAuthServer)
	bosvr.Mount(mux, boServer)
	bocontactsvr.Mount(mux, boContactServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	var srv *http.Server = &http.Server{Addr: fmt.Sprintf(":%d", cnf.Port), Handler: handler}
	logger.Printf(`
	
	 ██████╗  ██████╗  █████╗     ██╗  ██╗     ██████╗ ███╗   ███╗
	██╔════╝ ██╔═══██╗██╔══██╗    ╚██╗██╔╝    ██╔════╝ ████╗ ████║
	██║  ███╗██║   ██║███████║     ╚███╔╝     ██║  ███╗██╔████╔██║
	██║   ██║██║   ██║██╔══██║     ██╔██╗     ██║   ██║██║╚██╔╝██║
	╚██████╔╝╚██████╔╝██║  ██║    ██╔╝ ██╗    ╚██████╔╝██║ ╚═╝ ██║
	 ╚═════╝  ╚═════╝ ╚═╝  ╚═╝    ╚═╝  ╚═╝     ╚═════╝ ╚═╝     ╚═╝
																  
																  `)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			errc <- srv.ListenAndServe()
			logger.Printf("HTTP server listening on %q", u.Host)
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
