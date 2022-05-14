// Code generated by goa v3.7.3, DO NOT EDIT.
//
// bo HTTP server types
//
// Command:
// $ goa gen webcup/design

package server

import (
	"unicode/utf8"
	bo "webcup/gen/bo"

	goa "goa.design/goa/v3/pkg"
)

// DeleteBoManyUsersRequestBody is the type of the "bo" service
// "deleteBoManyUsers" endpoint HTTP request body.
type DeleteBoManyUsersRequestBody struct {
	Tab []string `form:"tab,omitempty" json:"tab,omitempty" xml:"tab,omitempty"`
}

// UpdateBoUserRequestBody is the type of the "bo" service "updateBoUser"
// endpoint HTTP request body.
type UpdateBoUserRequestBody struct {
	User *PayloadUserRequestBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
}

// GetBoUsersResponseBody is the type of the "bo" service "getBoUsers" endpoint
// HTTP response body.
type GetBoUsersResponseBody struct {
	// All users
	Users []*ResUserResponseBody `form:"users" json:"users" xml:"users"`
	// total of users
	Count   int64 `form:"count" json:"count" xml:"count"`
	Success bool  `form:"success" json:"success" xml:"success"`
}

// GetBoDataResponseBody is the type of the "bo" service "getBoData" endpoint
// HTTP response body.
type GetBoDataResponseBody struct {
	// All datas
	Data []*ResDataResponseBody `form:"data" json:"data" xml:"data"`
	// total of data
	Count   int64 `form:"count" json:"count" xml:"count"`
	Success bool  `form:"success" json:"success" xml:"success"`
}

// DeleteBoUserResponseBody is the type of the "bo" service "deleteBoUser"
// endpoint HTTP response body.
type DeleteBoUserResponseBody struct {
	Success bool `form:"success" json:"success" xml:"success"`
}

// DeleteBoManyUsersResponseBody is the type of the "bo" service
// "deleteBoManyUsers" endpoint HTTP response body.
type DeleteBoManyUsersResponseBody struct {
	Success bool `form:"success" json:"success" xml:"success"`
}

// UpdateBoUserResponseBody is the type of the "bo" service "updateBoUser"
// endpoint HTTP response body.
type UpdateBoUserResponseBody struct {
	// Result is an Object
	User    *ResUserResponseBody `form:"user" json:"user" xml:"user"`
	Success bool                 `form:"success" json:"success" xml:"success"`
}

// GetBoUserResponseBody is the type of the "bo" service "getBoUser" endpoint
// HTTP response body.
type GetBoUserResponseBody struct {
	// Result is an object
	User    *ResUserResponseBody `form:"user" json:"user" xml:"user"`
	Success bool                 `form:"success" json:"success" xml:"success"`
}

// GetBoUsersUnknownErrorResponseBody is the type of the "bo" service
// "getBoUsers" endpoint HTTP response body for the "unknown_error" error.
type GetBoUsersUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// GetBoDataUnknownErrorResponseBody is the type of the "bo" service
// "getBoData" endpoint HTTP response body for the "unknown_error" error.
type GetBoDataUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// DeleteBoUserUnknownErrorResponseBody is the type of the "bo" service
// "deleteBoUser" endpoint HTTP response body for the "unknown_error" error.
type DeleteBoUserUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// DeleteBoManyUsersUnknownErrorResponseBody is the type of the "bo" service
// "deleteBoManyUsers" endpoint HTTP response body for the "unknown_error"
// error.
type DeleteBoManyUsersUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// UpdateBoUserUnknownErrorResponseBody is the type of the "bo" service
// "updateBoUser" endpoint HTTP response body for the "unknown_error" error.
type UpdateBoUserUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// GetBoUserUnknownErrorResponseBody is the type of the "bo" service
// "getBoUser" endpoint HTTP response body for the "unknown_error" error.
type GetBoUserUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// ResUserResponseBody is used to define fields on response body types.
type ResUserResponseBody struct {
	ID        string `form:"id" json:"id" xml:"id"`
	Firstname string `form:"firstname" json:"firstname" xml:"firstname"`
	Lastname  string `form:"lastname" json:"lastname" xml:"lastname"`
	Username  string `form:"username" json:"username" xml:"username"`
	Email     string `form:"email" json:"email" xml:"email"`
	// User is admin or not
	Role   string `form:"role" json:"role" xml:"role"`
	Avatar string `form:"avatar" json:"avatar" xml:"avatar"`
}

// ResDataResponseBody is used to define fields on response body types.
type ResDataResponseBody struct {
	ID          string `form:"id" json:"id" xml:"id"`
	Title       string `form:"title" json:"title" xml:"title"`
	Description string `form:"description" json:"description" xml:"description"`
	// Url of the logo and stock in db
	Image    string `form:"image" json:"image" xml:"image"`
	Category string `form:"category" json:"category" xml:"category"`
	UserID   string `form:"user_id" json:"user_id" xml:"user_id"`
}

// PayloadUserRequestBody is used to define fields on request body types.
type PayloadUserRequestBody struct {
	Email     *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Username  *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	Firstname *string `form:"firstname,omitempty" json:"firstname,omitempty" xml:"firstname,omitempty"`
	Lastname  *string `form:"lastname,omitempty" json:"lastname,omitempty" xml:"lastname,omitempty"`
	// Url of the avatar and stock in db
	Avatar *string `form:"avatar,omitempty" json:"avatar,omitempty" xml:"avatar,omitempty"`
	// role of the user
	Role *string `form:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
}

// NewGetBoUsersResponseBody builds the HTTP response body from the result of
// the "getBoUsers" endpoint of the "bo" service.
func NewGetBoUsersResponseBody(res *bo.GetBoUsersResult) *GetBoUsersResponseBody {
	body := &GetBoUsersResponseBody{
		Count:   res.Count,
		Success: res.Success,
	}
	if res.Users != nil {
		body.Users = make([]*ResUserResponseBody, len(res.Users))
		for i, val := range res.Users {
			body.Users[i] = marshalBoResUserToResUserResponseBody(val)
		}
	}
	return body
}

// NewGetBoDataResponseBody builds the HTTP response body from the result of
// the "getBoData" endpoint of the "bo" service.
func NewGetBoDataResponseBody(res *bo.GetBoDataResult) *GetBoDataResponseBody {
	body := &GetBoDataResponseBody{
		Count:   res.Count,
		Success: res.Success,
	}
	if res.Data != nil {
		body.Data = make([]*ResDataResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalBoResDataToResDataResponseBody(val)
		}
	}
	return body
}

// NewDeleteBoUserResponseBody builds the HTTP response body from the result of
// the "deleteBoUser" endpoint of the "bo" service.
func NewDeleteBoUserResponseBody(res *bo.DeleteBoUserResult) *DeleteBoUserResponseBody {
	body := &DeleteBoUserResponseBody{
		Success: res.Success,
	}
	return body
}

// NewDeleteBoManyUsersResponseBody builds the HTTP response body from the
// result of the "deleteBoManyUsers" endpoint of the "bo" service.
func NewDeleteBoManyUsersResponseBody(res *bo.DeleteBoManyUsersResult) *DeleteBoManyUsersResponseBody {
	body := &DeleteBoManyUsersResponseBody{
		Success: res.Success,
	}
	return body
}

// NewUpdateBoUserResponseBody builds the HTTP response body from the result of
// the "updateBoUser" endpoint of the "bo" service.
func NewUpdateBoUserResponseBody(res *bo.UpdateBoUserResult) *UpdateBoUserResponseBody {
	body := &UpdateBoUserResponseBody{
		Success: res.Success,
	}
	if res.User != nil {
		body.User = marshalBoResUserToResUserResponseBody(res.User)
	}
	return body
}

// NewGetBoUserResponseBody builds the HTTP response body from the result of
// the "getBoUser" endpoint of the "bo" service.
func NewGetBoUserResponseBody(res *bo.GetBoUserResult) *GetBoUserResponseBody {
	body := &GetBoUserResponseBody{
		Success: res.Success,
	}
	if res.User != nil {
		body.User = marshalBoResUserToResUserResponseBody(res.User)
	}
	return body
}

// NewGetBoUsersUnknownErrorResponseBody builds the HTTP response body from the
// result of the "getBoUsers" endpoint of the "bo" service.
func NewGetBoUsersUnknownErrorResponseBody(res *bo.UnknownError) *GetBoUsersUnknownErrorResponseBody {
	body := &GetBoUsersUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewGetBoDataUnknownErrorResponseBody builds the HTTP response body from the
// result of the "getBoData" endpoint of the "bo" service.
func NewGetBoDataUnknownErrorResponseBody(res *bo.UnknownError) *GetBoDataUnknownErrorResponseBody {
	body := &GetBoDataUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewDeleteBoUserUnknownErrorResponseBody builds the HTTP response body from
// the result of the "deleteBoUser" endpoint of the "bo" service.
func NewDeleteBoUserUnknownErrorResponseBody(res *bo.UnknownError) *DeleteBoUserUnknownErrorResponseBody {
	body := &DeleteBoUserUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewDeleteBoManyUsersUnknownErrorResponseBody builds the HTTP response body
// from the result of the "deleteBoManyUsers" endpoint of the "bo" service.
func NewDeleteBoManyUsersUnknownErrorResponseBody(res *bo.UnknownError) *DeleteBoManyUsersUnknownErrorResponseBody {
	body := &DeleteBoManyUsersUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewUpdateBoUserUnknownErrorResponseBody builds the HTTP response body from
// the result of the "updateBoUser" endpoint of the "bo" service.
func NewUpdateBoUserUnknownErrorResponseBody(res *bo.UnknownError) *UpdateBoUserUnknownErrorResponseBody {
	body := &UpdateBoUserUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewGetBoUserUnknownErrorResponseBody builds the HTTP response body from the
// result of the "getBoUser" endpoint of the "bo" service.
func NewGetBoUserUnknownErrorResponseBody(res *bo.UnknownError) *GetBoUserUnknownErrorResponseBody {
	body := &GetBoUserUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewGetBoUsersPayload builds a bo service getBoUsers endpoint payload.
func NewGetBoUsersPayload(offset int32, limit int32, field string, direction string, oauth *string, jwtToken *string) *bo.GetBoUsersPayload {
	v := &bo.GetBoUsersPayload{}
	v.Offset = offset
	v.Limit = limit
	v.Field = field
	v.Direction = direction
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// NewGetBoDataPayload builds a bo service getBoData endpoint payload.
func NewGetBoDataPayload(offset int32, limit int32, field string, direction string, oauth *string, jwtToken *string) *bo.GetBoDataPayload {
	v := &bo.GetBoDataPayload{}
	v.Offset = offset
	v.Limit = limit
	v.Field = field
	v.Direction = direction
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// NewDeleteBoUserPayload builds a bo service deleteBoUser endpoint payload.
func NewDeleteBoUserPayload(id string, oauth *string, jwtToken *string) *bo.DeleteBoUserPayload {
	v := &bo.DeleteBoUserPayload{}
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// NewDeleteBoManyUsersPayload builds a bo service deleteBoManyUsers endpoint
// payload.
func NewDeleteBoManyUsersPayload(body *DeleteBoManyUsersRequestBody, oauth *string, jwtToken *string) *bo.DeleteBoManyUsersPayload {
	v := &bo.DeleteBoManyUsersPayload{}
	v.Tab = make([]string, len(body.Tab))
	for i, val := range body.Tab {
		v.Tab[i] = val
	}
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// NewUpdateBoUserPayload builds a bo service updateBoUser endpoint payload.
func NewUpdateBoUserPayload(body *UpdateBoUserRequestBody, id string, oauth *string, jwtToken *string) *bo.UpdateBoUserPayload {
	v := &bo.UpdateBoUserPayload{}
	v.User = unmarshalPayloadUserRequestBodyToBoPayloadUser(body.User)
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// NewGetBoUserPayload builds a bo service getBoUser endpoint payload.
func NewGetBoUserPayload(id string, oauth *string, jwtToken *string) *bo.GetBoUserPayload {
	v := &bo.GetBoUserPayload{}
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// ValidateDeleteBoManyUsersRequestBody runs the validations defined on
// DeleteBoManyUsersRequestBody
func ValidateDeleteBoManyUsersRequestBody(body *DeleteBoManyUsersRequestBody) (err error) {
	if body.Tab == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tab", "body"))
	}
	return
}

// ValidateUpdateBoUserRequestBody runs the validations defined on
// UpdateBoUserRequestBody
func ValidateUpdateBoUserRequestBody(body *UpdateBoUserRequestBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.User != nil {
		if err2 := ValidatePayloadUserRequestBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePayloadUserRequestBody runs the validations defined on
// payloadUserRequestBody
func ValidatePayloadUserRequestBody(body *PayloadUserRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.Firstname == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firstname", "body"))
	}
	if body.Lastname == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastname", "body"))
	}
	if body.Avatar == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("avatar", "body"))
	}
	if body.Role == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("role", "body"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.Username != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.username", *body.Username, "^[a-z0-9_\\-]+$"))
	}
	if body.Username != nil {
		if utf8.RuneCountInString(*body.Username) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.username", *body.Username, utf8.RuneCountInString(*body.Username), 2, true))
		}
	}
	if body.Username != nil {
		if utf8.RuneCountInString(*body.Username) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.username", *body.Username, utf8.RuneCountInString(*body.Username), 20, false))
		}
	}
	if body.Firstname != nil {
		if utf8.RuneCountInString(*body.Firstname) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", *body.Firstname, utf8.RuneCountInString(*body.Firstname), 3, true))
		}
	}
	if body.Firstname != nil {
		if utf8.RuneCountInString(*body.Firstname) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.firstname", *body.Firstname, utf8.RuneCountInString(*body.Firstname), 20, false))
		}
	}
	if body.Lastname != nil {
		if utf8.RuneCountInString(*body.Lastname) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.lastname", *body.Lastname, utf8.RuneCountInString(*body.Lastname), 3, true))
		}
	}
	if body.Lastname != nil {
		if utf8.RuneCountInString(*body.Lastname) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.lastname", *body.Lastname, utf8.RuneCountInString(*body.Lastname), 20, false))
		}
	}
	if body.Role != nil {
		if !(*body.Role == "admin" || *body.Role == "user" || *body.Role == "pro") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.role", *body.Role, []interface{}{"admin", "user", "pro"}))
		}
	}
	return
}
