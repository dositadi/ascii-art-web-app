package services

import (
	"html/template"
	"net/http"
	"time"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) RenderHistoryPage(w http.ResponseWriter, r *http.Request) *m.Error {
	funcMap := template.FuncMap{
		"Format": func(t time.Time) string {
			return t.Format("02 Jan 2006 15:04")
		},
	}
	temp, err := template.New("history.html").Funcs(funcMap).ParseFiles("web/static/internal_pages/history.html", "web/templates/history_partial.html", "web/templates/history_tabs_partial.html")
	if err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_CODE,
			Details: err.Error(),
			Code:    h.SERVER_ERR,
		}
	}

	val := r.Context().Value("user_id")
	var user_id string

	if id, ok := val.(string); ok {
		user_id = id
	}

	_, userName, _, err2 := s.Repository.GetHashedPasswordIDAndName(r.Context(), &user_id, nil)
	if err2 != nil {
		return err2
	}

	asciiArts, err5 := s.Repository.GetAllUsersSavedAscii(r.Context(), user_id)
	if err5 != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_CODE,
			Details: err.Error(),
			Code:    h.SERVER_ERR,
		}
	}

	namePrefix := s.GetNamePrefix(userName)

	historyPageDetail := struct {
		AsciiArts         []m.Ascii
		UserName          string
		NamePrefix        string
		AsciiRoute        string
		AboutRoute        string
		HelpRoute         string
		ContributorsRoute string
	}{
		AsciiArts:         asciiArts,
		UserName:          userName,
		NamePrefix:        namePrefix,
		AsciiRoute:        h.ASCII_ROUTE,
		AboutRoute:        h.ABOUT_US_ROUTE,
		HelpRoute:         h.HELP_ROUTE,
		ContributorsRoute: h.CONTRIBUTORS_ROUTE,
	}

	if s.GetHxRequestStatus(r) {
		if err3 := temp.ExecuteTemplate(w, "history", historyPageDetail); err3 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err3.Error(),
				Code:    h.SERVER_ERR,
			}
		}
	} else {
		if err4 := temp.Execute(w, historyPageDetail); err4 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err4.Error(),
				Code:    h.SERVER_ERR,
			}
		}
	}

	return nil
}
