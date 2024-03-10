package main

import "errors"

var (
	ErrUserIDInvalid     = errors.New("user id is invalid")
	ErrUserRoleInvalid   = errors.New("user role is invalid")
	ErrActionTypeInvalid = errors.New("action type is invalid")
	ErrCannotMakeRequest = errors.New("cannot make request with resty client")
	ErrInvalidAddress    = errors.New("invalid host for log minder")
)
