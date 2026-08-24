package repository

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (r *ServiceRepo) PingDB(ctx context.Context) *m.Error {
	err := r.DB.Ping(ctx)
	if err != nil {
		return &m.Error{
			Error:   "Ping error.",
			Details: err.Error(),
			Code:    "500",
		}
	}
	return nil
}
