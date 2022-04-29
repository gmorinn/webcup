// Code generated by goa v3.7.2, DO NOT EDIT.
//
// users client HTTP transport
//
// Command:
// $ goa gen webcup/back/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the users service endpoint HTTP clients.
type Client struct {
	// DeleteUser Doer is the HTTP client used to make requests to the deleteUser
	// endpoint.
	DeleteUserDoer goahttp.Doer

	// GetUserByID Doer is the HTTP client used to make requests to the getUserByID
	// endpoint.
	GetUserByIDDoer goahttp.Doer

	// UpdateDescription Doer is the HTTP client used to make requests to the
	// updateDescription endpoint.
	UpdateDescriptionDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the users service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		DeleteUserDoer:        doer,
		GetUserByIDDoer:       doer,
		UpdateDescriptionDoer: doer,
		CORSDoer:              doer,
		RestoreResponseBody:   restoreBody,
		scheme:                scheme,
		host:                  host,
		decoder:               dec,
		encoder:               enc,
	}
}

// DeleteUser returns an endpoint that makes HTTP requests to the users service
// deleteUser server.
func (c *Client) DeleteUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteUserRequest(c.encoder)
		decodeResponse = DecodeDeleteUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteUserDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("users", "deleteUser", err)
		}
		return decodeResponse(resp)
	}
}

// GetUserByID returns an endpoint that makes HTTP requests to the users
// service getUserByID server.
func (c *Client) GetUserByID() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetUserByIDRequest(c.encoder)
		decodeResponse = DecodeGetUserByIDResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetUserByIDRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetUserByIDDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("users", "getUserByID", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateDescription returns an endpoint that makes HTTP requests to the users
// service updateDescription server.
func (c *Client) UpdateDescription() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateDescriptionRequest(c.encoder)
		decodeResponse = DecodeUpdateDescriptionResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateDescriptionRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateDescriptionDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("users", "updateDescription", err)
		}
		return decodeResponse(resp)
	}
}
