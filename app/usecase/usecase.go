package usecase

import (
	"dans/app/repository"
	"dans/app/usecase/job"
	"dans/app/usecase/user"
)

type Usecase struct {
	User user.UserUsecase
	Job  job.JobUsecase
}

func Init(repository *repository.Repository) *Usecase {
	uc := &Usecase{
		User: user.NewUserUsecase(
			repository.User,
		),
		Job: job.NewJobUsecase(
			repository.Job,
		),
	}
	return uc
}
