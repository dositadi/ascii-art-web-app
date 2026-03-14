package models

type Request struct {
	Text   string
	Banner string
}

type RequestResponse struct {
	Original    string
	Transformed string
	BannerFont  string
	CreatedAt   string
}
