package services

import (
	"fmt"
	"net/http"
	"time"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) TransformText(w http.ResponseWriter, r *http.Request, text, banner string, start time.Time) *m.Error {
	latinWords, err := s.AsciiTransformer.SplitInputByNewline(text)
	if err != nil {
		return &m.Error{
			Error:   h.PROCESS_TEXT_ERR,
			Details: h.PROCESS_TEXT_ERR_DETAIL,
			Code:    h.SERVER_ERR_CODE,
		}
	}

	asciiWords, cliEquivalent, err2 := s.AsciiTransformer.ReadWords(latinWords, banner)
	if err2 != nil {
		return err
	}

	formattedAsciiWords := s.AsciiTransformer.FormatAsciiWords(asciiWords)

	uiCliInput := fmt.Sprintf(`forge --text "%s" --font %s --size small`, cliEquivalent, banner)
	AsciiForgeHeader := fmt.Sprintf(`ASCIIForge v1.0.0  ·  font: %s  ·  size: small  ·  border: none`, banner)
	responseTime := fmt.Sprintf("✓  output rendered in %s", time.Since(start).Abs().String())
	toolbarFont := fmt.Sprintf("font: %s", banner)
	toolbarChars := fmt.Sprintf("chars: %v", len(text))
	toolbarLines := fmt.Sprintf("lines: %v", len(latinWords))

	var maxWord int

	for _, word := range latinWords {
		if len(word) > maxWord {
			maxWord = len(word)
		}
	}

	AsciiForgeFooter := fmt.Sprintf("width: %v chars  ·  height: %v lines  ·  encoding: UTF-8", maxWord, len(latinWords))

	err3 := s.AsciiTransformer.RenderAsciiArtOutput(w, r, text, banner, cliEquivalent, formattedAsciiWords, uiCliInput, AsciiForgeHeader, responseTime, toolbarFont, toolbarChars, toolbarLines, AsciiForgeFooter)
	if err3 != nil {
		return err3
	}
	return nil
}
