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
	RequestID uuid.UUID   `json:"request_id"`
	Name      string      `json:"name"`
	Value     pgtype.Text `json:"value"`
}

type CreateQueryParametersParams struct {
	RequestID uuid.UUID   `json:"request_id"`
	Name      string      `json:"name"`
	Value     pgtype.Text `json:"value"`
}

const createRequest = `-- name: CreateRequest :exec
INSERT INTO requests (id, method, content, source_ip, response_code, timestamp)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateRequestParams struct {
	ID           uuid.UUID          `json:"id"`
	Method       pgtype.Text        `json:"method"`
	Content      []byte             `json:"content"`
	SourceIp     pgtype.Text        `json:"source_ip"`
	ResponseCode pgtype.Int4        `json:"response_code"`
	Timestamp    pgtype.Timestamptz `json:"timestamp"`
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

const getRequests = `-- name: GetRequests :many
select id, method,content,source_ip,response_code,timestamp
from requests
`

func (q *Queries) GetRequests(ctx context.Context) ([]Request, error) {
	rows, err := q.db.Query(ctx, getRequests)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Request
	for rows.Next() {
		var i Request
		if err := rows.Scan(
			&i.ID,
			&i.Method,
			&i.Content,
			&i.SourceIp,
			&i.ResponseCode,
			&i.Timestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
