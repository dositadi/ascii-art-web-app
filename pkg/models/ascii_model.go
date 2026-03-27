package models

import "time"

type Ascii struct {
	Id              string
	UserId          string
	InputText       string
	Font            string
	AsciiText       string
	DownloadedAsImg bool
	DownloadedAsTxt bool
	UpdatedAt       time.Time
	CreatedAt       time.Time
}
