package tranformer

import (
	"html/template"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (at *AsciiTransform) RenderAsciiArtOutput(w http.ResponseWriter, r *http.Request, formattedAsciiWords, uiCliInput, asciiForgeHeader, responseTime, toolbarFont, toolbarChars, toolbarLines, asciiForgeFooter string) *m.Error {
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

	// Main Ascii page details for fallback
	asciiPageDetails := struct {
		LogOutRoute      string
		TextKey          string
		BannerKey        string
		PostRoute        string
		Output           string
		DownloadImgRoute string
		DownloadTxtRoute string
		SaveOutputRoute  string
		ViewHistoryRoute string
		RecievedOutput   bool
	}{
		TextKey:          h.TEXT_KEY,
		BannerKey:        h.BANNER_KEY,
		LogOutRoute:      "",
		PostRoute:        h.ASCII_ROUTE,
		RecievedOutput:   false,
		DownloadImgRoute: "",
		DownloadTxtRoute: "",
		SaveOutputRoute:  "",
		ViewHistoryRoute: "",
		Output:           "",
	}

	if r.Header.Get("HX-Request") == "true" {
		if err2 := temp.ExecuteTemplate(w, "ascii-response", responseDetails); err2 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err2.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
	} else {
		if err3 := temp.Execute(w, asciiPageDetails); err3 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err3.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
	}

	return nil
}
