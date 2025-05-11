package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tungnt76/vulcanlabs/service"
)

func main() {
	cinemaService := service.NewCinemaService()
	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	r.HandleFunc("/configure", cinemaService.Configure).Methods("POST")
	r.HandleFunc("/available-seats", cinemaService.GetAvailableSeats).Methods("GET")
	r.HandleFunc("/reserve-seats", cinemaService.ReserveSeats).Methods("POST")
	r.HandleFunc("/cancel-seats", cinemaService.CancelSeats).Methods("POST")
	r.HandleFunc("/seats", cinemaService.ListSeats).Methods("GET")

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
