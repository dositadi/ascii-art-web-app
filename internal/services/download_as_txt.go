package services

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/google/uuid"
)

func (s *Service) DownloadAsTxt(w http.ResponseWriter, r *http.Request, text, font, id string) *m.Error {
	formattedAscii, err := s.FormatAscii(text, font)
	if err != nil {
		return err
	}

	downloadId := uuid.NewString()

	attachment := fmt.Sprintf("attachment; filename=ascii_forge_%s_%s.txt", text, downloadId)

	w.Header().Set("Content-Disposition", attachment)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(formattedAscii)))

	_, err2 := io.Copy(w, strings.NewReader(formattedAscii))
	if err2 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err2.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	err3 := s.Repository.UpdateAsciiOutputsTable(r.Context(), id)
	if err3 != nil {
		fmt.Println(err3.Details)
		return err3
	}
	return nil
}
