package services

import (
	"fmt"
	"net/http"
	"slices"
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
	var maxWordLength int

	if len(latinWords) != 0 || latinWords != nil {
		maxWordLength = len(slices.Max(latinWords))
	}

	AsciiForgeFooter := fmt.Sprintf("width: %v chars  ·  height: %v lines  ·  encoding: UTF-8", maxWordLength, len(latinWords))

	err3 := s.AsciiTransformer.RenderAsciiArtOutput(w, r, formattedAsciiWords, uiCliInput, AsciiForgeHeader, responseTime, toolbarFont, toolbarChars, toolbarLines, AsciiForgeFooter)
	if err3 != nil {
		return err3
	}
	return nil
}
