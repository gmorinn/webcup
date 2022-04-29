package design

import (
	"fmt"
	"webcup/back/config"

	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
	_ "goa.design/plugins/v3/docs"
)

// API describes the global properties of the API server.
var _ = API("webcup", func() {

	// Get .env
	cnf := config.New()
	var host string

	if cnf.SSL {
		host = fmt.Sprintf("https://%s", cnf.Host)
	} else {
		host = fmt.Sprintf("http://%s", cnf.Host)
	}

	Title("webcup")
	Description("Best API REST building with GoaDesign")
	Version("1.0")

	cors.Origin(fmt.Sprintf("/.*%v.*/", "gm-influence.re"), func() {
		cors.Headers("Authorization", "Content-Type", "jwtToken", "Origin")
		cors.Methods("POST", "GET", "PUT", "OPTIONS", "DELETE", "PATCH")
		cors.Expose("Content-Type", "Origin")
		cors.MaxAge(600)
		cors.Credentials()
	})

	Server("webcup", func() {
		Host(cnf.Domain, func() {
			URI(host)
		})
	})

	Contact(func() {
		Name("Guillaume MORIN")
		Email("guillaume.morin@epitech.eu")
		URL("https://guillaume-morin.fr/")
	})

	License(func() {
		Name("GOA X GM")
		URL("https://guillaume-morin.fr/")
	})

})

// Download Postman
var _ = Service("openapi", func() {
	Files("/openapi.json", "openapi3.json", func() {
		Description("Postman generated by Goa")
		Docs(func() {
			Description("cd Public && live-server to see the Documentation of the API")
			URL("http://127.0.0.1:8080/")
		})
	})
})

// Access File
var _ = Service("fileapi", func() {
	Files("/public/{*path}", "bin/public", func() {
		Description("Serve static content.")
	})
})

var unknownError = Type("unknownError", func() {
	Field(1, "err", String, func() {
		Example("sql no rows affected")
	})
	Field(2, "error_code", String, func() {
		Example("TX_UPDATE_ITEM")
	})
	Field(3, "success", Boolean, func() {
		Default(false)
	})
	Required("err", "success", "error_code")
})