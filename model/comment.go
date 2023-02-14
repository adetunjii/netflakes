package model

import "time"

type Comment struct {
	ID        int64     `json:"id"`
	Body      string    `json:"body"`
	MovieID   int64     `json:"movie_id"`
	MovieUrl  string    `json:"movie_url"`
	SenderIP  string    `json:"sender_ip"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}
