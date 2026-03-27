package repository

import (
	"context"
	"database/sql"
	"errors"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *ServiceRepo) GetAllUsersSavedAscii(ctx context.Context, user_id string) ([]m.Ascii, *m.Error) {
	rows, err := s.DB.QueryContext(ctx, "", user_id)
	if err != nil {
		if errors.Is(err, sql.ErrConnDone) {
			return nil, &m.Error{
				Error:   h.CONN_LOST_ERR,
				Details: err.Error(),
				Code:    h.CONN_LOST_ERR_CODE,
			}
		} else {
			return nil, &m.Error{
				Error:   h.CONN_LOST_ERR,
				Details: err.Error(),
				Code:    h.CONN_LOST_ERR_CODE,
			}
		}
	}

	asciiArts := []m.Ascii{}

	for rows.Next() {
		asciiArt := m.Ascii{}

		if err2 := rows.Scan(&asciiArt.Id, &asciiArt.InputText, &asciiArt.Font, &asciiArt.AsciiText, &asciiArt.CreatedAt); err2 != nil {
			if errors.Is(err2, sql.ErrConnDone) {
				return nil, &m.Error{
					Error:   h.CONN_LOST_ERR,
					Details: err2.Error(),
					Code:    h.CONN_LOST_ERR_CODE,
				}
			} else {
				return nil, &m.Error{
					Error:   h.CONN_LOST_ERR,
					Details: err.Error(),
					Code:    h.CONN_LOST_ERR_CODE,
				}
			}
		}

		asciiArts = append(asciiArts, asciiArt)
	}

	return asciiArts, nil
}
