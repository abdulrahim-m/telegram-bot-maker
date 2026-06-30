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
type Session struct {
	SessionID string    `db:"session_id"`
	AdminID   int64     `db:"admin_id"`
	ExpieryAt time.Time `db:"expiery_at"`
}
