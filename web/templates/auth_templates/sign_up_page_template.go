package authtemplates

import (
	"html/template"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func SignUpPageTemplate(w http.ResponseWriter, message *string) *m.Error {
	temp, err := template.New("signup_page.html").ParseFiles("web/static/auth_pages/signup_page.html", "web/static/auth_pages/signup_partials.html")
	if err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}

	signupPageDetail := struct {
		NameKey      string
		PasswordKey  string
		EmailKey     string
		PostUrl      string
		BackURL      string
		LoginPageUrl string
		ErrorMessage string
	}{
		NameKey:      "name",
		PasswordKey:  "password",
		EmailKey:     "email",
		PostUrl:      h.SIGNUP_ROUTE,
		BackURL:      h.WELCOME_ROUTE,
		LoginPageUrl: h.LOGIN_ROUTE,
	}

	if message != nil {
		signupPageDetail.ErrorMessage = *message
	}

	if err2 := temp.Execute(w, signupPageDetail); err2 != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err2.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}
	return nil
}
