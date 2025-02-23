// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateHeadersParams struct {
	RequestID uuid.UUID
	Name      string
	Value     pgtype.Text
}

type CreateQueryParametersParams struct {
	RequestID uuid.UUID
	Name      string
	Value     pgtype.Text
}

const createRequest = `-- name: CreateRequest :exec
INSERT INTO requests (id, method, content, source_ip, response_code, timestamp)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateRequestParams struct {
	ID           uuid.UUID
	Method       pgtype.Text
	Content      []byte
	SourceIp     pgtype.Text
	ResponseCode pgtype.Int4
	Timestamp    pgtype.Timestamptz
}

func (q *Queries) CreateRequest(ctx context.Context, arg CreateRequestParams) error {
	_, err := q.db.Exec(ctx, createRequest,
		arg.ID,
		arg.Method,
		arg.Content,
		arg.SourceIp,
		arg.ResponseCode,
		arg.Timestamp,
	)
	return err
}
