// Code generated by sqlc. DO NOT EDIT.
// source: data.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const countData = `-- name: CountData :one
SELECT COUNT(*) FROM data
WHERE deleted_at IS NULL
`

func (q *Queries) CountData(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countData)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countDataByUserID = `-- name: CountDataByUserID :one
SELECT COUNT(*) FROM data
WHERE user_id = $1 AND
deleted_at IS NULL
`

func (q *Queries) CountDataByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, countDataByUserID, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createData = `-- name: CreateData :one
INSERT INTO data (title, description, img, category, user_id) 
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, deleted_at, title, description, user_id, img, category
`

type CreateDataParams struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Img         sql.NullString `json:"img"`
	Category    Futur          `json:"category"`
	UserID      uuid.UUID      `json:"user_id"`
}

func (q *Queries) CreateData(ctx context.Context, arg CreateDataParams) (Datum, error) {
	row := q.db.QueryRowContext(ctx, createData,
		arg.Title,
		arg.Description,
		arg.Img,
		arg.Category,
		arg.UserID,
	)
	var i Datum
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Description,
		&i.UserID,
		&i.Img,
		&i.Category,
	)
	return i, err
}

const getDataByID = `-- name: GetDataByID :one
SELECT id, created_at, updated_at, deleted_at, title, description, user_id, img, category FROM data
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetDataByID(ctx context.Context, id uuid.UUID) (Datum, error) {
	row := q.db.QueryRowContext(ctx, getDataByID, id)
	var i Datum
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Title,
		&i.Description,
		&i.UserID,
		&i.Img,
		&i.Category,
	)
	return i, err
}

const getDataByUserID = `-- name: GetDataByUserID :many
SELECT id, created_at, updated_at, deleted_at, title, description, user_id, img, category FROM data
WHERE user_id = $1
AND deleted_at IS NULL
LIMIT $2 OFFSET $3
`

type GetDataByUserIDParams struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) GetDataByUserID(ctx context.Context, arg GetDataByUserIDParams) ([]Datum, error) {
	rows, err := q.db.QueryContext(ctx, getDataByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Datum{}
	for rows.Next() {
		var i Datum
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Title,
			&i.Description,
			&i.UserID,
			&i.Img,
			&i.Category,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listData = `-- name: ListData :many
SELECT id, created_at, updated_at, deleted_at, title, description, user_id, img, category FROM data
WHERE deleted_at IS NULL
AND (title ILIKE $1 OR description ILIKE $1)
LIMIT 5
`

func (q *Queries) ListData(ctx context.Context, title string) ([]Datum, error) {
	rows, err := q.db.QueryContext(ctx, listData, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Datum{}
	for rows.Next() {
		var i Datum
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Title,
			&i.Description,
			&i.UserID,
			&i.Img,
			&i.Category,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDataMostRecent = `-- name: ListDataMostRecent :many
SELECT id, created_at, updated_at, deleted_at, title, description, user_id, img, category FROM data
WHERE deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type ListDataMostRecentParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListDataMostRecent(ctx context.Context, arg ListDataMostRecentParams) ([]Datum, error) {
	rows, err := q.db.QueryContext(ctx, listDataMostRecent, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Datum{}
	for rows.Next() {
		var i Datum
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Title,
			&i.Description,
			&i.UserID,
			&i.Img,
			&i.Category,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateData = `-- name: UpdateData :exec
UPDATE data
SET title = $1,
    description = $2,
    img = $3,
    category = $4,
    updated_at = NOW()
WHERE id = $5
`

type UpdateDataParams struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Img         sql.NullString `json:"img"`
	Category    Futur          `json:"category"`
	ID          uuid.UUID      `json:"id"`
}

func (q *Queries) UpdateData(ctx context.Context, arg UpdateDataParams) error {
	_, err := q.db.ExecContext(ctx, updateData,
		arg.Title,
		arg.Description,
		arg.Img,
		arg.Category,
		arg.ID,
	)
	return err
}
