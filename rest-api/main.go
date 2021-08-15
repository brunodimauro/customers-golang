package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	CustomerId int64
	FirstName  string
	LastName   string
	Address    *Address
}

type Address struct {
	City  string
	State string
}

var customers []Customer

func main() {
	//just for now / later I'll add a NoSQL database
	customers = append(customers, Customer{CustomerId: 1, FirstName: "Jack", LastName: "Pearson", Address: &Address{City: "Pittsburgh", State: "Pennsylvania"}})
	customers = append(customers, Customer{CustomerId: 2, FirstName: "Rebecca", LastName: "Malone", Address: &Address{City: "Pittsburgh", State: "Pennsylvania"}})
	customers = append(customers, Customer{CustomerId: 3, FirstName: "Kate", LastName: "Pearson", Address: &Address{City: "Los Angeles", State: "California"}})
	customers = append(customers, Customer{CustomerId: 4, FirstName: "Kevin", LastName: "Pearson", Address: &Address{City: "New York", State: "New York"}})
	customers = append(customers, Customer{CustomerId: 5, FirstName: "Randall", LastName: "Pearson", Address: &Address{City: "Philadelphia", State: "Pennsylvania"}})

	//Define endpoints
	router := mux.NewRouter()
	router.HandleFunc("/customer", GetCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", GetCustomer).Methods("GET")
	router.HandleFunc("/customer/{id}", CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", DeleteCustomer).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramId, _ := strconv.ParseInt(params["id"], 10, 64)
	for _, item := range customers {
		if item.CustomerId == paramId {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Customer{})
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var customer Customer
	_ = json.NewDecoder(r.Body).Decode(&customer)
	paramId, _ := strconv.ParseInt(params["id"], 10, 64)
	customer.CustomerId = paramId
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customers)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// Cast to Int64
	paramId, _ := strconv.ParseInt(params["id"], 10, 64)
	for index, item := range customers {
		if item.CustomerId == paramId {
			customers = append(customers[:index], customers[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(customers)
	}
}
