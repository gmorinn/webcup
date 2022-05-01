// Code generated by goa v3.7.3, DO NOT EDIT.
//
// bo HTTP server encoders and decoders
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
	bo "webcup/gen/bo"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetBoUsersResponse returns an encoder for responses returned by the bo
// getBoUsers endpoint.
func EncodeGetBoUsersResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bo.GetBoUsersResult)
		enc := encoder(ctx, w)
		body := NewGetBoUsersResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetBoUsersRequest returns a decoder for requests sent to the bo
// getBoUsers endpoint.
func DecodeGetBoUsersRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
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
		payload := NewGetBoUsersPayload(offset, limit, field, direction, oauth, jwtToken)
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

// EncodeGetBoUsersError returns an encoder for errors returned by the
// getBoUsers bo endpoint.
func EncodeGetBoUsersError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *bo.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetBoUsersUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteBoUserResponse returns an encoder for responses returned by the
// bo deleteBoUser endpoint.
func EncodeDeleteBoUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bo.DeleteBoUserResult)
		enc := encoder(ctx, w)
		body := NewDeleteBoUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteBoUserRequest returns a decoder for requests sent to the bo
// deleteBoUser endpoint.
func DecodeDeleteBoUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
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
		payload := NewDeleteBoUserPayload(id, oauth, jwtToken)
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

// EncodeDeleteBoUserError returns an encoder for errors returned by the
// deleteBoUser bo endpoint.
func EncodeDeleteBoUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *bo.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteBoUserUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteBoManyUsersResponse returns an encoder for responses returned by
// the bo deleteBoManyUsers endpoint.
func EncodeDeleteBoManyUsersResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bo.DeleteBoManyUsersResult)
		enc := encoder(ctx, w)
		body := NewDeleteBoManyUsersResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteBoManyUsersRequest returns a decoder for requests sent to the bo
// deleteBoManyUsers endpoint.
func DecodeDeleteBoManyUsersRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body DeleteBoManyUsersRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateDeleteBoManyUsersRequestBody(&body)
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
		payload := NewDeleteBoManyUsersPayload(&body, oauth, jwtToken)
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

// EncodeDeleteBoManyUsersError returns an encoder for errors returned by the
// deleteBoManyUsers bo endpoint.
func EncodeDeleteBoManyUsersError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *bo.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteBoManyUsersUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateBoUserResponse returns an encoder for responses returned by the
// bo updateBoUser endpoint.
func EncodeUpdateBoUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bo.UpdateBoUserResult)
		enc := encoder(ctx, w)
		body := NewUpdateBoUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateBoUserRequest returns a decoder for requests sent to the bo
// updateBoUser endpoint.
func DecodeUpdateBoUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateBoUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateBoUserRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id       string
			oauth    *string
			jwtToken *string

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
		payload := NewUpdateBoUserPayload(&body, id, oauth, jwtToken)
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

// EncodeUpdateBoUserError returns an encoder for errors returned by the
// updateBoUser bo endpoint.
func EncodeUpdateBoUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *bo.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateBoUserUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetBoUserResponse returns an encoder for responses returned by the bo
// getBoUser endpoint.
func EncodeGetBoUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*bo.GetBoUserResult)
		enc := encoder(ctx, w)
		body := NewGetBoUserResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetBoUserRequest returns a decoder for requests sent to the bo
// getBoUser endpoint.
func DecodeGetBoUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
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
		payload := NewGetBoUserPayload(id, oauth, jwtToken)
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

// EncodeGetBoUserError returns an encoder for errors returned by the getBoUser
// bo endpoint.
func EncodeGetBoUserError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *bo.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetBoUserUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalBoResUserToResUserResponseBody builds a value of type
// *ResUserResponseBody from a value of type *bo.ResUser.
func marshalBoResUserToResUserResponseBody(v *bo.ResUser) *ResUserResponseBody {
	res := &ResUserResponseBody{
		ID:        v.ID,
		Firstname: v.Firstname,
		Lastname:  v.Lastname,
		Username:  v.Username,
		Email:     v.Email,
		Role:      v.Role,
		Avatar:    v.Avatar,
	}

	return res
}

// unmarshalPayloadUserRequestBodyToBoPayloadUser builds a value of type
// *bo.PayloadUser from a value of type *PayloadUserRequestBody.
func unmarshalPayloadUserRequestBodyToBoPayloadUser(v *PayloadUserRequestBody) *bo.PayloadUser {
	res := &bo.PayloadUser{
		Email:     *v.Email,
		Username:  *v.Username,
		Firstname: *v.Firstname,
		Lastname:  *v.Lastname,
		Avatar:    *v.Avatar,
		Role:      *v.Role,
	}

	return res
}