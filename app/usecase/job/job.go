package job

import (
	"dans/app/model"
	"dans/app/repository/job"
)

type JobUsecase interface {
	GetByID(id string) (model.Job, error)
	List(params model.ParamsJob) ([]model.Job, error)
}

type jobUsecase struct {
	jobRepository job.JobRepository
}

func NewJobUsecase(jobRepository job.JobRepository) JobUsecase {
	return &jobUsecase{
		jobRepository: jobRepository,
	}
}

func (j *jobUsecase) GetByID(id string) (model.Job, error) {
	return j.jobRepository.GetByID(id)
}

func (j *jobUsecase) List(params model.ParamsJob) ([]model.Job, error) {
	return j.jobRepository.List(params)
}
