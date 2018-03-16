package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// StartServer starts the HTTP server and configures all necessary routes.
func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/transaction/{year}/{month}", listHandler)
	r.HandleFunc("/api/transaction/{year}/{month}/budget", budgetHandler)
	r.HandleFunc("/api/transaction", postHandler).
		Methods("POST")
	// TODO ML Use gobindata later to pack all files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	port := ":8080"
	log.Println("Starting to listen on port", port)
	http.ListenAndServe(port, r)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List transactions handler called. vars=", mux.Vars(r))
	year, month, err := parseStandardFields(w, r)
	if err != nil {
		return
	}

	w.Header()["Content-Type"] = []string{"application/json"}
	ts := Load(year, month)
	enc := json.NewEncoder(w)
	enc.Encode(ts)
}

func budgetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Budget transactions handler called. vars=", mux.Vars(r))
	year, month, err := parseStandardFields(w, r)
	if err != nil {
		return
	}

	w.Header()["Content-Type"] = []string{"application/json"}
	ts := Load(year, month)
	budget := ComputeBudget(ts)
	enc := json.NewEncoder(w)
	enc.Encode(budget)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Post transaction handler called. vars=", vars)

	dec := json.NewDecoder(r.Body)
	var t Transaction
	dec.Decode(&t)
	w.WriteHeader(http.StatusOK)
	trans := NewTransaction(t.Description, t.Amount.StringFixed(2))
	Save(trans)
}

func parseStandardFields(w http.ResponseWriter, r *http.Request) (int, int, error) {
	vars := mux.Vars(r)
	year, err := strconv.Atoi(vars["year"])
	if err != nil {
		log.Println("Unable to parse year", err)
		w.WriteHeader(http.StatusBadRequest)
		return 0, 0, err
	}
	month, err := strconv.Atoi(vars["month"])
	if err != nil {
		log.Println("Unable to parse month", err)
		w.WriteHeader(http.StatusBadRequest)
		return 0, 0, err
	}

	return year, month, nil
}
