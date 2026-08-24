package repository

import (
	"context"
	"fmt"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/jackc/pgx/v5"
)

func (r *ServiceRepo) GetTableLenght(ctx context.Context, user_id, font string) (int, *m.Error) {
	var row pgx.Row

	if font == "" {
		row = r.DB.QueryRow(ctx, h.GET_TABLE_LENGHT, user_id)
	} else {
		row = r.DB.QueryRow(ctx, h.GET_TABLE_LENGHT_WITH_FONT, user_id, font)
	}

	var length int

	if err := row.Scan(&length); err != nil {
		fmt.Println(err.Error())
		return 0, &m.Error{
			Error:   h.SERVER_ERR,
			Details: err.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}
	return length, nil
}
