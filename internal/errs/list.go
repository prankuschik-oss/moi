package errs

import "errors"

var (
	ErrNotfound           = errors.New("not found")
	ErrEmployeesNotfound      = errors.New("employees not found")
	ErrInvalidEmployeesID     = errors.New("invalid employees id")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidFieldValue  = errors.New("invalid field value")
	ErrInvalidEmployeesName  = errors.New("invalid eployees name, min 4 symbols")
)
