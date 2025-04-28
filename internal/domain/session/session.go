package entity

import "time"

type Session struct {
	ID        string
	UserID    string
	UserName  string
	AvatarURL string
	ExpiresAt time.Time
}
