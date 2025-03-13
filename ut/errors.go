package ut

import "go-micro.dev/v5/errors"

func ParseMicError(err error) (*errors.Error, bool) {
	if err == nil {
		return nil, false
	}

	e, ok := err.(*errors.Error)

	return e, ok
}
