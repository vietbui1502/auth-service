package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/vietbui1502/auth-service/domain"
	"github.com/vietbui1502/auth-service/service"
)

func Start() {
	router := mux.NewRouter()

	//Create AuthRepositoryDB
	authRepositoryDB := domain.NewAuthRepositoryDB()

	//Create Auth service
	authService1 := service.NewAuthService(authRepositoryDB)

	//Create authentication handlers
	ah := Handlers{authService: authService1}

	//Define API
	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	//router.HandleFunc(path:"/auth/register", ah.NotImplementHandler).Methods(http.MethodPost)
	//router.HandleFunc(path:"/auth/verify", ah.Verify).Methods(http.MethodGet)

	//Start Service
	log.Fatal(http.ListenAndServe(":8889", router))
}
