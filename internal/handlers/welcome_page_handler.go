package handlers

import (
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	wt "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/web/templates/welcome_template"
)

func (s *Handler) WelcomePageHandler(w http.ResponseWriter, r *http.Request) {
	err := wt.WelcomePageTemplate(w)
	if err != nil {
		err := h.ErrorToJson(m.Error{Error: err.Error, Details: err.Details, Code: err.Code})
		h.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}
