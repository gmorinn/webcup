// Code generated by goa v3.7.3, DO NOT EDIT.
//
// boContact HTTP client CLI support package
//
// Command:
// $ goa gen webcup/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	bocontact "webcup/gen/bo_contact"

	goa "goa.design/goa/v3/pkg"
)

// BuildGetBoContactPayload builds the payload for the boContact getBoContact
// endpoint from CLI flags.
func BuildGetBoContactPayload(boContactGetBoContactOffset string, boContactGetBoContactLimit string, boContactGetBoContactField string, boContactGetBoContactDirection string, boContactGetBoContactOauth string, boContactGetBoContactJWTToken string) (*bocontact.GetBoContactPayload, error) {
	var err error
	var offset int32
	{
		var v int64
		v, err = strconv.ParseInt(boContactGetBoContactOffset, 10, 32)
		offset = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for offset, must be INT32")
		}
	}
	var limit int32
	{
		var v int64
		v, err = strconv.ParseInt(boContactGetBoContactLimit, 10, 32)
		limit = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for limit, must be INT32")
		}
	}
	var field string
	{
		if boContactGetBoContactField != "" {
			field = boContactGetBoContactField
		}
	}
	var direction string
	{
		if boContactGetBoContactDirection != "" {
			direction = boContactGetBoContactDirection
			if !(direction == "asc" || direction == "desc") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("direction", direction, []interface{}{"asc", "desc"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var oauth *string
	{
		if boContactGetBoContactOauth != "" {
			oauth = &boContactGetBoContactOauth
		}
	}
	var jwtToken *string
	{
		if boContactGetBoContactJWTToken != "" {
			jwtToken = &boContactGetBoContactJWTToken
		}
	}
	v := &bocontact.GetBoContactPayload{}
	v.Offset = offset
	v.Limit = limit
	v.Field = field
	v.Direction = direction
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}

// BuildDeleteBoContactPayload builds the payload for the boContact
// deleteBoContact endpoint from CLI flags.
func BuildDeleteBoContactPayload(boContactDeleteBoContactID string, boContactDeleteBoContactOauth string, boContactDeleteBoContactJWTToken string) (*bocontact.DeleteBoContactPayload, error) {
	var err error
	var id string
	{
		id = boContactDeleteBoContactID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if boContactDeleteBoContactOauth != "" {
			oauth = &boContactDeleteBoContactOauth
		}
	}
	var jwtToken *string
	{
		if boContactDeleteBoContactJWTToken != "" {
			jwtToken = &boContactDeleteBoContactJWTToken
		}
	}
	v := &bocontact.DeleteBoContactPayload{}
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}

// BuildGetBoContactByIDPayload builds the payload for the boContact
// getBoContactByID endpoint from CLI flags.
func BuildGetBoContactByIDPayload(boContactGetBoContactByIDID string, boContactGetBoContactByIDOauth string, boContactGetBoContactByIDJWTToken string) (*bocontact.GetBoContactByIDPayload, error) {
	var err error
	var id string
	{
		id = boContactGetBoContactByIDID
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if boContactGetBoContactByIDOauth != "" {
			oauth = &boContactGetBoContactByIDOauth
		}
	}
	var jwtToken *string
	{
		if boContactGetBoContactByIDJWTToken != "" {
			jwtToken = &boContactGetBoContactByIDJWTToken
		}
	}
	v := &bocontact.GetBoContactByIDPayload{}
	v.ID = id
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}

// BuildDeleteBoManyContactPayload builds the payload for the boContact
// deleteBoManyContact endpoint from CLI flags.
func BuildDeleteBoManyContactPayload(boContactDeleteBoManyContactBody string, boContactDeleteBoManyContactOauth string, boContactDeleteBoManyContactJWTToken string) (*bocontact.DeleteBoManyContactPayload, error) {
	var err error
	var body DeleteBoManyContactRequestBody
	{
		err = json.Unmarshal([]byte(boContactDeleteBoManyContactBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"tab\": [\n         \"Adipisci harum ipsum.\",\n         \"In beatae fugiat unde eos possimus mollitia.\"\n      ]\n   }'")
		}
		if body.Tab == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("tab", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if boContactDeleteBoManyContactOauth != "" {
			oauth = &boContactDeleteBoManyContactOauth
		}
	}
	var jwtToken *string
	{
		if boContactDeleteBoManyContactJWTToken != "" {
			jwtToken = &boContactDeleteBoManyContactJWTToken
		}
	}
	v := &bocontact.DeleteBoManyContactPayload{}
	if body.Tab != nil {
		v.Tab = make([]string, len(body.Tab))
		for i, val := range body.Tab {
			v.Tab[i] = val
		}
	}
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}
