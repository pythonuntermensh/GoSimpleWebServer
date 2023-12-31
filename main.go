package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserDetails struct {
	Login    string `json:"login"    validate:"required"`
	Email    string `json:"email"    validate:"required"`
	Password string `json:"password"    validate:"required"`
}

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK")
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Supports only POST method")
		return
	}

	decoder := json.NewDecoder(req.Body)

	var data UserDetails
	if err := decoder.Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid JSON")
		return
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missed fields in JSON")
		return
	}

	log.Println(data.Email, data.Login, data.Password)

	fmt.Fprintf(w, "Logined successfully by:\n")
	fmt.Fprintf(w, "\tLogin = %s\n", data.Login)
	fmt.Fprintf(w, "\tEmail = %s\n", data.Email)
	fmt.Fprintf(w, "\tPassword = %s\n", data.Password)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/health", healthCheckHandler)
	log.Fatal((http.ListenAndServe(":80", nil)))
}
