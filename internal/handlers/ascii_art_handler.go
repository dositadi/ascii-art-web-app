package handlers

import (
	"net/http"
)

func (s *Handler) AsciiArtPageHandler(w http.ResponseWriter, r *http.Request) {
	err := s.Service.RenderAsciiArtPage(w, r)
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
	}
}
