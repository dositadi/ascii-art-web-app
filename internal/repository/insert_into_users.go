package repository

import (
	"context"
	"database/sql"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (r *ServiceRepo) InsertUser(ctx context.Context, user m.User) *m.Error {
	tx, err1 := r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err1 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err1.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	exists, err := r.CheckIfUserExists(ctx, user.Email)
	if err != nil {
		return err
	}

	if exists {
		return &m.Error{
			Error:   h.CONFLICT_ERR,
			Details: "User exists already.",
			Code:    h.CONFLICT_ERR_CODE,
		}
	}

	_, err2 := tx.ExecContext(ctx, h.INSERT_INTO_USERS, user.Id, user.Name, user.Email, user.HashedPassword)
	if err2 != nil {
		if err3 := tx.Rollback(); err3 != nil {
			return &m.Error{
				Error:   h.SERVER_ERR,
				Details: err3.Error(),
				Code:    h.SERVER_ERR_CODE,
			}
		}
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err2.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	if err4 := tx.Commit(); err4 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err4.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}
	return nil
}

func (r *ServiceRepo) CheckIfUserExists(ctx context.Context, email string) (bool, *m.Error) {
	var exists bool
	row := r.DB.QueryRowContext(ctx, h.CHECK_USER_EXISTS, email)

	if err := row.Scan(&exists); err != nil {
		return false, &m.Error{
			Error:   h.SERVER_ERR,
			Details: err.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	return exists, nil
}
