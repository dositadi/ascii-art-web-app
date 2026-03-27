package services

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	"github.com/google/uuid"
)

func (s *Service) SaveAscii(ctx context.Context, text, banner, user_id string) *m.Error {
	latinWords, err := s.AsciiTransformer.SplitInputByNewline(text)
	if err != nil {
		return err
	}

	asciiWords, _, err2 := s.AsciiTransformer.ReadWords(latinWords, banner)
	if err2 != nil {
		return err2
	}

	formattedAsciiWords := s.AsciiTransformer.FormatAsciiWords(asciiWords)

	ascii := m.Ascii{
		Id:        uuid.NewString(),
		UserId:    user_id,
		InputText: text,
		Font:      banner,
		AsciiText: formattedAsciiWords,
	}

	err3 := s.Repository.InsertAscii(ctx, ascii, user_id)
	if err3 != nil {
		return err3
	}
	return nil
}
