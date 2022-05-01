package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("publicUsers", func() {
	Description("public route of users")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Security(OAuth2)

	Error("unknown_error", unknownError, "Error not identified (500)")
	Error("oauth_error", String, "Error when a request is send before asking for oAuth")

	HTTP(func() {
		Path("/v1/web/public/user")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
		Response("oauth_error", StatusForbidden)
	})

	Method("getUserByUsername", func() {
		Description("Get one User by username")
		Payload(func() {
			Attribute("username", String, func() {
				Example("guillaumemoriin")
				MinLength(3)
				MaxLength(20)
				Pattern("^[a-z0-9_\\-]+$")
			})
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("username")
		})

		HTTP(func() {
			GET("/{username}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("user", resUser, "Result is an object")
			Attribute("success", Boolean)
			Required("user", "success")
		})
	})

	Method("listUsers", func() {
		Description("List users for search bar")
		Payload(func() {
			Attribute("key", String, func() {
				Example("guillaumemoriin")
			})
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("key")
		})

		HTTP(func() {
			PATCH("/search")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("users", ArrayOf(resUser), "Result is an an array of user")
			Attribute("success", Boolean)
			Required("users", "success")
		})
	})

	Method("listUsersMostRecent", func() {
		Description("List users the most recent")
		Payload(func() {
			AccessTokenField(1, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Attribute("offset", Int32, func() {
				Description("Offset for pagination")
				Example(0)
				Minimum(0)
			})
			Attribute("limit", Int32, func() {
				Description("Limit of items listed for pagination")
				Example(5)
				Minimum(0)
			})
			Required("limit", "offset")
		})

		HTTP(func() {
			GET("/recents/{offset}/{limit}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("users", ArrayOf(resUser), "Result is an an array of user")
			Attribute("success", Boolean)
			Attribute("count", Int64, "total of users")
			Required("users", "success", "count")
		})
	})
})
