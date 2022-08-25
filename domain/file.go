package domain

import "gopkg.in/guregu/null.v4"

type File struct {
	ID          int         `json:"id"`
	FolderShaID null.String `json:"folder_sha_id"`
	ShaID       string      `json:"sha_id"`
	UserID      int         `json:"user_id"`
	Path        string      `json:"path"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	CreatedAt   null.Time   `json:"created_at"`
	UpdatedAt   null.Time   `json:"updated_at"`
}
