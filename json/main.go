package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Middleware func(f http.HandlerFunc) http.HandlerFunc

func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				w.WriteHeader(405)
			}
			f(w, r)
		}
	}
}

func Token() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			token, ok := r.Header["Authorization"]
			if !ok {
				w.WriteHeader(406)
				return
			}
			thetoken := strings.Split(token[0], " ")[1]
			fmt.Printf("Token: %q\n", thetoken)
			f(w, r)
		}
	}
}

func ChainMiddlewares(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Decode(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Printf("Username: %q , Email: %q, Password: %q \n", user.Username, user.Email, user.Password)

	w.WriteHeader(200)
}

func Encode(w http.ResponseWriter, r *http.Request) {
	user := User{
		Username: "ndzhwr",
		Email:    "ndizihiweregis06@gmail.com",
		Password: "thequickbrownfox",
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&user)
}

func main() {
	http.HandleFunc("/decode", ChainMiddlewares(Decode, Token(), Method("POST")))
	http.HandleFunc("/encode", ChainMiddlewares(Encode, Token(), Method("GET")))
	http.ListenAndServe(":8080", nil)
}
