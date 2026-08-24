package repository

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (r *ServiceRepo) UpdateAsciiOutputsTable(ctx context.Context, id string) *m.Error {
	_, err2 := r.DB.Exec(ctx, h.UPDATE_ASCII_DOWNLOAD_AS_TXT, id)
	if err2 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err2.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	return nil
}
