package handlers

import (
	"fmt"
	"net/http"
	"strings"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
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
			s.Service.RenderSignupPage(w, r, &err2.Details)
			return
		} else if err2.Error == h.SERVER_ERR {
			s.Service.RenderSignupPage(w, r, &err2.Details) // Change this to the error page
			return
		} else {
			s.Service.RenderSignupPage(w, r, &err2.Details) // change this to the error page
			return
		}
	}

	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

func (s *Handler) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	err := s.Service.RenderSignupPage(w, r, nil)
	if err != nil {
		err := h.ErrorToJson(m.Error{Error: err.Error, Details: err.Details, Code: err.Code})
		h.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}
