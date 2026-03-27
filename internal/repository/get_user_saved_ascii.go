package repository

import (
	"context"
	"fmt"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *ServiceRepo) GetAllUsersSavedAscii(ctx context.Context, user_id string) ([]m.Ascii, *m.Error) {
	rows, err := s.DB.QueryContext(ctx, h.GET_ALL_USER_SAVED_ASCII, user_id)
	fmt.Println(user_id)
	if err != nil {
		return nil, &m.Error{
			Error:   h.CONN_LOST_ERR,
			Details: err.Error(),
			Code:    h.CONN_LOST_ERR_CODE,
		}
	}

	defer rows.Close()

	asciiArts := []m.Ascii{}

	for rows.Next() {
		var asciiArt m.Ascii

		if err2 := rows.Scan(&asciiArt.Id, &asciiArt.InputText, &asciiArt.Font, &asciiArt.AsciiText, &asciiArt.CreatedAt); err2 != nil {
			return nil, &m.Error{
				Error:   h.CONN_LOST_ERR,
				Details: err2.Error(),
				Code:    h.CONN_LOST_ERR_CODE,
			}
		}

		asciiArts = append(asciiArts, asciiArt)
		//fmt.Println("FORMATTED: ", asciiArt.CreatedAt.Format("02 Jan 2006 15:04"))
	}

	if err3 := rows.Err(); err3 != nil {
		return nil, &m.Error{
			Error:   h.SERVER_ERR,
			Details: err3.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	return asciiArts, nil
}
