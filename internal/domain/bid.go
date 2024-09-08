package domain

import "time"

type Bid struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	TenderID    string    `json:"tender_id"`
	AuthorType  string    `json:"author_type"`
	AuthorID    string    `json:"author_id"`
	Version     int       `json:"version"`
	CreatedAt   time.Time `json:"created_at"`
}
