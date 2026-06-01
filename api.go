package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	jwt "github.com/golang-jwt/jwt/v5"
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

	router.Handle("/account",  withJWTAuth(makeHTTPHandleFunc(s.handleAccount)))
	router.Handle("/account/{id}",  makeHTTPHandleFunc(s.handleGetById))

	log.Println("Json Api server running on port ",s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w,r)	
	}
	if r.Method == "POST"{
		return s.handleCreateAccount(w,r)
	}
	if r.Method == "DELETE"{
		return  s.handleDeleteAccount(w,r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetById(w http.ResponseWriter, r *http.Request) error {
	// vars := mux.Vars(r)["id"]
	account := NewAccount("Sahal","Sabith")
	
	return WriteJSON(w,http.StatusOK,account)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}

	return WriteJSON(w,http.StatusOK,accounts)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(&createAccountReq); err != nil{
		return err
	}

	account := NewAccount(createAccountReq.FirstName,createAccountReq.LatsName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	tokenString,err := CreateJWT(account)

	if err != nil{
		return err
	}

	fmt.Println("JWT token :",tokenString)

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateJWT(account *Account) (string, error){
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"accountNumber":account.Number,
	}

	secret := "s4h41s4b1th"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func withJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling jwt auth middleware....")

		tokenStirng := r.Header.Get("x-jwt-token")

		_, err := validateJWT(tokenStirng)

		if err != nil {
			WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
			return 
		}



		handlerFunc(w,r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := "s4h41s4b1th"

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("Unexpected signing method : %v",token.Header["alg"])
		}
		return []byte(secret),nil
	})
}