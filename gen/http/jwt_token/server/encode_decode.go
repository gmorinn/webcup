// Code generated by goa v3.7.3, DO NOT EDIT.
//
// jwtToken HTTP server encoders and decoders
//
// Command:
// $ goa gen webcup/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	jwttoken "webcup/gen/jwt_token"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeSignupResponse returns an encoder for responses returned by the
// jwtToken signup endpoint.
func EncodeSignupResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*jwttoken.Sign)
		enc := encoder(ctx, w)
		body := NewSignupResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSignupRequest returns a decoder for requests sent to the jwtToken
// signup endpoint.
func DecodeSignupRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body SignupRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSignupRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			oauth *string
		)
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		payload := NewSignupPayload(&body, oauth)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}

		return payload, nil
	}
}

// EncodeSignupError returns an encoder for errors returned by the signup
// jwtToken endpoint.
func EncodeSignupError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *jwttoken.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSignupUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeSigninResponse returns an encoder for responses returned by the
// jwtToken signin endpoint.
func EncodeSigninResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*jwttoken.Sign)
		enc := encoder(ctx, w)
		body := NewSigninResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSigninRequest returns a decoder for requests sent to the jwtToken
// signin endpoint.
func DecodeSigninRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body SigninRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSigninRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			oauth *string
		)
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		payload := NewSigninPayload(&body, oauth)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}

		return payload, nil
	}
}

// EncodeSigninError returns an encoder for errors returned by the signin
// jwtToken endpoint.
func EncodeSigninError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *jwttoken.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSigninUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRefreshResponse returns an encoder for responses returned by the
// jwtToken refresh endpoint.
func EncodeRefreshResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*jwttoken.Sign)
		enc := encoder(ctx, w)
		body := NewRefreshResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRefreshRequest returns a decoder for requests sent to the jwtToken
// refresh endpoint.
func DecodeRefreshRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body RefreshRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateRefreshRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			oauth *string
		)
		oauthRaw := r.Header.Get("Authorization")
		if oauthRaw != "" {
			oauth = &oauthRaw
		}
		payload := NewRefreshPayload(&body, oauth)
		if payload.Oauth != nil {
			if strings.Contains(*payload.Oauth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Oauth, " ", 2)[1]
				payload.Oauth = &cred
			}
		}

		return payload, nil
	}
}

// EncodeRefreshError returns an encoder for errors returned by the refresh
// jwtToken endpoint.
func EncodeRefreshError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unknown_error":
			var res *jwttoken.UnknownError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRefreshUnknownErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}