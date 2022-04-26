package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/home", homePage)
	r.HandleFunc("/student", returnAllstudents)
	log.Fatal(http.ListenAndServe(":10000", r))
}

type Student struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

var Students = []Student{Student{ID: 1, Name: "budi", Age: 5}}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MENTARI KINDERGARTEN HOMEPAGE")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllstudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllstudents")
	json.NewEncoder(w).Encode(Students)
}

func main() {
	fmt.Println("Mentari Kindergarten Database")
	handleRequests()
	//localhost:10000/home : MENTARI KINDERGARTEN HOMEPAGE
	//Register siswa
	//localhost:10000/student : [{"id":1,"name":"budi","age":5}]

}
