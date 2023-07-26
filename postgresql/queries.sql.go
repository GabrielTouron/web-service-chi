// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: queries.sql

package postgresql

import (
	"context"
)

const listCommands = `-- name: ListCommands :many
SELECT id, name, command, created_at from commands
ORDER BY name
`

func (q *Queries) ListCommands(ctx context.Context) ([]Command, error) {
	rows, err := q.db.QueryContext(ctx, listCommands)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Command
	for rows.Next() {
		var i Command
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Command,
			&i.CreatedAt,
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
