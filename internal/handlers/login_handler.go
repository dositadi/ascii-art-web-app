package handlers

import (
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	at "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/web/templates/auth_templates"
)

func (s *Handler) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	err := at.LoginPageTemplate(w, nil)
	if err != nil {
		err := h.ErrorToJson(m.Error{Error: err.Error, Details: err.Details, Code: err.Code})
		h.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}

func (s *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	ctx := r.Context()

	activeUser, err := s.Service.LoginUser(ctx, email, password)
	if err != nil {
		if err.Error == h.UNAUTHORIZED_ERR {
			at.LoginPageTemplate(w, &err.Details)
			return
		} else {
			// Display in the page for errors
		}
	}

	refreshToken, err2 := h.GenerateRefreshJWT(activeUser)
	if err2 != nil {
		at.LoginPageTemplate(w, &err2.Details)
		return
	}

	accessToken, err3 := h.GenerateAccessJWT(activeUser)
	if err3 != nil {
		at.LoginPageTemplate(w, &err3.Details)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "refresh-jwt", Value: refreshToken, HttpOnly: true, Secure: true, SameSite: http.SameSiteLaxMode})
	http.SetCookie(w, &http.Cookie{Name: "access-token", Value: accessToken, HttpOnly: true, SameSite: http.SameSiteLaxMode, Path: "/"})
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
