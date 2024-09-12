package common

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func isPgErrorCode(err error, code string) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == code {
			return true
		}
	}
	return false
}

func IsDuplicateKeyError(err error) bool {
	return isPgErrorCode(err, "23505")
}
