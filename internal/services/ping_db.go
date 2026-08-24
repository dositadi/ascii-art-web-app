package services

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (s *Service) CheckDBHealth(ctx context.Context) *m.Error {
	if err := s.Repository.PingDB(ctx); err != nil {
		return err
	}
	return nil
}
