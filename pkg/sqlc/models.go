// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Request struct {
	ID           pgtype.UUID
	Method       pgtype.Text
	Content      []byte
	SourceIp     pgtype.Text
	ResponseCode pgtype.Int4
}

type RequestHeader struct {
	RequestID pgtype.UUID
	Name      string
	Value     pgtype.Text
}

type RequestQueryParameter struct {
	RequestID pgtype.UUID
	Name      string
	Value     pgtype.Text
}
