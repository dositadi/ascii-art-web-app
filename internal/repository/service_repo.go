package repository

import "database/sql"

type ServiceRepo struct {
	DB *sql.DB
}

func ConstructNewRepo(db *sql.DB) *ServiceRepo {
	return &ServiceRepo{
		DB: db,
	}
}
