package services

import (
	"net/http"
	"text/template"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) RenderWelcomePage(w http.ResponseWriter, r *http.Request) *m.Error {
	temp, err := template.New("welcome_page.html").ParseFiles("web/static/welcome_page/welcome_page.html", "web/templates/welcome_partial.html")
	if err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}

	welcomePageData := struct {
		LoginRoute     string
		SignupRoute    string
		WelcomeMessage string
		WelcomeRoute   string
	}{
		WelcomeRoute: h.WELCOME_ROUTE,
		LoginRoute:   h.LOGIN_ROUTE,
		SignupRoute:  h.SIGNUP_ROUTE,
		WelcomeMessage: `Transform your words into stunning ASCII banners.
        Choose your style, craft your message, and share
        something truly memorable.`,
	}

	if s.GetHxRequestStatus(r) {
		if err2 := temp.ExecuteTemplate(w, "welcome", welcomePageData); err2 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err2.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
	} else {
		if err3 := temp.Execute(w, welcomePageData); err3 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err3.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
	}
	return nil
}
