package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"todo-app/internal/db"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
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

// created todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	database := db.InitDb()
	defer database.Close()

	var t Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Got todo: %s (done=%v)\n", t.Task, t.Done)

	res, err := database.Exec(
		"INSERT INTO todos(task, done) VALUES(?, ?)",
		t.Task, t.Done,
	)

	if err != nil {
		http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	t.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

// get to handler
func GetTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	database := db.InitDb()
	defer database.Close()

	rows, err := database.Query("SELECT id, task, done FROM todos")
	if err != nil {
		http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Task, &t.Done); err != nil {
			http.Error(w, "scan error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "rows error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

//update todo to done or undone
//delete todo
