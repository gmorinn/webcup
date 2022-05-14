// Code generated by goa v3.7.3, DO NOT EDIT.
//
// data service
//
// Command:
// $ goa gen webcup/design

package data

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// futristics data of the api
type Service interface {
	// List data the most recent
	ListDataMostRecent(context.Context, *ListDataMostRecentPayload) (res *ListDataMostRecentResult, err error)
	// Create one data
	CreateData(context.Context, *CreateDataPayload) (res *CreateDataResult, err error)
	// Update one data
	UpdateData(context.Context, *UpdateDataPayload) (res *UpdateDataResult, err error)
	// Get one data user id
	GetDataByUserID(context.Context, *GetDataByUserIDPayload) (res *GetDataByUserIDResult, err error)
	// Get one data by id
	GetDataByID(context.Context, *GetDataByIDPayload) (res *GetDataByIDResult, err error)
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
const ServiceName = "data"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"listDataMostRecent", "createData", "updateData", "getDataByUserID", "getDataByID"}

// CreateDataPayload is the payload type of the data service createData method.
type CreateDataPayload struct {
	Data *PayloadData
	// JWT used for authentication after Signin/Signup
	JWTToken *string
	// Use to generate Oauth with /authorization
	Oauth *string
}

// CreateDataResult is the result type of the data service createData method.
type CreateDataResult struct {
	// Result is an object
	Data    *ResData
	Success bool
}

// GetDataByIDPayload is the payload type of the data service getDataByID
// method.
type GetDataByIDPayload struct {
	// Unique ID of the data
	ID string
	// JWT used for authentication after Signin/Signup
	JWTToken *string
	// Use to generate Oauth with /authorization
	Oauth *string
}

// GetDataByIDResult is the result type of the data service getDataByID method.
type GetDataByIDResult struct {
	// Result is an object
	Data    *ResData
	Success bool
}

// GetDataByUserIDPayload is the payload type of the data service
// getDataByUserID method.
type GetDataByUserIDPayload struct {
	// Unique ID of the user
	UserID string
	// JWT used for authentication after Signin/Signup
	JWTToken *string
	// Use to generate Oauth with /authorization
	Oauth *string
}

// GetDataByUserIDResult is the result type of the data service getDataByUserID
// method.
type GetDataByUserIDResult struct {
	// Result is an object
	Data    []*ResData
	Success bool
}

// ListDataMostRecentPayload is the payload type of the data service
// listDataMostRecent method.
type ListDataMostRecentPayload struct {
	// JWT used for authentication after Signin/Signup
	JWTToken *string
	// Use to generate Oauth with /authorization
	Oauth *string
	// Offset for pagination
	Offset int32
	// Limit of items listed for pagination
	Limit int32
}

// ListDataMostRecentResult is the result type of the data service
// listDataMostRecent method.
type ListDataMostRecentResult struct {
	// Result is an an array of data
	Data    []*ResData
	Success bool
	// total of datas
	Count int64
}

type PayloadData struct {
	Title       string
	Description string
	// Url of the logo and stock in db
	Image    string
	Category string
	UserID   string
}

type ResData struct {
	ID          string
	Title       string
	Description string
	// Url of the logo and stock in db
	Image    string
	Category string
	UserID   string
}

// UpdateDataPayload is the payload type of the data service updateData method.
type UpdateDataPayload struct {
	ID   string
	Data *PayloadData
	// JWT used for authentication after Signin/Signup
	JWTToken *string
	// Use to generate Oauth with /authorization
	Oauth *string
}

// UpdateDataResult is the result type of the data service updateData method.
type UpdateDataResult struct {
	// Result is an Object
	Data    *ResData
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