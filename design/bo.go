package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("bo", func() {
	Description("bo of the api")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Security(OAuth2, JWTAuth)

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/v1/bo")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Header("jwtToken:jwtToken", String, "Jwt token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("getBoUsers", func() {
		Description("Get All users")
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
			GET("/users/{offset}/{limit}")
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
			Attribute("users", ArrayOf(resUser), "All users")
			Attribute("count", Int64, "total of users")
			Attribute("success", Boolean)
			Required("users", "success", "count")
		})
	})

	Method("getBoData", func() {
		Description("Get All data")
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
			GET("/users/datas/{offset}/{limit}")
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
			Attribute("data", ArrayOf(resData), "All datas")
			Attribute("count", Int64, "total of data")
			Attribute("success", Boolean)
			Required("data", "success", "count")
		})
	})

	Method("deleteBoUser", func() {
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
			DELETE("/user/remove/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})

	Method("deleteBoManyUsers", func() {
		Description("Delete many users with IDs send in body")
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
			PATCH("/users/remove")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("success", Boolean)
			Required("success")
		})
	})

	Method("updateBoUser", func() {
		Description("Update one User")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("user", payloadBoUser)
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("user", "id")
		})
		HTTP(func() {
			PUT("/user/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("user", resUser, "Result is an Object")
			Attribute("success", Boolean)
			Required("user", "success")
		})
	})

	Method("getBoUser", func() {
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
			GET("/user/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("user", resUser, "Result is an object")
			Attribute("success", Boolean)
			Required("user", "success")
		})
	})
})

var payloadBoUser = Type("payloadUser", func() {
	Attribute("email", String, func() {
		Example("guillaume@gmail.com")
		Format(FormatEmail)
	})
	Attribute("username", String, func() {
		Example("guillaumemoriin")
		MinLength(2)
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
	Attribute("avatar", String, func() {
		Description("Url of the avatar and stock in db")
	})
	Attribute("role", String, func() {
		Description("role of the user")
		Enum("admin", "user", "pro")
	})
	Required("email", "username", "firstname", "lastname", "avatar", "role")
})
