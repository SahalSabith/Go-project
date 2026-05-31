package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON( w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.Handle("/account",  makeHTTPHandleFunc(s.handleAccount))
	router.Handle("/account/{id}",  makeHTTPHandleFunc(s.handleAccount))

	log.Println("Json Api server running on port ",s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGet(w,r)	
	}
	if r.Method == "POST"{
		return s.handleCreate(w,r)
	}
	if r.Method == "DELETE"{
		return  s.handleDeleteAccount(w,r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGet(w http.ResponseWriter, r *http.Request) error {
	// vars := mux.Vars(r)["id"]
	account := NewAccount("Sahal","Sabith")
	
	return WriteJSON(w,http.StatusOK,account)
}

func (s *APIServer) handleCreate(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}