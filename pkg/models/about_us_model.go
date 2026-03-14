package models

import "time"

type AboutUs struct {
	Name      *string   `json:"name"`
	Content   *string   `json:"content"`
	UpdatedAt string    `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
