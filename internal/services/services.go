package services

import (
	"context"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

type Repo interface {
	/* InsertAscii(ascii m.Ascii) *m.Error
	GetAscii() ([]m.Ascii, *m.Error) */

	InsertUser(ctx context.Context, user m.User) *m.Error
	GetHashedPasswordIDAndName(ctx context.Context, email string) (string, string, string, *m.Error)
	PingDB() *m.Error
}

type Service struct {
	Repository Repo
}

func ConstructNewService(repo Repo) *Service {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) GetHxRequestStatus(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
