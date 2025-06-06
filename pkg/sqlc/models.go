// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Request struct {
	ID           uuid.UUID          `json:"id"`
	Method       pgtype.Text        `json:"method"`
	Content      []byte             `json:"content"`
	SourceIp     pgtype.Text        `json:"sourceIp"`
	ResponseCode pgtype.Int4        `json:"responseCode"`
	Timestamp    pgtype.Timestamptz `json:"timestamp"`
}

type RequestHeader struct {
	RequestID uuid.UUID   `json:"requestId"`
	Name      string      `json:"name"`
	Value     pgtype.Text `json:"value"`
}

type RequestQueryParameter struct {
	RequestID uuid.UUID   `json:"requestId"`
	Name      string      `json:"name"`
	Value     pgtype.Text `json:"value"`
}
