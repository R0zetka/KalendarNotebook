package main

import (
	"Kalendar/internal/app/endpoint/calendar"
	"Kalendar/internal/app/endpoint/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/tasklist/user", user.CreateAccount).Methods("POST")
	r.HandleFunc("/api/v1/tasklist/user/{personEmail}/{personPassword}", user.GetPerson).Methods("GET")

	r.HandleFunc("/api/v1/tasklist/user/{personEmail}/{personPassword}/reminder", calendar.CreateReminder).Methods("POST")
	r.HandleFunc("/api/v1/tasklist/user/{personEmail}/{personPassword}/reminder/{reminderID}", calendar.GetReminder).Methods("GET")
	r.HandleFunc("/api/v1/tasklist/user/{personEmail}/{personPassword}/reminder/{reminderID}/status", calendar.UpdateStatusReminder).Methods("PUT")
	r.HandleFunc("/api/v1/tasklist/user/{personEmail}/{personPassword}/reminder/{reminderID}", calendar.UpdateReminder).Methods("PUT")
	r.HandleFunc("/api/v1/tasklist/user/{personEmail}/{personPassword}/reminder/{reminderID}", calendar.DeleteReminder).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
