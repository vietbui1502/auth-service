package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	

	//Define API
	router.HandleFunc(path:"/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc(path:"/auth/register", ah.NotImplementHandler).Methods(http.MethodPost)
	router.HandleFunc(path:"/auth/verify", ah.Verify).Methods(http.MethodGet)

	//Start Service
	log.Fatal(http.ListenAndServe(":8888", router))
}

func getDbClient() *sqlx.DB {
	dbClient, err := sqlx.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		logger.Error(err.Error())
	}
	// See "Important settings" section.
	dbClient.SetConnMaxLifetime(time.Minute * 3)
	dbClient.SetMaxOpenConns(10)
	dbClient.SetMaxIdleConns(10)
	return dbClient
}
