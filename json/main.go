package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		fmt.Printf("Username: %q , Email: %q, Password: %q \n", user.Username, user.Email, user.Password)

		w.WriteHeader(200)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			Username: "ndzhwr",
			Email:    "ndizihiweregis06@gmail.com",
			Password: "thequickbrownfox",
		}
		json.NewEncoder(w).Encode(&user)
		w.WriteHeader(200)
	})

	http.ListenAndServe(":8080",nil);
}
