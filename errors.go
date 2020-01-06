package mic

import "github.com/micro/go-micro/errors"

// BadRequest generates a 400 error.
func ErrorBadRequest(id, format string, a ...interface{}) error {
	return errors.BadRequest(id, format, a...)
}

// NotFound generates a 404 error.
func ErrorNotFound(id, format string, a ...interface{}) error {
	return errors.NotFound(id, format, a...)
}

// Unauthorized generates a 401 error.
func ErrorUnauthorized(id, format string, a ...interface{}) error {
	return errors.Unauthorized(id, format, a...)
}

// Forbidden generates a 403 error.
func ErrorForbidden(id, format string, a ...interface{}) error {
	return errors.Forbidden(id, format, a...)
}

// InternalServer generates a 500 error.
func ErrorInternalServer(id, format string, a ...interface{}) error {
	return errors.InternalServerError(id, format, a...)
}
