package common

import (
	"golang-utility/pkg/errors"
	"net/http"
)

type NotFoundError struct {
	errors.BaseError
}

func (err *NotFoundError) StatusCode() int {
	return http.StatusNotFound
}
