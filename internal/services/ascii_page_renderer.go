package services

import (
	"net/http"
	"text/template"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) RenderAsciiArtPage(w http.ResponseWriter, r *http.Request) *m.Error {
	temp, err := template.New("ascii.html").ParseFiles("web/static/internal_pages/ascii.html", "web/templates/ascii_partial.html")
	if err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}

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

	if s.GetHxRequestStatus(r) {
		if err1 := temp.ExecuteTemplate(w, "ascii", asciiPageDetails); err1 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err1.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
	} else {
		if err2 := temp.Execute(w, asciiPageDetails); err2 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err2.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
	}
	return nil
}
