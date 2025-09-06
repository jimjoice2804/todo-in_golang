package handlers

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := Person{Name: "jimmy", Age: 25}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

// func SlowHandler(w http.ResponseWriter, r *http.Request){

// }
