// Code generated by goa v3.7.3, DO NOT EDIT.
//
// auth HTTP client CLI support package
//
// Command:
// $ goa gen webcup/design

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
	auth "webcup/gen/auth"

	goa "goa.design/goa/v3/pkg"
)

// BuildEmailExistPayload builds the payload for the auth email-exist endpoint
// from CLI flags.
func BuildEmailExistPayload(authEmailExistBody string, authEmailExistOauth string) (*auth.EmailExistPayload, error) {
	var err error
	var body EmailExistRequestBody
	{
		err = json.Unmarshal([]byte(authEmailExistBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"guillaume@gmail.com\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if authEmailExistOauth != "" {
			oauth = &authEmailExistOauth
		}
	}
	v := &auth.EmailExistPayload{
		Email: body.Email,
	}
	v.Oauth = oauth

	return v, nil
}

// BuildSendConfirmationPayload builds the payload for the auth
// send-confirmation endpoint from CLI flags.
func BuildSendConfirmationPayload(authSendConfirmationBody string, authSendConfirmationOauth string) (*auth.SendConfirmationPayload, error) {
	var err error
	var body SendConfirmationRequestBody
	{
		err = json.Unmarshal([]byte(authSendConfirmationBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"guillaume@gmail.com\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if authSendConfirmationOauth != "" {
			oauth = &authSendConfirmationOauth
		}
	}
	v := &auth.SendConfirmationPayload{
		Email: body.Email,
	}
	v.Oauth = oauth

	return v, nil
}

// BuildResetPasswordPayload builds the payload for the auth reset-password
// endpoint from CLI flags.
func BuildResetPasswordPayload(authResetPasswordBody string, authResetPasswordOauth string) (*auth.ResetPasswordPayload, error) {
	var err error
	var body ResetPasswordRequestBody
	{
		err = json.Unmarshal([]byte(authResetPasswordBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"code\": \"ZGI5E\",\n      \"confirm_password\": \"JeSuisUnTest974\",\n      \"email\": \"guillaume@gmail.com\",\n      \"password\": \"JeSuisUnTest974\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

		if utf8.RuneCountInString(body.Code) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.code", body.Code, utf8.RuneCountInString(body.Code), 5, true))
		}
		if utf8.RuneCountInString(body.Code) > 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.code", body.Code, utf8.RuneCountInString(body.Code), 5, false))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.password", body.Password, "\\d"))
		if utf8.RuneCountInString(body.Password) < 9 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", body.Password, utf8.RuneCountInString(body.Password), 9, true))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.confirm_password", body.ConfirmPassword, "\\d"))
		if utf8.RuneCountInString(body.ConfirmPassword) < 9 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.confirm_password", body.ConfirmPassword, utf8.RuneCountInString(body.ConfirmPassword), 9, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var oauth *string
	{
		if authResetPasswordOauth != "" {
			oauth = &authResetPasswordOauth
		}
	}
	v := &auth.ResetPasswordPayload{
		Email:           body.Email,
		Code:            body.Code,
		Password:        body.Password,
		ConfirmPassword: body.ConfirmPassword,
	}
	v.Oauth = oauth

	return v, nil
}