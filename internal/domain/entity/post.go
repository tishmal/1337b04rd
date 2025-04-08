package entity

import "time"

type Post struct {
	ID        string
	Title     string
	Content   string
	ImageURL  string
	UserID    string
	UserName  string
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}
