package entity

import "time"

type Post struct {
	ID           string
	PostID       string
	Title        string
	Content      string
	SessionID    string
	ImageURL     string
	UserAvatarID int
	UserName     string
	AvatarURL    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
