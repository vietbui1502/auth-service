package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vietbui1502/auth-service/dto"
	"github.com/vietbui1502/auth-service/service"
)

type Handlers struct {
	authService service.AuthService
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	//log.Panicln("Enter login function")
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println("Error while decoding login request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	} else {
		token, err := h.authService.Login(loginRequest)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprintf(w, *token)
		}
	}
}
