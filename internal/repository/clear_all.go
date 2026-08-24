package repository

import (
	"context"
	"log"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/jackc/pgx/v5"
)

func (r *ServiceRepo) ClearAll(ctx context.Context, user_id string) *m.Error {
	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	defer func() {
		if err3 := tx.Rollback(ctx); err3 != nil {
			log.Printf("%s", err3)
		}
	}()

	_, err2 := tx.Exec(ctx, h.CLEAR_ALL_USER_DATA, user_id)
	if err2 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err2.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	if err4 := tx.Commit(ctx); err4 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err4.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}
	return nil
}
