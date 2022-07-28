package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func URLRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "It works"}`))
}

func getAPIHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`"output": { "health":"OK"}`))
}

func myFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`"output": { "myFunction":"called"}`))
	var1 := mux.Vars(r)["var1"]
	fmt.Printf(var1)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", URLRoot).Methods(http.MethodGet)
	router.HandleFunc("/_health", getAPIHealth).Methods(http.MethodGet)
	router.HandleFunc("/path1/path2/{var1}", myFunction).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":3001", router))
}
