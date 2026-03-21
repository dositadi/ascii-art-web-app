package handlers

import "net/http"

func (s *Handler) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	err := s.Service.RenderHomePage(w, r)
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}
}
