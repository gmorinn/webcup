// Code generated by goa v3.7.3, DO NOT EDIT.
//
// contacts HTTP server types
//
// Command:
// $ goa gen webcup/design

package server

import (
	"unicode/utf8"
	contacts "webcup/gen/contacts"

	goa "goa.design/goa/v3/pkg"
)

// AddMessageRequestBody is the type of the "contacts" service "addMessage"
// endpoint HTTP request body.
type AddMessageRequestBody struct {
	UserID *string `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	Msg    *string `form:"msg,omitempty" json:"msg,omitempty" xml:"msg,omitempty"`
}

// AddMessageResponseBody is the type of the "contacts" service "addMessage"
// endpoint HTTP response body.
type AddMessageResponseBody struct {
	Success bool `form:"success" json:"success" xml:"success"`
}

// AddMessageUnknownErrorResponseBody is the type of the "contacts" service
// "addMessage" endpoint HTTP response body for the "unknown_error" error.
type AddMessageUnknownErrorResponseBody struct {
	Err       string `form:"err" json:"err" xml:"err"`
	ErrorCode string `form:"error_code" json:"error_code" xml:"error_code"`
	Success   bool   `form:"success" json:"success" xml:"success"`
}

// NewAddMessageResponseBody builds the HTTP response body from the result of
// the "addMessage" endpoint of the "contacts" service.
func NewAddMessageResponseBody(res *contacts.AddMessageResult) *AddMessageResponseBody {
	body := &AddMessageResponseBody{
		Success: res.Success,
	}
	return body
}

// NewAddMessageUnknownErrorResponseBody builds the HTTP response body from the
// result of the "addMessage" endpoint of the "contacts" service.
func NewAddMessageUnknownErrorResponseBody(res *contacts.UnknownError) *AddMessageUnknownErrorResponseBody {
	body := &AddMessageUnknownErrorResponseBody{
		Err:       res.Err,
		ErrorCode: res.ErrorCode,
		Success:   res.Success,
	}
	return body
}

// NewAddMessagePayload builds a contacts service addMessage endpoint payload.
func NewAddMessagePayload(body *AddMessageRequestBody, oauth *string, jwtToken *string) *contacts.AddMessagePayload {
	v := &contacts.AddMessagePayload{
		UserID: *body.UserID,
		Msg:    *body.Msg,
	}
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v
}

// ValidateAddMessageRequestBody runs the validations defined on
// AddMessageRequestBody
func ValidateAddMessageRequestBody(body *AddMessageRequestBody) (err error) {
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "body"))
	}
	if body.Msg == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("msg", "body"))
	}
	if body.UserID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.user_id", *body.UserID, goa.FormatUUID))
	}
	if body.Msg != nil {
		if utf8.RuneCountInString(*body.Msg) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.msg", *body.Msg, utf8.RuneCountInString(*body.Msg), 2, true))
		}
	}
	if body.Msg != nil {
		if utf8.RuneCountInString(*body.Msg) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.msg", *body.Msg, utf8.RuneCountInString(*body.Msg), 500, false))
		}
	}
	return
}
