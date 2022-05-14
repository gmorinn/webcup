package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"webcup/api"
	"webcup/gen/auth"
	bo "webcup/gen/bo"
	bo_contact "webcup/gen/bo_contact"
	contacts "webcup/gen/contacts"
	data "webcup/gen/data"
	fileapi "webcup/gen/fileapi"
	files "webcup/gen/files"
	jwttoken "webcup/gen/jwt_token"
	oauth "webcup/gen/o_auth"
	public_users "webcup/gen/public_users"
	users "webcup/gen/users"
)

type ApiEndpoints struct {
	dataEndpoints         *data.Endpoints
	contactsEndpoints     *contacts.Endpoints
	public_usersEndpoints *public_users.Endpoints
	usersEndpoints        *users.Endpoints
	fileapiEndpoints      *fileapi.Endpoints
	filesEndpoints        *files.Endpoints
	authEndpoints         *auth.Endpoints
	jwtTokenEndpoints     *jwttoken.Endpoints
	oAuthEndpoints        *oauth.Endpoints
	boEndpoints           *bo.Endpoints
	boContactEndpoints    *bo_contact.Endpoints
}

func main() {
	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
		server *api.Server
	)
	{
		logger = log.New(os.Stderr, fmt.Sprintf("[%v]", "GM API"), log.Ltime)
		server = api.NewServer()
	}

	// Initialize the services.
	var (
		authSvc         auth.Service         = api.NewAuth(logger, server)
		jwtTokenSvc     jwttoken.Service     = api.NewJWTToken(logger, server)
		dataSvc         data.Service         = api.NewData(logger, server)
		boSvc           bo.Service           = api.NewBo(logger, server)
		contactsSvc     contacts.Service     = api.NewContacts(logger, server)
		public_usersSvc public_users.Service = api.NewPublicUsers(logger, server)
		usersSvc        users.Service        = api.NewUsers(logger, server)
		fileapiSvc      fileapi.Service      = api.NewFileapi(logger)
		filesSvc        files.Service        = api.NewFiles(logger, server)
		oAuthSvc        oauth.Service        = api.NewOAuth(logger, server)
		boContactSvc    bo_contact.Service   = api.NewBoContact(logger, server)
	)

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		apiEndpoints ApiEndpoints = ApiEndpoints{
			dataEndpoints:         data.NewEndpoints(dataSvc),
			contactsEndpoints:     contacts.NewEndpoints(contactsSvc),
			public_usersEndpoints: public_users.NewEndpoints(public_usersSvc),
			usersEndpoints:        users.NewEndpoints(usersSvc),
			authEndpoints:         auth.NewEndpoints(authSvc),
			fileapiEndpoints:      fileapi.NewEndpoints(fileapiSvc),
			filesEndpoints:        files.NewEndpoints(filesSvc),
			jwtTokenEndpoints:     jwttoken.NewEndpoints(jwtTokenSvc),
			oAuthEndpoints:        oauth.NewEndpoints(oAuthSvc),
			boEndpoints:           bo.NewEndpoints(boSvc),
			boContactEndpoints:    bo_contact.NewEndpoints(boContactSvc),
		}
	)
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", server.Config.Domain, "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", server.Config.SSL, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	var addr string = fmt.Sprintf("http://%s", server.Config.Host)

	fmt.Printf("API on %s\n", addr)

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case server.Config.Domain:
		{
			addr := addr
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":3000"
			}
			handleHTTPServer(ctx, u, &apiEndpoints, &wg, errc, logger, *dbgF)
		}
	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: %v)\n", *hostF, server.Config.Host)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
