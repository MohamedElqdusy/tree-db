package main

import (
	"fmt"
	"net/http"
	"tree/config"
	"tree/db"
	"tree/logger"
	"tree/service"
)

const appName = "Tree-service"

func main() {

	log := logger.NewLogger()
	logger.InitLogger(log)
	defer logger.Sync()

	initiatPostgre()

	//  create a new *router instance
	router := service.NewRouter()
	logger.Infof("Start Listen And Serve on 5475")
	if err := http.ListenAndServe(":5475", router); err != nil {
		logger.Fatal(err)
	}
}

func initiatPostgre() {
	pc := config.IniatilizePostgreConfig()
	postgersAddress := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pc.PostgresHost, pc.PostgresPort, pc.PostgresUser, pc.PostgresPassword, pc.PostgresDataBase)
	repository, err := db.NewPostgre(postgersAddress)
	logger.Error(err)
	db.SetRepository(repository)
}
