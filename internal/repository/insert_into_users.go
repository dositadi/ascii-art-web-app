package repository

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (r *ServiceRepo) InsertUser(ctx context.Context, user m.User) *m.Error {
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

	_, err2 := r.DB.Exec(ctx, h.INSERT_INTO_USERS, user.Id, user.Name, user.Email, user.HashedPassword)
	if err2 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err2.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}
	return nil
}
