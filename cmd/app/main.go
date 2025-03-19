package main

import (
	"adp_practice1/internal/handler/user_handler"
	"adp_practice1/internal/repository/user_repo"
	"adp_practice1/internal/usecase"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	userRepo := user_repo.NewInMemoryUserRepository()

	userUseCase := usecase.NewUserUsecase(userRepo)

	userHandler := user_handler.NewUserHandler(userUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUserHandler).Methods("GET")

	port := ":8080"
	fmt.Println("Server is running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
