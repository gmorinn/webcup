// Code generated by goa v3.7.3, DO NOT EDIT.
//
// data client
//
// Command:
// $ goa gen webcup/design

package data

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "data" service client.
type Client struct {
	ListDataMostRecentEndpoint goa.Endpoint
	CreateDataEndpoint         goa.Endpoint
	UpdateDataEndpoint         goa.Endpoint
	GetDataByUserIDEndpoint    goa.Endpoint
	GetDataByIDEndpoint        goa.Endpoint
}

// NewClient initializes a "data" service client given the endpoints.
func NewClient(listDataMostRecent, createData, updateData, getDataByUserID, getDataByID goa.Endpoint) *Client {
	return &Client{
		ListDataMostRecentEndpoint: listDataMostRecent,
		CreateDataEndpoint:         createData,
		UpdateDataEndpoint:         updateData,
		GetDataByUserIDEndpoint:    getDataByUserID,
		GetDataByIDEndpoint:        getDataByID,
	}
}

// ListDataMostRecent calls the "listDataMostRecent" endpoint of the "data"
// service.
func (c *Client) ListDataMostRecent(ctx context.Context, p *ListDataMostRecentPayload) (res *ListDataMostRecentResult, err error) {
	var ires interface{}
	ires, err = c.ListDataMostRecentEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ListDataMostRecentResult), nil
}

// CreateData calls the "createData" endpoint of the "data" service.
func (c *Client) CreateData(ctx context.Context, p *CreateDataPayload) (res *CreateDataResult, err error) {
	var ires interface{}
	ires, err = c.CreateDataEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CreateDataResult), nil
}

// UpdateData calls the "updateData" endpoint of the "data" service.
func (c *Client) UpdateData(ctx context.Context, p *UpdateDataPayload) (res *UpdateDataResult, err error) {
	var ires interface{}
	ires, err = c.UpdateDataEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UpdateDataResult), nil
}

// GetDataByUserID calls the "getDataByUserID" endpoint of the "data" service.
func (c *Client) GetDataByUserID(ctx context.Context, p *GetDataByUserIDPayload) (res *GetDataByUserIDResult, err error) {
	var ires interface{}
	ires, err = c.GetDataByUserIDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetDataByUserIDResult), nil
}

// GetDataByID calls the "getDataByID" endpoint of the "data" service.
func (c *Client) GetDataByID(ctx context.Context, p *GetDataByIDPayload) (res *GetDataByIDResult, err error) {
	var ires interface{}
	ires, err = c.GetDataByIDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetDataByIDResult), nil
}