// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: activity_list.sql

package db

import (
	"context"
	"time"
)

const createActivity = `-- name: CreateActivity :one
INSERT INTO activity_lists(
    trip_owner,
    activity,
    date_time
) VALUES (
    $1, $2, $3
) RETURNING id, trip_owner, activity, date_time, checked
`

type CreateActivityParams struct {
	TripOwner int64     `json:"trip_owner"`
	Activity  string    `json:"activity"`
	DateTime  time.Time `json:"date_time"`
}

func (q *Queries) CreateActivity(ctx context.Context, arg CreateActivityParams) (ActivityList, error) {
	row := q.queryRow(ctx, q.createActivityStmt, createActivity, arg.TripOwner, arg.Activity, arg.DateTime)
	var i ActivityList
	err := row.Scan(
		&i.ID,
		&i.TripOwner,
		&i.Activity,
		&i.DateTime,
		&i.Checked,
	)
	return i, err
}

const deleteActivity = `-- name: DeleteActivity :exec
DELETE FROM activity_lists WHERE trip_owner = &1
`

func (q *Queries) DeleteActivity(ctx context.Context) error {
	_, err := q.exec(ctx, q.deleteActivityStmt, deleteActivity)
	return err
}

const getActivityUpdate = `-- name: GetActivityUpdate :one
SELECT FROM activity_lists
WHERE trip_owner = $1
LIMIT 1
FOR NO KEY UPDATE
`

type GetActivityUpdateRow struct {
}

func (q *Queries) GetActivityUpdate(ctx context.Context, tripOwner int64) (GetActivityUpdateRow, error) {
	row := q.queryRow(ctx, q.getActivityUpdateStmt, getActivityUpdate, tripOwner)
	var i GetActivityUpdateRow
	err := row.Scan()
	return i, err
}

const listActivities = `-- name: ListActivities :many
SELECT FROM activity_lists
WHERE trip_owner = $1
LIMIT $2
OFFSET $3
`

type ListActivitiesParams struct {
	TripOwner int64 `json:"trip_owner"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

type ListActivitiesRow struct {
}

func (q *Queries) ListActivities(ctx context.Context, arg ListActivitiesParams) ([]ListActivitiesRow, error) {
	rows, err := q.query(ctx, q.listActivitiesStmt, listActivities, arg.TripOwner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListActivitiesRow{}
	for rows.Next() {
		var i ListActivitiesRow
		if err := rows.Scan(); err != nil {
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