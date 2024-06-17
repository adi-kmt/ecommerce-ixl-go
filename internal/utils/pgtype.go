package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertUUIDToPgType(uuid uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: uuid,
		Valid: true,
	}
}

func ConvertPgUUIDToString(uuid pgtype.UUID) string {
	return fmt.Sprintf("%x", uuid.Bytes)
}
