package repository

import (
	"context"
	"fmt"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/jackc/pgx/v5"
)

func (r *ServiceRepo) GetAllUsersSavedAscii(ctx context.Context, user_id string, limit, offset int, font string) ([]m.Ascii, *m.Error) {
	var rows pgx.Rows
	var err error

	if font == "" {
		rows, err = r.DB.Query(ctx, h.GET_ALL_USER_SAVED_ASCII, user_id, limit, offset)
		if err != nil {
			fmt.Println(err.Error())
			return nil, &m.Error{
				Error:   h.CONN_LOST_ERR,
				Details: err.Error(),
				Code:    h.CONN_LOST_ERR_CODE,
			}
		}
	} else {
		rows, err = r.DB.Query(ctx, h.FILTER_ASCII, user_id, font, limit, offset)
		if err != nil {
			fmt.Println(err.Error())
			return nil, &m.Error{
				Error:   h.SERVER_ERR,
				Details: err.Error(),
				Code:    h.SERVER_ERR_CODE,
			}
		}
	}

	defer rows.Close()

	asciiArts := []m.Ascii{}

	for rows.Next() {
		var asciiArt m.Ascii

		if err2 := rows.Scan(&asciiArt.Id, &asciiArt.InputText, &asciiArt.Font, &asciiArt.AsciiText, &asciiArt.CreatedAt); err2 != nil {
			fmt.Println("entered 2", err2.Error())
			return nil, &m.Error{
				Error:   h.CONN_LOST_ERR,
				Details: err2.Error(),
				Code:    h.CONN_LOST_ERR_CODE,
			}
		}
		asciiArts = append(asciiArts, asciiArt)
	}

	if err3 := rows.Err(); err3 != nil {
		fmt.Println(err3.Error())
		return nil, &m.Error{
			Error:   h.SERVER_ERR,
			Details: err3.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	return asciiArts, nil
}
