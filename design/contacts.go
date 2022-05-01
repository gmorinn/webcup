package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("contacts", func() {
	Description("contacts of the api")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Security(OAuth2, JWTAuth)

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/v1/web/contacts")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Header("jwtToken:jwtToken", String, "Jwt token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("addMessage", func() {
		Description("user ask for something")
		Payload(func() {
			Attribute("user_id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("msg", String, func() {
				Example("Je reprends l'app pour un million")
				MinLength(2)
				MaxLength(500)
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("user_id", "msg")
		})
		HTTP(func() {
			POST("/add")
			Response(StatusCreated)
		})
		Result(func() {
			Attribute("success", Boolean, func() {
				Default(false)
			})
			Required("success")
		})
	})
})
