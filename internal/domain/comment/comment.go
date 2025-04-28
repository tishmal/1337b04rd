package entity

import "time"

type Comment struct {
	ID        string
	Content   string
	PostID    string
	ReplyToID string // ID поста или комментария, на который отвечают
	UserID    string
	UserName  string
	AvatarURL string
	CreatedAt time.Time
}
