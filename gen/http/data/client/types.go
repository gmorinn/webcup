// Code generated by goa v3.7.3, DO NOT EDIT.
//
// data HTTP client types
//
// Command:
// $ goa gen webcup/design

package client

import (
	"unicode/utf8"
	data "webcup/gen/data"

	goa "goa.design/goa/v3/pkg"
)

// ListDataRequestBody is the type of the "data" service "listData" endpoint
// HTTP request body.
type ListDataRequestBody struct {
	Key string `form:"key" json:"key" xml:"key"`
}

// CreateDataRequestBody is the type of the "data" service "createData"
// endpoint HTTP request body.
type CreateDataRequestBody struct {
	Data *PayloadDataRequestBody `form:"data" json:"data" xml:"data"`
}

// UpdateDataRequestBody is the type of the "data" service "updateData"
// endpoint HTTP request body.
type UpdateDataRequestBody struct {
	Data *PayloadDataRequestBody `form:"data" json:"data" xml:"data"`
}

// ListDataResponseBody is the type of the "data" service "listData" endpoint
// HTTP response body.
type ListDataResponseBody struct {
	// Result is an an array of user
	Data    []*ResDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Success *bool                  `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// ListDataMostRecentResponseBody is the type of the "data" service
// "listDataMostRecent" endpoint HTTP response body.
type ListDataMostRecentResponseBody struct {
	// Result is an an array of data
	Data    []*ResDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Success *bool                  `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
	// total of datas
	Count *int64 `form:"count,omitempty" json:"count,omitempty" xml:"count,omitempty"`
}

// CreateDataResponseBody is the type of the "data" service "createData"
// endpoint HTTP response body.
type CreateDataResponseBody struct {
	// Result is an object
	Data    *ResDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Success *bool                `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// UpdateDataResponseBody is the type of the "data" service "updateData"
// endpoint HTTP response body.
type UpdateDataResponseBody struct {
	// Result is an Object
	Data    *ResDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Success *bool                `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// GetDataByUserIDResponseBody is the type of the "data" service
// "getDataByUserID" endpoint HTTP response body.
type GetDataByUserIDResponseBody struct {
	// Result is an object
	Data    []*ResDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Success *bool                  `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// GetDataByIDResponseBody is the type of the "data" service "getDataByID"
// endpoint HTTP response body.
type GetDataByIDResponseBody struct {
	// Result is an object
	Data    *ResDataResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Success *bool                `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// ListDataUnknownErrorResponseBody is the type of the "data" service
// "listData" endpoint HTTP response body for the "unknown_error" error.
type ListDataUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// ListDataMostRecentUnknownErrorResponseBody is the type of the "data" service
// "listDataMostRecent" endpoint HTTP response body for the "unknown_error"
// error.
type ListDataMostRecentUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// CreateDataUnknownErrorResponseBody is the type of the "data" service
// "createData" endpoint HTTP response body for the "unknown_error" error.
type CreateDataUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// UpdateDataUnknownErrorResponseBody is the type of the "data" service
// "updateData" endpoint HTTP response body for the "unknown_error" error.
type UpdateDataUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// GetDataByUserIDUnknownErrorResponseBody is the type of the "data" service
// "getDataByUserID" endpoint HTTP response body for the "unknown_error" error.
type GetDataByUserIDUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// GetDataByIDUnknownErrorResponseBody is the type of the "data" service
// "getDataByID" endpoint HTTP response body for the "unknown_error" error.
type GetDataByIDUnknownErrorResponseBody struct {
	Err       *string `form:"err,omitempty" json:"err,omitempty" xml:"err,omitempty"`
	ErrorCode *string `form:"error_code,omitempty" json:"error_code,omitempty" xml:"error_code,omitempty"`
	Success   *bool   `form:"success,omitempty" json:"success,omitempty" xml:"success,omitempty"`
}

// ResDataResponseBody is used to define fields on response body types.
type ResDataResponseBody struct {
	ID          *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Title       *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Url of the logo and stock in db
	Image    *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	Category *string `form:"category,omitempty" json:"category,omitempty" xml:"category,omitempty"`
	UserID   *string `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// PayloadDataRequestBody is used to define fields on request body types.
type PayloadDataRequestBody struct {
	Title       string `form:"title" json:"title" xml:"title"`
	Description string `form:"description" json:"description" xml:"description"`
	// Url of the logo and stock in db
	Image    *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	Category string  `form:"category" json:"category" xml:"category"`
	UserID   string  `form:"user_id" json:"user_id" xml:"user_id"`
}

// NewListDataRequestBody builds the HTTP request body from the payload of the
// "listData" endpoint of the "data" service.
func NewListDataRequestBody(p *data.ListDataPayload) *ListDataRequestBody {
	body := &ListDataRequestBody{
		Key: p.Key,
	}
	return body
}

// NewCreateDataRequestBody builds the HTTP request body from the payload of
// the "createData" endpoint of the "data" service.
func NewCreateDataRequestBody(p *data.CreateDataPayload) *CreateDataRequestBody {
	body := &CreateDataRequestBody{}
	if p.Data != nil {
		body.Data = marshalDataPayloadDataToPayloadDataRequestBody(p.Data)
	}
	return body
}

// NewUpdateDataRequestBody builds the HTTP request body from the payload of
// the "updateData" endpoint of the "data" service.
func NewUpdateDataRequestBody(p *data.UpdateDataPayload) *UpdateDataRequestBody {
	body := &UpdateDataRequestBody{}
	if p.Data != nil {
		body.Data = marshalDataPayloadDataToPayloadDataRequestBody(p.Data)
	}
	return body
}

// NewListDataResultOK builds a "data" service "listData" endpoint result from
// a HTTP "OK" response.
func NewListDataResultOK(body *ListDataResponseBody) *data.ListDataResult {
	v := &data.ListDataResult{
		Success: *body.Success,
	}
	v.Data = make([]*data.ResData, len(body.Data))
	for i, val := range body.Data {
		v.Data[i] = unmarshalResDataResponseBodyToDataResData(val)
	}

	return v
}

// NewListDataUnknownError builds a data service listData endpoint
// unknown_error error.
func NewListDataUnknownError(body *ListDataUnknownErrorResponseBody) *data.UnknownError {
	v := &data.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewListDataMostRecentResultOK builds a "data" service "listDataMostRecent"
// endpoint result from a HTTP "OK" response.
func NewListDataMostRecentResultOK(body *ListDataMostRecentResponseBody) *data.ListDataMostRecentResult {
	v := &data.ListDataMostRecentResult{
		Success: *body.Success,
		Count:   *body.Count,
	}
	v.Data = make([]*data.ResData, len(body.Data))
	for i, val := range body.Data {
		v.Data[i] = unmarshalResDataResponseBodyToDataResData(val)
	}

	return v
}

// NewListDataMostRecentUnknownError builds a data service listDataMostRecent
// endpoint unknown_error error.
func NewListDataMostRecentUnknownError(body *ListDataMostRecentUnknownErrorResponseBody) *data.UnknownError {
	v := &data.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewCreateDataResultCreated builds a "data" service "createData" endpoint
// result from a HTTP "Created" response.
func NewCreateDataResultCreated(body *CreateDataResponseBody) *data.CreateDataResult {
	v := &data.CreateDataResult{
		Success: *body.Success,
	}
	v.Data = unmarshalResDataResponseBodyToDataResData(body.Data)

	return v
}

// NewCreateDataUnknownError builds a data service createData endpoint
// unknown_error error.
func NewCreateDataUnknownError(body *CreateDataUnknownErrorResponseBody) *data.UnknownError {
	v := &data.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewUpdateDataResultOK builds a "data" service "updateData" endpoint result
// from a HTTP "OK" response.
func NewUpdateDataResultOK(body *UpdateDataResponseBody) *data.UpdateDataResult {
	v := &data.UpdateDataResult{
		Success: *body.Success,
	}
	v.Data = unmarshalResDataResponseBodyToDataResData(body.Data)

	return v
}

// NewUpdateDataUnknownError builds a data service updateData endpoint
// unknown_error error.
func NewUpdateDataUnknownError(body *UpdateDataUnknownErrorResponseBody) *data.UnknownError {
	v := &data.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewGetDataByUserIDResultOK builds a "data" service "getDataByUserID"
// endpoint result from a HTTP "OK" response.
func NewGetDataByUserIDResultOK(body *GetDataByUserIDResponseBody) *data.GetDataByUserIDResult {
	v := &data.GetDataByUserIDResult{
		Success: *body.Success,
	}
	v.Data = make([]*data.ResData, len(body.Data))
	for i, val := range body.Data {
		v.Data[i] = unmarshalResDataResponseBodyToDataResData(val)
	}

	return v
}

// NewGetDataByUserIDUnknownError builds a data service getDataByUserID
// endpoint unknown_error error.
func NewGetDataByUserIDUnknownError(body *GetDataByUserIDUnknownErrorResponseBody) *data.UnknownError {
	v := &data.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// NewGetDataByIDResultOK builds a "data" service "getDataByID" endpoint result
// from a HTTP "OK" response.
func NewGetDataByIDResultOK(body *GetDataByIDResponseBody) *data.GetDataByIDResult {
	v := &data.GetDataByIDResult{
		Success: *body.Success,
	}
	v.Data = unmarshalResDataResponseBodyToDataResData(body.Data)

	return v
}

// NewGetDataByIDUnknownError builds a data service getDataByID endpoint
// unknown_error error.
func NewGetDataByIDUnknownError(body *GetDataByIDUnknownErrorResponseBody) *data.UnknownError {
	v := &data.UnknownError{
		Err:       *body.Err,
		ErrorCode: *body.ErrorCode,
		Success:   *body.Success,
	}

	return v
}

// ValidateListDataResponseBody runs the validations defined on
// ListDataResponseBody
func ValidateListDataResponseBody(body *ListDataResponseBody) (err error) {
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	for _, e := range body.Data {
		if e != nil {
			if err2 := ValidateResDataResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateListDataMostRecentResponseBody runs the validations defined on
// ListDataMostRecentResponseBody
func ValidateListDataMostRecentResponseBody(body *ListDataMostRecentResponseBody) (err error) {
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.Count == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("count", "body"))
	}
	for _, e := range body.Data {
		if e != nil {
			if err2 := ValidateResDataResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateCreateDataResponseBody runs the validations defined on
// CreateDataResponseBody
func ValidateCreateDataResponseBody(body *CreateDataResponseBody) (err error) {
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.Data != nil {
		if err2 := ValidateResDataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateDataResponseBody runs the validations defined on
// UpdateDataResponseBody
func ValidateUpdateDataResponseBody(body *UpdateDataResponseBody) (err error) {
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.Data != nil {
		if err2 := ValidateResDataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateGetDataByUserIDResponseBody runs the validations defined on
// GetDataByUserIDResponseBody
func ValidateGetDataByUserIDResponseBody(body *GetDataByUserIDResponseBody) (err error) {
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	for _, e := range body.Data {
		if e != nil {
			if err2 := ValidateResDataResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateGetDataByIDResponseBody runs the validations defined on
// GetDataByIDResponseBody
func ValidateGetDataByIDResponseBody(body *GetDataByIDResponseBody) (err error) {
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.Data != nil {
		if err2 := ValidateResDataResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateListDataUnknownErrorResponseBody runs the validations defined on
// listData_unknown_error_response_body
func ValidateListDataUnknownErrorResponseBody(body *ListDataUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}

// ValidateListDataMostRecentUnknownErrorResponseBody runs the validations
// defined on listDataMostRecent_unknown_error_response_body
func ValidateListDataMostRecentUnknownErrorResponseBody(body *ListDataMostRecentUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}

// ValidateCreateDataUnknownErrorResponseBody runs the validations defined on
// createData_unknown_error_response_body
func ValidateCreateDataUnknownErrorResponseBody(body *CreateDataUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}

// ValidateUpdateDataUnknownErrorResponseBody runs the validations defined on
// updateData_unknown_error_response_body
func ValidateUpdateDataUnknownErrorResponseBody(body *UpdateDataUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}

// ValidateGetDataByUserIDUnknownErrorResponseBody runs the validations defined
// on getDataByUserID_unknown_error_response_body
func ValidateGetDataByUserIDUnknownErrorResponseBody(body *GetDataByUserIDUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}

// ValidateGetDataByIDUnknownErrorResponseBody runs the validations defined on
// getDataByID_unknown_error_response_body
func ValidateGetDataByIDUnknownErrorResponseBody(body *GetDataByIDUnknownErrorResponseBody) (err error) {
	if body.Err == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("err", "body"))
	}
	if body.Success == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("success", "body"))
	}
	if body.ErrorCode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("error_code", "body"))
	}
	return
}

// ValidateResDataResponseBody runs the validations defined on
// resDataResponseBody
func ValidateResDataResponseBody(body *ResDataResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Image == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image", "body"))
	}
	if body.Category == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("category", "body"))
	}
	if body.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "body"))
	}
	if body.ID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.id", *body.ID, goa.FormatUUID))
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", *body.Title, utf8.RuneCountInString(*body.Title), 3, true))
		}
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", *body.Title, utf8.RuneCountInString(*body.Title), 20, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2, true))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 500, false))
		}
	}
	if body.Category != nil {
		if !(*body.Category == "robotics" || *body.Category == "space" || *body.Category == "brain" || *body.Category == "animals") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.category", *body.Category, []interface{}{"robotics", "space", "brain", "animals"}))
		}
	}
	if body.UserID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.user_id", *body.UserID, goa.FormatUUID))
	}
	return
}

// ValidatePayloadDataRequestBody runs the validations defined on
// payloadDataRequestBody
func ValidatePayloadDataRequestBody(body *PayloadDataRequestBody) (err error) {
	if utf8.RuneCountInString(body.Title) < 3 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", body.Title, utf8.RuneCountInString(body.Title), 3, true))
	}
	if utf8.RuneCountInString(body.Title) > 20 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", body.Title, utf8.RuneCountInString(body.Title), 20, false))
	}
	if utf8.RuneCountInString(body.Description) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", body.Description, utf8.RuneCountInString(body.Description), 2, true))
	}
	if utf8.RuneCountInString(body.Description) > 500 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", body.Description, utf8.RuneCountInString(body.Description), 500, false))
	}
	if !(body.Category == "robotics" || body.Category == "space" || body.Category == "brain" || body.Category == "animals") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.category", body.Category, []interface{}{"robotics", "space", "brain", "animals"}))
	}
	err = goa.MergeErrors(err, goa.ValidateFormat("body.user_id", body.UserID, goa.FormatUUID))

	return
}
