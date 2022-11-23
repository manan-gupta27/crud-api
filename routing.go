package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	
)

func HandlerRouting() {
	r := mux.NewRouter();
	r.HandleFunc("/employee",CreateEmployee).Methods("POST")
	r.HandleFunc("/employees",GetEmployees).Methods("GET")

	r.HandleFunc("/employees/{eid}",GetEmployeeByID).Methods("GET")
	
	r.HandleFunc("/employees/{eid}",UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{eid}",DeleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080",r))
}