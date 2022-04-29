// Code generated by goa v3.7.2, DO NOT EDIT.
//
// contacts HTTP client CLI support package
//
// Command:
// $ goa gen webcup/back/design

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
	contacts "webcup/back/gen/contacts"

	goa "goa.design/goa/v3/pkg"
)

// BuildAddMessagePayload builds the payload for the contacts addMessage
// endpoint from CLI flags.
func BuildAddMessagePayload(contactsAddMessageBody string, contactsAddMessageOauth string, contactsAddMessageJWTToken string) (*contacts.AddMessagePayload, error) {
	var err error
	var body AddMessageRequestBody
	{
		err = json.Unmarshal([]byte(contactsAddMessageBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"msg\": \"Je reprends l\\'app pour un million\",\n      \"user_id\": \"5dfb0bf7-597a-4250-b7ad-63a43ff59c25\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.user_id", body.UserID, goa.FormatUUID))

		if utf8.RuneCountInString(body.Msg) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.msg", body.Msg, utf8.RuneCountInString(body.Msg), 2, true))
		}
		if utf8.RuneCountInString(body.Msg) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.msg", body.Msg, utf8.RuneCountInString(body.Msg), 500, false))
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if contactsAddMessageOauth != "" {
			oauth = &contactsAddMessageOauth
		}
	}
	var jwtToken *string
	{
		if contactsAddMessageJWTToken != "" {
			jwtToken = &contactsAddMessageJWTToken
		}
	}
	v := &contacts.AddMessagePayload{
		UserID: body.UserID,
		Msg:    body.Msg,
	}
	v.Oauth = oauth
	v.JWTToken = jwtToken

	return v, nil
}