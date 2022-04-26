package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

//register

type Student struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

var data = []Student{
	Student{1, "budi", 5},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result []byte
		var err error
		id, err := strconv.ParseInt(r.FormValue("id")[0:], 10, 64)
		if err != nil {
			w.Write([]byte(`{"error": "internal error"}`))
		}
		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
