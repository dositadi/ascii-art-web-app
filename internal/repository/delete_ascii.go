package repository

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (r *ServiceRepo) DeleteFromAscii(ctx context.Context, id string) *m.Error {
	exists, err1 := r.CheckIfAsciiExists(ctx, id)
	if err1 != nil {
		return err1
	}

	if !exists {
		return &m.Error{
			Error:   h.NOT_FOUND_ERR,
			Details: h.NOT_FOUND_DETAIL,
			Code:    h.NOT_FOUND_CODE,
		}
	}

	_, err2 := r.DB.Exec(ctx, h.DELETE_ASCII, id)
	if err2 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err2.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	return nil
}
