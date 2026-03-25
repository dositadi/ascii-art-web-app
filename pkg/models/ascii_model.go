package models

import "time"

type Ascii struct {
	Id              string    `json:"id"`
	UserId          string    `json:"user_id"`
	InputText        string    `json:"input_text"`
	Font            string    `json:"banner"`
	AsciiText       string    `json:"ascii_text"`
	DownloadedAsImg bool      `json:"downloaded_as_image"`
	DownloadedAsTxt bool      `json:"downloaded_as_txt"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time `json:"created_at"`
}
