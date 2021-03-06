// Code generated by goa v3.7.3, DO NOT EDIT.
//
// jwtToken client
//
// Command:
// $ goa gen webcup/design

package jwttoken

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "jwtToken" service client.
type Client struct {
	SignupEndpoint   goa.Endpoint
	SigninEndpoint   goa.Endpoint
	RefreshEndpoint  goa.Endpoint
	SigninBoEndpoint goa.Endpoint
}

// NewClient initializes a "jwtToken" service client given the endpoints.
func NewClient(signup, signin, refresh, signinBo goa.Endpoint) *Client {
	return &Client{
		SignupEndpoint:   signup,
		SigninEndpoint:   signin,
		RefreshEndpoint:  refresh,
		SigninBoEndpoint: signinBo,
	}
}

// Signup calls the "signup" endpoint of the "jwtToken" service.
func (c *Client) Signup(ctx context.Context, p *SignupPayload) (res *Sign, err error) {
	var ires interface{}
	ires, err = c.SignupEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Sign), nil
}

// Signin calls the "signin" endpoint of the "jwtToken" service.
func (c *Client) Signin(ctx context.Context, p *SigninPayload) (res *Sign, err error) {
	var ires interface{}
	ires, err = c.SigninEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Sign), nil
}

// Refresh calls the "refresh" endpoint of the "jwtToken" service.
func (c *Client) Refresh(ctx context.Context, p *RefreshPayload) (res *Sign, err error) {
	var ires interface{}
	ires, err = c.RefreshEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Sign), nil
}

// SigninBo calls the "signinBo" endpoint of the "jwtToken" service.
func (c *Client) SigninBo(ctx context.Context, p *SigninBoPayload) (res *Sign, err error) {
	var ires interface{}
	ires, err = c.SigninBoEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Sign), nil
}
