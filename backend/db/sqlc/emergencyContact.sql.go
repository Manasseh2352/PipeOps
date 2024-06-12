// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: emergencyContact.sql

package db

import (
	"context"
)

const createEmergencyContact = `-- name: CreateEmergencyContact :one
INSERT INTO emergency_contacts(
    owner,
    email,
    phone_number
) VALUES (
    $1, $2, $3
) RETURNING id, owner, email, phone_number
`

type CreateEmergencyContactParams struct {
	Owner       int64  `json:"owner"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (q *Queries) CreateEmergencyContact(ctx context.Context, arg CreateEmergencyContactParams) (EmergencyContact, error) {
	row := q.queryRow(ctx, q.createEmergencyContactStmt, createEmergencyContact, arg.Owner, arg.Email, arg.PhoneNumber)
	var i EmergencyContact
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Email,
		&i.PhoneNumber,
	)
	return i, err
}

const getEmergencyContactUpdate = `-- name: GetEmergencyContactUpdate :one
SELECT id, owner, email, phone_number FROM emergency_contacts
WHERE owner = $1
LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetEmergencyContactUpdate(ctx context.Context, owner int64) (EmergencyContact, error) {
	row := q.queryRow(ctx, q.getEmergencyContactUpdateStmt, getEmergencyContactUpdate, owner)
	var i EmergencyContact
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Email,
		&i.PhoneNumber,
	)
	return i, err
}

const listsAccounts = `-- name: ListsAccounts :many
SELECT id, owner, email, phone_number FROM emergency_contacts
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListsAccountsParams struct {
	Owner  int64 `json:"owner"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListsAccounts(ctx context.Context, arg ListsAccountsParams) ([]EmergencyContact, error) {
	rows, err := q.query(ctx, q.listsAccountsStmt, listsAccounts, arg.Owner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EmergencyContact{}
	for rows.Next() {
		var i EmergencyContact
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Email,
			&i.PhoneNumber,
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
