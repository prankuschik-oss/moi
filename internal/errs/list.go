package errs

import "errors"

var (
	ErrNotfound           = errors.New("not found")
	ErrUsersNotfound      = errors.New("users not found")
	ErrInvalidUsersID     = errors.New("invalid users id")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidFieldValue  = errors.New("invalid field value")
	ErrInvalidUserstName  = errors.New("invalid users name, min 4 symbols")
)
