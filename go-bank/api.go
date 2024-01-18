package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
}

type ApiError struct {
	Error string `json:"error"`
}

func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{listenAddr: listenAddr}
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHttpHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
		}
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	log.Println("Server running on port: " + s.listenAddr)

	router.HandleFunc("/account", makeHttpHandlerFunc(s.handleAccount))
	http.ListenAndServe(s.listenAddr, router)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.Header().Add("x-provider", "gabril-banking")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	default:
		return fmt.Errorf("method %s not allowed", r.Method)
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	}
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("Gabriel", "Peres")
	return WriteJson(w, http.StatusOK, account)
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
