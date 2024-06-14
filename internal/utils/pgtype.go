package utils

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertUUIDToPgType(uuid uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: uuid,
		Valid: true,
	}
}
