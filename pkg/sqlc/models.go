// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Request struct {
	ID           uuid.UUID
	Method       pgtype.Text
	Content      []byte
	SourceIp     pgtype.Text
	ResponseCode pgtype.Int4
	Timestamp    pgtype.Timestamptz
}

type RequestHeader struct {
	RequestID uuid.UUID
	Name      string
	Value     pgtype.Text
}

type RequestQueryParameter struct {
	RequestID uuid.UUID
	Name      string
	Value     pgtype.Text
}
