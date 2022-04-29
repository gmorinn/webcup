// Code generated by goa v3.7.2, DO NOT EDIT.
//
// contacts service
//
// Command:
// $ goa gen webcup/back/design

package contacts

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// contacts of the api
type Service interface {
	// user ask for something
	AddMessage(context.Context, *AddMessagePayload) (res *AddMessageResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// OAuth2Auth implements the authorization logic for the OAuth2 security scheme.
	OAuth2Auth(ctx context.Context, token string, schema *security.OAuth2Scheme) (context.Context, error)
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "contacts"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"addMessage"}

// AddMessagePayload is the payload type of the contacts service addMessage
// method.
type AddMessagePayload struct {
	UserID string
	Msg    string
	// JWT used for authentication after Signin/Signup
	JWTToken *string
	// Use to generate Oauth with /authorization
	Oauth *string
}

// AddMessageResult is the result type of the contacts service addMessage
// method.
type AddMessageResult struct {
	Success bool
}

type UnknownError struct {
	Err       string
	ErrorCode string
	Success   bool
}

// Error returns an error description.
func (e *UnknownError) Error() string {
	return ""
}

// ErrorName returns "unknownError".
func (e *UnknownError) ErrorName() string {
	return "unknown_error"
}

// MakeTimeout builds a goa.ServiceError from an error.
func MakeTimeout(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "timeout",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
		Timeout: true,
	}
}