package services

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) LoginUser(ctx context.Context, email *string, password string) (m.ActiveUser, *m.Error) {
	id, name, hashed_password, err := s.Repository.GetHashedPasswordIDAndName(ctx, nil, email)
	if err != nil {
		return m.ActiveUser{}, err
	}

	err2 := h.ComparePasswordAndHash(hashed_password, password)
	if err2 != nil {
		return m.ActiveUser{}, err2
	}
	return m.ActiveUser{Id: id, Name: name, Email: *email}, nil
}
