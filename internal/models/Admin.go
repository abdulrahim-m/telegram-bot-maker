package models

import "time"

type AdminRole string

const (
	RoleCreator AdminRole = "creator"
	RoleManager AdminRole = "manager"
	RoleDefault AdminRole = "default"
)

type Admin struct {
	ID        int64     `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Role      AdminRole `db:"role"`
	CreatedAt time.Time `db:"created_at"`
}
