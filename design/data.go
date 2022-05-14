package design

import . "goa.design/goa/v3/dsl"

// Service describes a service
var _ = Service("data", func() {
	Description("futristics data of the api")

	Error("timeout", func() { // Use default error type
		Timeout()
	})

	Security(OAuth2, JWTAuth)

	Error("unknown_error", unknownError, "Error not identified (500)")

	HTTP(func() {
		Path("/web/data")
		Header("oauth:Authorization", String, "OAuth token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Header("jwtToken:jwtToken", String, "Jwt token", func() {
			Pattern("^Bearer [^ ]+$")
		})
		Response("unknown_error", StatusInternalServerError)
	})

	Method("listDataMostRecent", func() {
		Description("List data the most recent")
		Payload(func() {
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
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
			Attribute("data", ArrayOf(resData), "Result is an an array of data")
			Attribute("success", Boolean)
			Attribute("count", Int64, "total of datas")
			Required("data", "success", "count")
		})
	})

	Method("createData", func() {
		Description("Create one data")
		Payload(func() {
			Attribute("data", payloadData)
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("data")
		})
		HTTP(func() {
			POST("/add")
			Response(StatusCreated)
		})
		Result(func() {
			Attribute("data", resData, "Result is an object")
			Attribute("success", Boolean)
			Required("data", "success")
		})
	})

	Method("updateData", func() {
		Description("Update one data")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			Attribute("data", payloadData)
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("data", "id")
		})
		HTTP(func() {
			PUT("/{id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("data", resData, "Result is an Object")
			Attribute("success", Boolean)
			Required("data", "success")
		})
	})

	Method("getDataByUserID", func() {
		Description("Get one data user id")
		Payload(func() {
			Attribute("user_id", String, func() {
				Format(FormatUUID)
				Description("Unique ID of the user")
				Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
			})
			TokenField(1, "jwtToken", String, func() {
				Description("JWT used for authentication after Signin/Signup")
			})
			AccessTokenField(2, "oauth", String, func() {
				Description("Use to generate Oauth with /authorization")
			})
			Required("user_id")
		})

		HTTP(func() {
			GET("user/{user_id}")
			Response(StatusOK)
		})
		Result(func() {
			Attribute("data", ArrayOf(resData), "Result is an object")
			Attribute("success", Boolean)
			Required("data", "success")
		})
	})

	Method("getDataByID", func() {
		Description("Get one data by id")
		Payload(func() {
			Attribute("id", String, func() {
				Format(FormatUUID)
				Description("Unique ID of the data")
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
			Attribute("data", resData, "Result is an object")
			Attribute("success", Boolean)
			Required("data", "success")
		})
	})
})

var resData = Type("resData", func() {
	Attribute("id", String, func() {
		Format(FormatUUID)
		Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
	})
	Attribute("title", String, func() {
		Example("Air max360 d'il y a 10 millions d'années")
		MinLength(3)
		MaxLength(20)
	})
	Attribute("description", String, func() {
		Example("Unique Air max restant au monde")
		MinLength(2)
		MaxLength(500)
	})
	Attribute("image", String, func() {
		Description("Url of the logo and stock in db")
	})
	Attribute("category", String, func() {
		Enum("robotics", "space", "brain", "animals")
		Default("robotics")
	})
	Attribute("user_id", String, func() {
		Format(FormatUUID)
		Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
	})
	Required("id", "title", "description", "image", "category", "user_id")
})

var payloadData = Type("payloadData", func() {
	Attribute("title", String, func() {
		Example("Air max360 d'il y a 10 millions d'années")
		MinLength(3)
		MaxLength(20)
	})
	Attribute("description", String, func() {
		Example("Unique Air max restant au monde")
		MinLength(2)
		MaxLength(500)
	})
	Attribute("image", String, func() {
		Description("Url of the logo and stock in db")
	})
	Attribute("category", String, func() {
		Enum("robotics", "space", "brain", "animals")
		Default("robotics")
	})
	Attribute("user_id", String, func() {
		Format(FormatUUID)
		Example("5dfb0bf7-597a-4250-b7ad-63a43ff59c25")
	})
	Required("title", "description", "image", "category", "user_id")
})
