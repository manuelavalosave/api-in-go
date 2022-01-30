package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetNameEndpoint(w http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	json.NewEncoder(w).Encode("Hola " + param["name"])

}
func main() {
	router := mux.NewRouter()

	//endpoinsts
	router.HandleFunc("/empleado/hola/{name}", GetNameEndpoint).Methods("GET")

	//start server
	log.Fatal(http.ListenAndServe(":3000", router))
}
