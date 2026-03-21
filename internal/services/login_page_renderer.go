package services

import (
	"net/http"
	"text/template"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) RenderLoginPage(w http.ResponseWriter, r *http.Request, message *string) *m.Error {
	temp, err := template.New("login_page.html").ParseFiles("web/static/auth_pages/login_page.html", "web/templates/login_error_partial.html", "web/templates/login_partial.html")
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
		if err2 := temp.ExecuteTemplate(w, "error", loginPageData); err2 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err2.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
		return nil
	}

	if s.GetHxRequestStatus(r) {
		if err3 := temp.ExecuteTemplate(w, "login", loginPageData); err3 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err3.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
	} else {
		if err4 := temp.Execute(w, loginPageData); err4 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_ERROR,
				Details: err4.Error(),
				Code:    h.PAGE_PARSING_CODE,
			}
		}
	}
	return nil
}
