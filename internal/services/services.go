package services

import (
	"context"
	"net/http"
	"strings"

	at "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/tranformer"
	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

type Repo interface {
	/* InsertAscii(ascii m.Ascii) *m.Error
	GetAscii() ([]m.Ascii, *m.Error) */

	// User CRUD operations
	InsertUser(ctx context.Context, user m.User) *m.Error
	GetHashedPasswordIDAndName(ctx context.Context, user_id, email *string) (string, string, string, *m.Error)
	PingDB() *m.Error

	// Ascii table CRUD operations
	InsertAscii(ctx context.Context, ascii m.Ascii) *m.Error
	GetAllUsersSavedAscii(ctx context.Context, user_id string, limit, offset int, font string) ([]m.Ascii, *m.Error)
	DeleteFromAscii(ctx context.Context, user_id string) *m.Error
	ClearAll(ctx context.Context, user_id string) *m.Error
	GetTableLenght(ctx context.Context, user_id, font string) (int, *m.Error)
	UpdateAsciiOutputsTable(ctx context.Context, id string) *m.Error
}

type Transformer interface {
	SplitInputByNewline(input string) ([]string, *m.Error)
	ReadAsciiFromFont(rn rune, banner string) ([]string, *m.Error)
	ReadWords(input []string, banner string) ([][][]string, string, *m.Error)
	FormatAsciiWords(asciiWords [][][]string) string
	RenderAsciiArtOutput(w http.ResponseWriter, r *http.Request, text, banner, cliEquivalent, formattedAsciiWords, uiCliInput, asciiForgeHeader, responseTime, toolbarFont, toolbarChars, toolbarLines, asciiForgeFooter string) *m.Error
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

func (s *Service) GetNamePrefix(userName string) string {
	names := strings.Fields(userName)
	var namesPrefix strings.Builder

	for i, name := range names {
		if i > 2 {
			namesPrefix.WriteString("X")
			break
		}
		if name != "" {
			namesPrefix.WriteString(string(name[0]))
		}
	}
	return namesPrefix.String()
}

func (s *Service) FormatAscii(text, banner string) (string, *m.Error) {
	latinWords, err := s.AsciiTransformer.SplitInputByNewline(text)
	if err != nil {
		return "", err
	}

	asciiWords, _, err2 := s.AsciiTransformer.ReadWords(latinWords, banner)
	if err2 != nil {
		return "", err2
	}

	formattedAsciiWords := s.AsciiTransformer.FormatAsciiWords(asciiWords)

	return formattedAsciiWords, nil
}
