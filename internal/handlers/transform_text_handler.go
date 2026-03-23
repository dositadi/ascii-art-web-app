package handlers

import (
	"net/http"
	"time"

	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Handler) TransformTextHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	text := r.FormValue(h.TEXT_KEY)
	banner := r.FormValue(h.BANNER_KEY)

	err := s.Service.TransformText(w, r, text, banner, start)
	if err != nil {
		s.Service.RenderAsciiArtPage(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}
