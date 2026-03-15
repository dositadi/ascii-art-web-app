package welcometemplate

import (
	"html/template"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func WelcomePageTemplate(w http.ResponseWriter) *m.Error {
	temp, err := template.New("welcome_page.html").ParseFiles("web/static/welcome_page/welcome_page.html")
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
	}{
		LoginRoute:  h.LOGIN_ROUTE,
		SignupRoute: h.SIGNUP_ROUTE,
		WelcomeMessage: `Transform your words into stunning ASCII banners.
        Choose your style, craft your message, and share
        something truly memorable.`,
	}

	if err := temp.Execute(w, welcomePageData); err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}
	return nil
}
