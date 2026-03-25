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

	val := r.Context().Value("user_id")
	var user_id string

	if id, ok := val.(string); ok {
		user_id = id
	}

	_, userName, _, err1 := s.Repository.GetHashedPasswordIDAndName(r.Context(), &user_id, nil)
	if err1 != nil {
		return err1
	}

	namePrefix := s.GetNamePrefix(userName)

	asciiPageDetails := struct {
		UserName          string
		NamePrefix        string
		LogOutRoute       string
		TextKey           string
		BannerKey         string
		PostRoute         string
		Output            string
		DownloadImgRoute  string
		DownloadTxtRoute  string
		SaveOutputRoute   string
		ViewHistoryRoute  string
		RecievedOutput    bool
		AboutRoute        string
		HelpRoute         string
		ContributorsRoute string
		HomePageRoute     string
	}{
		UserName:          userName,
		NamePrefix:        namePrefix,
		TextKey:           h.TEXT_KEY,
		BannerKey:         h.BANNER_KEY,
		PostRoute:         h.ASCII_ROUTE,
		RecievedOutput:    false,
		ViewHistoryRoute:  h.HISTORY_ROUTE,
		AboutRoute:        h.ABOUT_US_ROUTE,
		HelpRoute:         h.HELP_ROUTE,
		ContributorsRoute: h.CONTRIBUTORS_ROUTE,
		HomePageRoute:     h.HOME_ROUTE,
	}

	if s.GetHxRequestStatus(r) {
		if err2 := temp.ExecuteTemplate(w, "ascii", asciiPageDetails); err2 != nil {
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
