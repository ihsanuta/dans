package main

import (
	"dans/app/repository"
	"dans/app/usecase"
	"dans/handler"
	"dans/lib/mysql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// Business Layer
	repo *repository.Repository
	uc   *usecase.Usecase

	h handler.Handler
)

func main() {
	// konek to mysql
	db := mysql.GetMysqlConnection()

	// Business layer Initialization
	repo = repository.Init(
		db,
	)
	uc = usecase.Init(repo)
	h = handler.Init(uc)
}
