package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"io"
	"log/slog"
	"net/http"
	db2 "request-bin/pkg/db"
	"request-bin/pkg/sqlc"
	"strings"
	"time"
)

func handleRequest(ctx context.Context, currUUid uuid.UUID, respCode int, req *http.Request) error {

	ts := time.Now()

	db := db2.MustGetDatabase()

	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func(tx pgx.Tx, ctx context.Context) {
		errC := tx.Rollback(ctx)
		if errC != nil {
			slog.Error("Error rolling back transaction", "error", errC)
		}
	}(tx, ctx)

	queries := sqlc.New(db).WithTx(tx)

	var bodyBytes []byte

	if req.ContentLength != 0 {
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return err
		}
	}

	err = queries.CreateRequest(ctx, sqlc.CreateRequestParams{
		ID: pgtype.UUID{
			Bytes: currUUid,
			Valid: true,
		},
		Method: pgtype.Text{
			String: req.Method,
			Valid:  true,
		},
		Content: bodyBytes,
		SourceIp: pgtype.Text{
			String: req.RemoteAddr,
			Valid:  true,
		},
		ResponseCode: pgtype.Int4{
			Int32: int32(respCode),
			Valid: true,
		},
		Timestamp: pgtype.Timestamptz{
			Time:             ts,
			InfinityModifier: 0,
			Valid:            true,
		},
	})
	if err != nil {
		return err
	}

	headersArray := make([]sqlc.CreateHeadersParams, 0)
	for name, value := range req.Header {
		curr := sqlc.CreateHeadersParams{
			RequestID: pgtype.UUID{
				Bytes: currUUid,
				Valid: true,
			},
			Name: name,
			Value: pgtype.Text{
				String: strings.Join(value, ","),
				Valid:  true,
			},
		}
		headersArray = append(headersArray, curr)
	}
	_, err = queries.CreateHeaders(ctx, headersArray)
	if err != nil {
		return err
	}

	queryArray := make([]sqlc.CreateQueryParametersParams, 0)
	for name, value := range req.URL.Query() {
		curr := sqlc.CreateQueryParametersParams{
			RequestID: pgtype.UUID{
				Bytes: currUUid,
				Valid: true,
			},
			Name: name,
			Value: pgtype.Text{
				String: strings.Join(value, ","),
				Valid:  true,
			},
		}
		queryArray = append(queryArray, curr)
	}

	_, err = queries.CreateQueryParameters(ctx, queryArray)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func DefaultRoute(c *gin.Context) {

	currUUid, err := uuid.NewV7()
	if err != nil {
		slog.Error("Error generating uuid", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = handleRequest(c.Request.Context(), currUUid, http.StatusOK, c.Request)
	if err != nil {
		slog.Error("Error handling request", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.Status(http.StatusOK)
}
