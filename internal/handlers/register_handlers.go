package handlers

import (
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/google/uuid"
)

/*
Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
*/

func (s *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	id := uuid.NewString()

	if name == "" {
		err := h.ErrorToJson(m.Error{Error: h.EMPTY_NAME_FIELD, Details: h.EMPTY_NAME_FIELD_DETAIL, Code: h.EMPTY_NAME_FIELD_CODE})
		h.ErrorResponse(w, err, http.StatusBadRequest)
		return
	} else if len(password) < 8 || password == "" {
		err := h.ErrorToJson(m.Error{Error: h.EMPTY_PASSWORD_FIELD, Details: h.EMPTY_PASSWORD_FIELD_DETAIL, Code: h.EMPTY_PASSWORD_FIELD_CODE})
		h.ErrorResponse(w, err, http.StatusBadRequest)
		return
	} else if email == "" {
		err := h.ErrorToJson(m.Error{Error: h.EMPTY_EMAIL_FIELD, Details: h.EMPTY_EMAIL_FIELD_DETAIL, Code: h.EMPTY_EMAIL_FIELD_CODE})
		h.ErrorResponse(w, err, http.StatusBadRequest)
		return
	} else if !h.IsEmail(email) {
		err := h.ErrorToJson(m.Error{Error: h.BAD_EMAIL_FORMAT, Details: h.BAD_EMAIL_FORMAT_DETAIL, Code: h.BAD_EMAIL_FORMAT_CODE})
		h.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	user := m.User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	ctx := r.Context()

	err := s.Service.RegisterUser(ctx, user)
	if err != nil {

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	http.Redirect(w, r, "/auth/login", http.StatusPermanentRedirect)
}

func (s *Handler) RenderRegisterPageHandler(w http.ResponseWriter, r *http.Request) {

}
