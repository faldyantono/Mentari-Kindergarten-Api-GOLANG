package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Student struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

var baseURL = "http://localhost:8080"
var data = []Student{
	Student{1, "budi", 5},
}

//get
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
func fetchUsers() ([]Student, error) {
	var err error
	var client = &http.Client{}
	var data []Student

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func fetchUser(ID string) (Student, error) {
	var err error
	var client = &http.Client{}
	var data Student

	var param = url.Values{}
	param.Set("id", ID)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
func main() {
	var users, err = fetchUser("1")
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("ID: %d\t Name: %s\t Age: %d\n", each.ID, each.Name, each.Age)
	}
	http.HandleFunc("/student", users)
	http.HandleFunc("/user", user)
	fmt.Println("Mentari Kindergarten School is asking for help to create student registration.")
	http.ListenAndServe(":8080", nil)
}
