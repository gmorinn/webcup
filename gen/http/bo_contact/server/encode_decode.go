// Code generated by goa v3.7.3, DO NOT EDIT.
//
// boContact HTTP server encoders and decoders
//
// Command:
// $ goa gen webcup/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	bocontact "webcup/gen/bo_contact"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetBoContactResponse returns an encoder for responses returned by the
// boContact getBoContact endpoint.
func EncodeGetBoContactResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bocontact.GetBoContactResult)
		enc := encoder(ctx, w)
		body := NewGetBoContactResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetBoContactRequest returns a decoder for requests sent to the
// boContact getBoContact endpoint.
func DecodeGetBoContactRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			offset    int32
			limit     int32
			field     string
			direction string
			oauth     *string
			jwtToken  *string
			err       error

			params = mux.Vars(r)
		)
		{
			offsetRaw := params["offset"]
			v, err2 := strconv.ParseInt(offsetRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("offset", offsetRaw, "integer"))
			}
			offset = int32(v)
		}
		{
			limitRaw := params["limit"]
			v, err2 := strconv.ParseInt(limitRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "integer"))
			}
			limit = int32(v)
		}
		fieldRaw := r.URL.Query().Get("field")
		if fieldRaw != "" {
			field = fieldRaw
		} else {
			field = "name"
		}
		directionRaw := r.URL.Query().Get("direction")
		if directionRaw != "" {
			direction = directionRaw
		} else {
			direction = "asc"
		}
		if !(direction == "asc" || direction == "desc") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("direction", direction, []interface{}{"asc", "desc"}))
		}
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetBoContactPayload(offset, limit, field, direction, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetBoContactError returns an encoder for errors returned by the
// getBoContact boContact endpoint.
func EncodeGetBoContactError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *bocontact.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetBoContactUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteBoContactResponse returns an encoder for responses returned by
// the boContact deleteBoContact endpoint.
func EncodeDeleteBoContactResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bocontact.DeleteBoContactResult)
		enc := encoder(ctx, w)
		body := NewDeleteBoContactResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteBoContactRequest returns a decoder for requests sent to the
// boContact deleteBoContact endpoint.
func DecodeDeleteBoContactRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id       string
			oauth    *string
			jwtToken *string
			err      error

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteBoContactPayload(id, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDeleteBoContactError returns an encoder for errors returned by the
// deleteBoContact boContact endpoint.
func EncodeDeleteBoContactError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *bocontact.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteBoContactUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetBoContactByIDResponse returns an encoder for responses returned by
// the boContact getBoContactByID endpoint.
func EncodeGetBoContactByIDResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bocontact.GetBoContactByIDResult)
		enc := encoder(ctx, w)
		body := NewGetBoContactByIDResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetBoContactByIDRequest returns a decoder for requests sent to the
// boContact getBoContactByID endpoint.
func DecodeGetBoContactByIDRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id       string
			oauth    *string
			jwtToken *string
			err      error

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetBoContactByIDPayload(id, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetBoContactByIDError returns an encoder for errors returned by the
// getBoContactByID boContact endpoint.
func EncodeGetBoContactByIDError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *bocontact.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetBoContactByIDUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteBoManyContactResponse returns an encoder for responses returned
// by the boContact deleteBoManyContact endpoint.
func EncodeDeleteBoManyContactResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bocontact.DeleteBoManyContactResult)
		enc := encoder(ctx, w)
		body := NewDeleteBoManyContactResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteBoManyContactRequest returns a decoder for requests sent to the
// boContact deleteBoManyContact endpoint.
func DecodeDeleteBoManyContactRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body DeleteBoManyContactRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateDeleteBoManyContactRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			oauth    *string
			jwtToken *string
		)
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		jwtTokenRaw := r.Header.Get("jwtToken")
		if jwtTokenRaw != "" {
			jwtToken = &jwtTokenRaw
		}
		payload := NewDeleteBoManyContactPayload(&body, oauth, jwtToken)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}
		if payload.JWTToken != nil {
			if strings.Contains(*payload.JWTToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.JWTToken, " ", 2)[1]
				payload.JWTToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDeleteBoManyContactError returns an encoder for errors returned by the
// deleteBoManyContact boContact endpoint.
func EncodeDeleteBoManyContactError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *bocontact.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteBoManyContactUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalBocontactResContactToResContactResponseBody builds a value of type
// *ResContactResponseBody from a value of type *bocontact.ResContact.
func marshalBocontactResContactToResContactResponseBody(v *bocontact.ResContact) *ResContactResponseBody {
	res := &ResContactResponseBody{
		ID:      v.ID,
		UserID:  v.UserID,
		Email:   v.Email,
		Message: v.Message,
	}

	return res
}
