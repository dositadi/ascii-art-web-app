package repository

import (
	"context"
	"database/sql"
	"errors"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

type ServiceRepo struct {
	DB *sql.DB
}

func ConstructNewRepo(db *sql.DB) *ServiceRepo {
	return &ServiceRepo{
		DB: db,
	}
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

func (r *ServiceRepo) CheckIfAsciiExists(ctx context.Context, id string) (bool, *m.Error) {
	var exists bool

	row := r.DB.QueryRowContext(ctx, "", id)

	if err := row.Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, &m.Error{
				Error:   h.CONFLICT_ERR,
				Details: "Ascii has been saved already.",
				Code:    h.CONFLICT_ERR_CODE,
			}
		} else {
			return false, &m.Error{
				Error:   h.CONFLICT_ERR,
				Details: "Ascii has been saved already.",
				Code:    h.CONFLICT_ERR_CODE,
			}
		}
	}
	return exists, nil
}
