package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("users", func() {
	Description("users of the api")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Security(OAuth2, JWTAuth)
	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/v1/web/user")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Header("jwtToken:jwtToken", String, "Jwt token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("deleteUser", func() {
		Description("Delete one User by ID")
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
			DELETE("/remove/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})

	Method("updateAvatar", func() {
		Description("Update avatar")

		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("avatar", String, func() {
				Description("Url of the avatar and stock in db")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("id", "avatar")
		})

		Result(func() {
			Attribute("success", Boolean)
			Attribute("user", resUser, "Result is an Object")
			Required("success")
		})

		HTTP(func() {
			PUT("/edit/avatar")
			Response(StatusOK)
		})
	})

	Method("updateNumberStockage", func() {
		Description("Update number stockage")

		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("number", Int64, func() {
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("id", "number")
		})

		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})

		HTTP(func() {
			PUT("/edit/stock")
			Response(StatusOK)
		})
	})

	Method("getUserByID", func() {
		Description("Get one User")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Description("Unique ID of the User")
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
			Attribute("user", resUser, "Result is an object")
			Attribute("success", Boolean)
			Required("user", "success")
		})
	})

	Method("updateDescription", func() {
		Description("Update main info like email, username, firstname, lastname")

		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("email", String, func() {
				Example("guillaume@gmail.com")
				Format(FormatEmail)
			})
			Attribute("username", String, func() {
				Example("guillaumemoriin")
				MinLength(3)
				MaxLength(20)
				Pattern("^[a-z0-9_\\-]+$")
			})
			Attribute("firstname", String, func() {
				Example("Guillaume")
				MinLength(3)
				MaxLength(20)
			})
			Attribute("lastname", String, func() {
				Example("Morin")
				MinLength(3)
				MaxLength(20)
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("email", "id", "username")
		})

		Result(func() {
			Attribute("success", Boolean)
			Attribute("user", resUser, "Result is an Object")
			Required("success")
		})

		HTTP(func() {
			PUT("/edit")
			Response(StatusOK)
		})
	})

})

var resUser = Type("resUser", func() {
	Attribute("id", String, func() {
		Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
	})
	Attribute("firstname", String, func() {
		Example("Guillaume")
	})
	Attribute("lastname", String, func() {
		Example("Morin")
	})
	Attribute("username", String, func() {
		Example("guillaumemorin")
	})
	Attribute("email", String, func() {
		Example("guillaume@gmail.com")
	})
	Attribute("role", String, func() {
		Description("User is admin or not")
	})
	Attribute("avatar", String, func() {
		Example("/public/uploads/2022/02/eedf427a-559a-4faf-9e75-357fbb5d65fb.png")
	})
	Required("id", "email", "username", "firstname", "lastname", "avatar", "role")
})
