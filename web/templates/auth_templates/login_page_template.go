package authtemplates

import (
	"html/template"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func LoginPageTemplate(w http.ResponseWriter, message *string) *m.Error {
	temp, err := template.New("login_page.html").ParseFiles("web/static/auth_pages/login_page.html", "web/static/auth_pages/login_partials.html")
	if err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}

	loginPageData := struct {
		PasswordKey  string
		EmailKey     string
		PostUrl      string
		SignUpURL    string
		BackURL      string
		ErrorMessage string
	}{
		PostUrl:     h.LOGIN_ROUTE,
		PasswordKey: "password",
		EmailKey:    "email",
		SignUpURL:   h.SIGNUP_ROUTE,
		BackURL:     h.WELCOME_ROUTE,
	}

	if message != nil {
		loginPageData.ErrorMessage = *message
	}

	if err := temp.Execute(w, loginPageData); err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}
	return nil
}
