package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("boContact", func() {
	Description("back office contacts of the api")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Security(OAuth2, JWTAuth)

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/v1/bo/contacts")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Header("jwtToken:jwtToken", String, "Jwt token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("getBoContact", func() {
		Description("Get All messages")
		Payload(func() {
			Attribute("offset", Int32, func() {
				Description("Offset for pagination")
				Example(0)
			})
			Attribute("limit", Int32, func() {
				Description("Limit of items listed for pagination")
				Example(9)
			})
			Attribute("field", String, func() {
				Description("Items order by {field}")
				Example("name")
				Default("name")
			})
			Attribute("direction", String, func() {
				Description("Items order by {field} ASC/DESC")
				Enum("asc", "desc")
				Example("asc")
				Default("asc")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("limit", "offset")
		})
		HTTP(func() {
			GET("/messages/{offset}/{limit}")
			Params(func() {
				Param("field", String, func() {
					Description("Items order by {field}")
					Example("name")
					Default("name")
				})
				Param("direction", String, func() {
					Description("Items order by {field} ASC/DESC")
					Enum("asc", "desc")
					Example("asc")
					Default("asc")
				})
			})
			Response(StatusOK)
		})
		Result(func() {
			Attribute("contacts", ArrayOf(resContact), "All messages")
			Attribute("count", Int64, "total of messages")
			Attribute("success", Boolean)
			Required("contacts", "success", "count")
		})
	})

	Method("deleteBoContact", func() {
		Description("Delete one contact by ID")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("id")
		})
		HTTP(func() {
			PUT("/remove/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})

	Method("getBoContactByID", func() {
		Description("get one contact by ID")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("id")
		})
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Attribute("contact", resContact, "One message")
			Required("success")
		})
	})

	Method("deleteBoManyContact", func() {
		Description("Delete many contact with IDs send in body")
		Payload(func() {
			Attribute("tab", ArrayOf(String))
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("tab")
		})
		HTTP(func() {
			PATCH("/remove")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})
})

var resContact = Type("resContact", func() {
	Attribute("id", String, func() {
		Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
	})
	Attribute("user_id", String, func() {
		Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
	})
	Attribute("email", String, func() {
		Example("guillaume@gmail.com")
	})
	Attribute("message", String, func() {
		Example("Vous embauchez ?")
	})
	Required("id", "email", "user_id", "message")
})
