// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package sqlcpg

import (
	"database/sql"
	"time"
)

type File struct {
	ID          int32
	FolderShaID string
	Name        string
	// folder, note
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ShaID     string
	Path      string
	UserID    int32
}

type Folder struct {
	ID        int32
	ShaID     string
	ParentID  sql.NullInt32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Note struct {
	ID        int32
	FileShaID string
	Note      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID          int32
	Username    string
	Email       sql.NullString
	PhoneNumber sql.NullString
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
}
