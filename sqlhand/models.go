package sqlhand

import "time"

type Book struct {
	ID       int
	Name     string
	Author   string
	Recensia string
	Readed   bool

	CreatedAt time.Time
	ReadedAt  *time.Time
}

