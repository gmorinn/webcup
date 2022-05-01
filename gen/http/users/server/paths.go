// Code generated by goa v3.7.3, DO NOT EDIT.
//
// HTTP request path constructors for the users service.
//
// Command:
// $ goa gen webcup/design

package server

import (
	"fmt"
)

// DeleteUserUsersPath returns the URL path to the users service deleteUser HTTP endpoint.
func DeleteUserUsersPath(id string) string {
	return fmt.Sprintf("/v1/web/user/remove/%v", id)
}

// GetUserByIDUsersPath returns the URL path to the users service getUserByID HTTP endpoint.
func GetUserByIDUsersPath(id string) string {
	return fmt.Sprintf("/v1/web/user/%v", id)
}

// UpdateDescriptionUsersPath returns the URL path to the users service updateDescription HTTP endpoint.
func UpdateDescriptionUsersPath() string {
	return "/v1/web/user/edit"
}