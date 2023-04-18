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

func checkHeader(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, ok := r.Header["Authorization"]
		if !ok {
			w.WriteHeader(406)
			return
		}
		thetoken := strings.Split(token[0], " ")[1]
		fmt.Printf("\n%q\n\n", thetoken)
		f(w, r)
	}
}
func main() {
	http.HandleFunc("/decode", checkHeader(func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Printf("Username: %q , Email: %q, Password: %q \n", user.Username, user.Email, user.Password)

		w.WriteHeader(200)
	}))

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			Username: "ndzhwr",
			Email:    "ndizihiweregis06@gmail.com",
			Password: "thequickbrownfox",
		}
		json.NewEncoder(w).Encode(&user)
		w.WriteHeader(200)
	})

	http.ListenAndServe(":8080", nil)
}
