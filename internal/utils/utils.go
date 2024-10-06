package utils

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

const (
	UniqueViolationErr = pq.ErrorCode("23505")
)

func IsDuplicateError(err error) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == UniqueViolationErr
	}
	return false
}

func IsNotFoundError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
