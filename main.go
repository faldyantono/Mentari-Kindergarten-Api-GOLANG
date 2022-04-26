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
	//register POST student into the system http://{host}:{port}/student
	r.HandleFunc("/student", returnAllstudents)
	//Get student By ID http://{host}:{port}/student/{id}
	r.HandleFunc("/student/{id}", returnSingleStudent)
	log.Fatal(http.ListenAndServe(":10000", r))
}

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

var Students []Student

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MENTARI KINDERGARTEN HOMEPAGE")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllstudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllstudents")
	json.NewEncoder(w).Encode(Students)
}
func returnSingleStudent(w http.ResponseWriter, r *http.Request) {
	//Get student By ID http://{host}:{port}/student/{id}
	vars := mux.Vars(r)
	key := vars["id"]
	for _, student := range Students {
		if student.Id == key {
			json.NewEncoder(w).Encode(student)
		}
	}
}
func main() {
	Students = []Student{Student{Id: "1", Name: "budi", Age: 5}}
	fmt.Println("Mentari Kindergarten Database")
	handleRequests()
	//localhost:10000/home : MENTARI KINDERGARTEN HOMEPAGE
	//Register siswa
	//localhost:10000/student : [{"id":1,"name":"budi","age":5}]
}
