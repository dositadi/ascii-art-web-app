package tranformer

import (
	"html/template"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (at *AsciiTransform) RenderAsciiArtOutput(w http.ResponseWriter, r *http.Request, text, banner, cliEquivalent, formattedAsciiWords, uiCliInput, asciiForgeHeader, responseTime, toolbarFont, toolbarChars, toolbarLines, asciiForgeFooter string) *m.Error {
	temp, err := template.New("ascii.html").ParseFiles("web/static/internal_pages/ascii.html", "web/templates/ascii_partial.html", "web/templates/ascii_response_partial.html")
	if err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}

	// Response details
	responseDetails := struct {
		FormattedAsciiWords string
		UiCliInput          string
		AsciiForgeHeader    string
		ResponseTime        string
		ToolbarFont         string
		ToolbarChars        string
		ToolbarLines        string
		AsciiForgeFooter    string
	}{
		FormattedAsciiWords: formattedAsciiWords,
		UiCliInput:          uiCliInput,
		AsciiForgeHeader:    asciiForgeHeader,
		ResponseTime:        responseTime,
		ToolbarFont:         toolbarFont,
		ToolbarChars:        toolbarChars,
		ToolbarLines:        toolbarLines,
		AsciiForgeFooter:    asciiForgeFooter,
	}

	if err2 := temp.ExecuteTemplate(w, "ascii-response", responseDetails); err2 != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err2.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}

	return nil
}
