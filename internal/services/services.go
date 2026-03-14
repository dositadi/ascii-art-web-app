package services

import (
	r "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/internal/repository"
)

type Service struct {
	ServiceRepo r.ServiceRepo
}

func ConstructNewService(serviceRepo r.ServiceRepo) *Service {
	return &Service{
		ServiceRepo: serviceRepo,
	}
}
