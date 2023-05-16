package postgres

import (
	"golang-utility/pkg/errors"
)

type PostgresConnectionError struct {
	errors.BaseError
}

type QueryExecError struct {
	errors.BaseError
}
