package common

import (
	"golang-utility/pkg/errors"
	"net/http"
)

type TransactionError struct {
	errors.BaseError
	Status int
}

func (err *TransactionError) StatusCode() int {
	if err.Status > 0 {
		return err.Status
	}
	return http.StatusInternalServerError
}
