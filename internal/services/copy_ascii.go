package services

import (
	"fmt"
	"html/template"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"golang.design/x/clipboard"
)

func (s *Service) CopyAscii(w http.ResponseWriter, r *http.Request) *m.Error {
	if err := r.ParseForm(); err != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	text := r.FormValue(h.TEXT_KEY)
	font := r.FormValue(h.BANNER_KEY)

	input := r.FormValue("input")
	fetchedFont := r.FormValue("font")

	if text == "" && font == "" {
		text = input
		font = fetchedFont
	}

	formattedAsciiWords, err := s.FormatAscii(text, font)
	if err != nil {
		return err
	}

	if err := clipboard.Init(); err != nil {
		fmt.Println(err.Error())
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	clipboard.Write(clipboard.FmtText, []byte(formattedAsciiWords))

	temp, err2 := template.New("ascii_response_partial.html").ParseFiles("web/templates/ascii_response_partial.html")
	if err2 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err2.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	if err3 := temp.ExecuteTemplate(w, "copy", nil); err3 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err3.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	return nil
}
