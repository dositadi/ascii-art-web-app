package services

import (
	"context"
	"net/http"

	at "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/tranformer"
	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

type Repo interface {
	/* InsertAscii(ascii m.Ascii) *m.Error
	GetAscii() ([]m.Ascii, *m.Error) */

	InsertUser(ctx context.Context, user m.User) *m.Error
	GetHashedPasswordIDAndName(ctx context.Context, user_id, email *string) (string, string, string, *m.Error)
	PingDB() *m.Error
}

type Transformer interface {
	SplitInputByNewline(input string) ([]string, *m.Error)
	ReadAsciiFromFont(rn rune, banner string) ([]string, *m.Error)
	ReadWords(input []string, banner string) ([][][]string, string, *m.Error)
	FormatAsciiWords(asciiWords [][][]string) string
	RenderAsciiArtOutput(w http.ResponseWriter, r *http.Request, formattedAsciiWords, uiCliInput, asciiForgeHeader, responseTime, toolbarFont, toolbarChars, toolbarLines, asciiForgeFooter string) *m.Error
}

type Service struct {
	Repository       Repo
	AsciiTransformer Transformer
}

func ConstructNewService(repo Repo) *Service {
	return &Service{
		Repository:       repo,
		AsciiTransformer: at.CreateNewAsciiTransformer(),
	}
}

func (s *Service) GetHxRequestStatus(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
