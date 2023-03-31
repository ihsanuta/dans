package repository

import (
	"dans/app/repository/job"
	"dans/app/repository/user"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	User user.UserRepository
	Job  job.JobRepository
}

func Init(db *gorm.DB) *Repository {
	repo := &Repository{
		User: user.NewUserRepository(
			db,
		),
		Job: job.NewJobRepository(
			"http://dev3.dansmultipro.co.id/api",
		),
	}
	return repo
}
