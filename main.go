package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/** It's in /vagrant/ NOT in /home/vagrant wtf... */
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/balance", GetBalance).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	newPort := newPortfolio()
	newPort.printBalance("ETH")
	newPort.printBalance("dollar")
	json.NewEncoder(w).Encode(newPort.balances)
}
