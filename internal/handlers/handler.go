package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

func SlowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Started slow request...")
	time.Sleep(10 * time.Second)
	fmt.Fprintln(w, "Finished slow request")
	fmt.Println("finished Slow request")
}
