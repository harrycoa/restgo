package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"./handlers"
	"./models"
)

func main(){
	// generar el mux
	mux := mux.NewRouter()
	models.SetDefaultUser()
	// rutas
	mux.HandleFunc("/api/v1/users/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/v1/users/", handlers.CreateUsers).Methods("POST")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/v1/users/{id:[0-9]+}", handlers.DeleteUsers).Methods("DELETE")


	log.Println("El servidor esta a la escucha en el puerto 4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
