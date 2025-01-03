// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: reports.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const deleteReportByID = `-- name: DeleteReportByID :exec
DELETE FROM reports
WHERE report_id = $1
`

func (q *Queries) DeleteReportByID(ctx context.Context, reportID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteReportByID, reportID)
	return err
}

const getReportByID = `-- name: GetReportByID :one
SELECT report_id, created_at, updated_at, post_id, user_id, reason FROM reports 
WHERE report_id = $1
`

func (q *Queries) GetReportByID(ctx context.Context, reportID uuid.UUID) (Report, error) {
	row := q.db.QueryRowContext(ctx, getReportByID, reportID)
	var i Report
	err := row.Scan(
		&i.ReportID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PostID,
		&i.UserID,
		&i.Reason,
	)
	return i, err
}

const listAllReports = `-- name: ListAllReports :many
SELECT report_id, created_at, updated_at, post_id, user_id, reason FROM reports
`

func (q *Queries) ListAllReports(ctx context.Context) ([]Report, error) {
	rows, err := q.db.QueryContext(ctx, listAllReports)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Report
	for rows.Next() {
		var i Report
		if err := rows.Scan(
			&i.ReportID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PostID,
			&i.UserID,
			&i.Reason,
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

const reportPost = `-- name: ReportPost :one
INSERT INTO reports (report_id, created_at, updated_at, post_id, user_id, reason)
VALUES(
   $1,
   NOW(),
   NOW(),
   $2,
   $3,
   $4
)
RETURNING report_id, created_at, updated_at, post_id, user_id, reason
`

type ReportPostParams struct {
	ReportID uuid.UUID
	PostID   uuid.UUID
	UserID   uuid.UUID
	Reason   string
}

func (q *Queries) ReportPost(ctx context.Context, arg ReportPostParams) (Report, error) {
	row := q.db.QueryRowContext(ctx, reportPost,
		arg.ReportID,
		arg.PostID,
		arg.UserID,
		arg.Reason,
	)
	var i Report
	err := row.Scan(
		&i.ReportID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PostID,
		&i.UserID,
		&i.Reason,
	)
	return i, err
}
