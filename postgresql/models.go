// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package postgresql

import (
	"time"
)

type Command struct {
	ID        int32
	Name      string
	Command   string
	CreatedAt time.Time
}

type User struct {
	ID        int32
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
