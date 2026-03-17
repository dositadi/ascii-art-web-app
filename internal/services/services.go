package services

import (
	"context"

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
