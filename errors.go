package mic

import "go-micro.dev/v5/errors"

// ErrorBadRequest generates a 400 error.
func ErrorBadRequest(id, format string, a ...interface{}) error {
	return errors.BadRequest(id, format, a...)
}

// ErrorNotFound generates a 404 error.
func ErrorNotFound(id, format string, a ...interface{}) error {
	return errors.NotFound(id, format, a...)
}

// ErrorUnauthorized generates a 401 error.
func ErrorUnauthorized(id, format string, a ...interface{}) error {
	return errors.Unauthorized(id, format, a...)
}

// ErrorForbidden generates a 403 error.
func ErrorForbidden(id, format string, a ...interface{}) error {
	return errors.Forbidden(id, format, a...)
}

// ErrorInternalServer generates a 500 error.
func ErrorInternalServer(id, format string, a ...interface{}) error {
	return errors.InternalServerError(id, format, a...)
}
