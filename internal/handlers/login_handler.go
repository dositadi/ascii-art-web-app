package handlers

import (
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	at "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/web/templates/auth_templates"
)

func (s *Handler) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	err := at.LoginPageTemplate(w)
	if err != nil {
		err := h.ErrorToJson(m.Error{Error: err.Error, Details: err.Details, Code: err.Code})
		h.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}
