package services

import (
	"fmt"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) TransformText(w http.ResponseWriter, r *http.Request, text, banner string) *m.Error {
	output, err := s.AsciiTransformer.SplitInputByNewline(text)
	if err != nil {
		return &m.Error{
			Error:   h.PROCESS_TEXT_ERR,
			Details: h.PROCESS_TEXT_ERR_DETAIL,
			Code:    h.SERVER_ERR_CODE,
		}
	}

	asciiWords, err2 := s.AsciiTransformer.ReadWords(output, banner)
	if err2 != nil {
		return err
	}

	formattedAsciiWords := s.AsciiTransformer.FormatAsciiWords(asciiWords)

	fmt.Println(formattedAsciiWords)
	return nil
}
