// Code generated by goa v3.7.3, DO NOT EDIT.
//
// jwtToken HTTP client encoders and decoders
//
// Command:
// $ goa gen webcup/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	jwttoken "webcup/gen/jwt_token"

	goahttp "goa.design/goa/v3/http"
)

// BuildSignupRequest instantiates a HTTP request object with method and path
// set to call the "jwtToken" service "signup" endpoint
func (c *Client) BuildSignupRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SignupJWTTokenPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("jwtToken", "signup", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSignupRequest returns an encoder for requests sent to the jwtToken
// signup server.
func EncodeSignupRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*jwttoken.SignupPayload)
		if !ok {
			return goahttp.ErrInvalidType("jwtToken", "signup", "*jwttoken.SignupPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewSignupRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("jwtToken", "signup", err)
		}
		return nil
	}
}

// DecodeSignupResponse returns a decoder for responses returned by the
// jwtToken signup endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeSignupResponse may return the following errors:
//	- "unknown_error" (type *jwttoken.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeSignupResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SignupResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("jwtToken", "signup", err)
			}
			err = ValidateSignupResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("jwtToken", "signup", err)
			}
			res := NewSignupSignOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body SignupUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("jwtToken", "signup", err)
			}
			err = ValidateSignupUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("jwtToken", "signup", err)
			}
			return nil, NewSignupUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("jwtToken", "signup", resp.StatusCode, string(body))
		}
	}
}

// BuildSigninRequest instantiates a HTTP request object with method and path
// set to call the "jwtToken" service "signin" endpoint
func (c *Client) BuildSigninRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SigninJWTTokenPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("jwtToken", "signin", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSigninRequest returns an encoder for requests sent to the jwtToken
// signin server.
func EncodeSigninRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*jwttoken.SigninPayload)
		if !ok {
			return goahttp.ErrInvalidType("jwtToken", "signin", "*jwttoken.SigninPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewSigninRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("jwtToken", "signin", err)
		}
		return nil
	}
}

// DecodeSigninResponse returns a decoder for responses returned by the
// jwtToken signin endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeSigninResponse may return the following errors:
//	- "unknown_error" (type *jwttoken.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeSigninResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SigninResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("jwtToken", "signin", err)
			}
			err = ValidateSigninResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("jwtToken", "signin", err)
			}
			res := NewSigninSignOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body SigninUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("jwtToken", "signin", err)
			}
			err = ValidateSigninUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("jwtToken", "signin", err)
			}
			return nil, NewSigninUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("jwtToken", "signin", resp.StatusCode, string(body))
		}
	}
}

// BuildRefreshRequest instantiates a HTTP request object with method and path
// set to call the "jwtToken" service "refresh" endpoint
func (c *Client) BuildRefreshRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RefreshJWTTokenPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("jwtToken", "refresh", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRefreshRequest returns an encoder for requests sent to the jwtToken
// refresh server.
func EncodeRefreshRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*jwttoken.RefreshPayload)
		if !ok {
			return goahttp.ErrInvalidType("jwtToken", "refresh", "*jwttoken.RefreshPayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewRefreshRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("jwtToken", "refresh", err)
		}
		return nil
	}
}

// DecodeRefreshResponse returns a decoder for responses returned by the
// jwtToken refresh endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeRefreshResponse may return the following errors:
//	- "unknown_error" (type *jwttoken.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeRefreshResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RefreshResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("jwtToken", "refresh", err)
			}
			err = ValidateRefreshResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("jwtToken", "refresh", err)
			}
			res := NewRefreshSignOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body RefreshUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("jwtToken", "refresh", err)
			}
			err = ValidateRefreshUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("jwtToken", "refresh", err)
			}
			return nil, NewRefreshUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("jwtToken", "refresh", resp.StatusCode, string(body))
		}
	}
}