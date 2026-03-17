package repository

import (
	"context"
	"database/sql"
	"errors"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *ServiceRepo) GetHashedPasswordIDAndName(ctx context.Context, email string) (string, string, string, *m.Error) {
	row := s.DB.QueryRowContext(ctx, h.GET_HPASS_ID_AND_NAME, email)

	var id, name, hashed_password string

	if err := row.Scan(&id, &name, &hashed_password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", "", &m.Error{
				Error:   h.UNAUTHORIZED_ERR,
				Details: h.UNAUTHORIZED_ERR_DETAIL,
				Code:    h.UNAUTHORIZED_ERR_CODE,
			}
		} else if errors.Is(err, sql.ErrConnDone) {
			return "", "", "", &m.Error{
				Error:   h.CONN_LOST_ERR,
				Details: h.CONN_LOST_ERR_DETAIL,
				Code:    h.CONN_LOST_ERR_CODE,
			}
		} else {
			return "", "", "", &m.Error{
				Error:   h.SERVER_ERR,
				Details: err.Error(),
				Code:    h.SERVER_ERR_CODE,
			}
		}
	}
	return id, name, hashed_password, nil
}
