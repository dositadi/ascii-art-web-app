package handlers

import (
	"fmt"
	"net/http"
	"strings"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	at "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/web/templates/auth_templates"
)

func (s *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	firstName := strings.TrimSpace(r.FormValue("first_name"))
	lastName := strings.TrimSpace(r.FormValue("last_name"))
	userEmail := strings.TrimSpace(r.FormValue("email"))
	userPassword := strings.TrimSpace(r.FormValue("password"))

	name := firstName + " " + lastName
	fmt.Println(userPassword)

	ctx := r.Context()

	err2 := s.Service.RegisterUser(ctx, name, userEmail, userPassword)
	if err2 != nil {
		if err2.Error == h.CONFLICT_ERR {
			at.SignUpPageTemplate(w, &err2.Details)
			return
		} else if err2.Error == h.SERVER_ERR {
			at.SignUpPageTemplate(w, &err2.Details)
			return
		} else {
			at.SignUpPageTemplate(w, &err2.Details)
			return
		}
	}

	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

func (s *Handler) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	err := at.SignUpPageTemplate(w, nil)
	if err != nil {
		err := h.ErrorToJson(m.Error{Error: err.Error, Details: err.Details, Code: err.Code})
		h.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}
