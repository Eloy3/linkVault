package models

import "time"

type Bookmark struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
