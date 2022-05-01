// Code generated by goa v3.7.3, DO NOT EDIT.
//
// bo client
//
// Command:
// $ goa gen webcup/design

package bo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "bo" service client.
type Client struct {
	GetBoUsersEndpoint        goa.Endpoint
	DeleteBoUserEndpoint      goa.Endpoint
	DeleteBoManyUsersEndpoint goa.Endpoint
	UpdateBoUserEndpoint      goa.Endpoint
	GetBoUserEndpoint         goa.Endpoint
}

// NewClient initializes a "bo" service client given the endpoints.
func NewClient(getBoUsers, deleteBoUser, deleteBoManyUsers, updateBoUser, getBoUser goa.Endpoint) *Client {
	return &Client{
		GetBoUsersEndpoint:        getBoUsers,
		DeleteBoUserEndpoint:      deleteBoUser,
		DeleteBoManyUsersEndpoint: deleteBoManyUsers,
		UpdateBoUserEndpoint:      updateBoUser,
		GetBoUserEndpoint:         getBoUser,
	}
}

// GetBoUsers calls the "getBoUsers" endpoint of the "bo" service.
func (c *Client) GetBoUsers(ctx context.Context, p *GetBoUsersPayload) (res *GetBoUsersResult, err error) {
	var ires interface{}
	ires, err = c.GetBoUsersEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetBoUsersResult), nil
}

// DeleteBoUser calls the "deleteBoUser" endpoint of the "bo" service.
func (c *Client) DeleteBoUser(ctx context.Context, p *DeleteBoUserPayload) (res *DeleteBoUserResult, err error) {
	var ires interface{}
	ires, err = c.DeleteBoUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DeleteBoUserResult), nil
}

// DeleteBoManyUsers calls the "deleteBoManyUsers" endpoint of the "bo" service.
func (c *Client) DeleteBoManyUsers(ctx context.Context, p *DeleteBoManyUsersPayload) (res *DeleteBoManyUsersResult, err error) {
	var ires interface{}
	ires, err = c.DeleteBoManyUsersEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DeleteBoManyUsersResult), nil
}

// UpdateBoUser calls the "updateBoUser" endpoint of the "bo" service.
func (c *Client) UpdateBoUser(ctx context.Context, p *UpdateBoUserPayload) (res *UpdateBoUserResult, err error) {
	var ires interface{}
	ires, err = c.UpdateBoUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UpdateBoUserResult), nil
}

// GetBoUser calls the "getBoUser" endpoint of the "bo" service.
func (c *Client) GetBoUser(ctx context.Context, p *GetBoUserPayload) (res *GetBoUserResult, err error) {
	var ires interface{}
	ires, err = c.GetBoUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetBoUserResult), nil
}