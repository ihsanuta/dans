package job

import (
	"dans/app/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type JobRepository interface {
	List(params model.ParamsJob) ([]model.Job, error)
	GetByID(id string) (model.Job, error)
}

type jobRepository struct {
	URL string
}

func NewJobRepository(URL string) JobRepository {
	return &jobRepository{
		URL: URL,
	}
}

func (j *jobRepository) List(params model.ParamsJob) ([]model.Job, error) {
	jobs := []model.Job{}
	client := http.Client{
		Timeout: time.Minute * 15,
	}

	baseURL, _ := url.Parse(fmt.Sprintf("%s%s", j.URL, "/recruitment/positions.json"))
	p := url.Values{}

	if params.Description != "" {
		p.Add("description", params.Description)
	}

	if params.Type != "" {
		p.Add("type", params.Type)
	}

	if params.Location != "" {
		p.Add("location", params.Location)
	}

	if params.Page > 0 {
		p.Add("page", fmt.Sprintf("%d", params.Page))
	}

	baseURL.RawQuery = p.Encode()
	log.Println(baseURL.String())
	req, err := http.NewRequest(http.MethodGet, baseURL.String(), nil)
	if err != nil {
		return jobs, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return jobs, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&jobs)
	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (j *jobRepository) GetByID(id string) (model.Job, error) {
	job := model.Job{}
	client := http.Client{
		Timeout: time.Minute * 15,
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s%s", j.URL, "/recruitment/positions/", id), nil)
	if err != nil {
		return job, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return job, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&job)
	if err != nil {
		return job, err
	}

	return job, nil
}
