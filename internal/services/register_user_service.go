package services

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/google/uuid"
)

func (s *Service) RegisterUser(ctx context.Context, name, email, password string) *m.Error {
	if name == "" {
		err := m.Error{Error: h.EMPTY_NAME_FIELD, Details: h.EMPTY_EMAIL_FIELD_DETAIL, Code: h.EMPTY_PASSWORD_FIELD_CODE}
		return &err
	}
	if email == "" {
		err := m.Error{Error: h.EMPTY_EMAIL_FIELD, Details: h.EMPTY_EMAIL_FIELD_DETAIL, Code: h.EMPTY_EMAIL_FIELD_CODE}
		return &err
	}
	if !h.IsEmail(email) {
		err := m.Error{Error: h.BAD_EMAIL_FORMAT, Details: h.BAD_EMAIL_FORMAT_DETAIL, Code: h.BAD_EMAIL_FORMAT_CODE}
		return &err
	}
	if password == "" {
		err := m.Error{Error: h.EMPTY_PASSWORD_FIELD, Details: h.EMPTY_PASSWORD_FIELD_DETAIL, Code: h.EMPTY_PASSWORD_FIELD_CODE}
		return &err
	}

	hashedPassword, err := h.HashPassword(password)
	if err != nil {
		return err
	}

	var user = m.User{
		Id:             uuid.NewString(),
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
	}

	err1 := s.Repository.InsertUser(ctx, user)
	if err1 != nil {
		return err1
	}
	return nil
}
