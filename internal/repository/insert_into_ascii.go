package repository

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (r *ServiceRepo) InsertAscii(ctx context.Context, ascii m.Ascii) *m.Error {
	exists, err2 := r.CheckIfAsciiExists(ctx, ascii.Id)
	if err2 != nil {
		return err2
	}

	if exists {
		return &m.Error{
			Error:   h.CONFLICT_ERR,
			Details: "Ascii has been saved already.",
			Code:    h.CONFLICT_ERR_CODE,
		}
	}

	_, err3 := r.DB.Exec(ctx, h.INSERT_INTO_ASCII_TEXTS, ascii.Id, ascii.UserId, ascii.InputText, ascii.Font, ascii.AsciiText)
	if err3 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err3.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}
	return nil
}
