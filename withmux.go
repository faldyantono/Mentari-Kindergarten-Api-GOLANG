package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//localhost:10000 bukan 8080

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/home", homePage)
	r.HandleFunc("/student", returnAllstudents)
	//Get student By ID http://{host}:{port}/student/{id}
	r.HandleFunc("/student/{id}", returnSingleStudent)
	//Create and update student
	r.HandleFunc("/student/{id}", createNewStudent).Methods("POST")
	//delete student
	r.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")
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

//CREATE (REGISTER) AND UPDATE STUDENT
func createNewStudent(w http.ResponseWriter, r *http.Request) {
	// CREATE
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
	var student Student
	json.Unmarshal(reqBody, &student)
	// update students{1} array to include
	Students = append(Students, student)
	json.NewEncoder(w).Encode(student)
}

//DELETING STUDENT
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, student := range Students {
		if student.Id == id {
			// updates Students array to remove the student
			Students = append(Students[:index], Students[index+1:]...)
		}
	}

}

func main() {
	Students = []Student{Student{Id: "1", Name: "budi kurniawan", Age: 5}}
	fmt.Println("Mentari Kindergarten Database")
	handleRequests()
	//localhost:10000/home : MENTARI KINDERGARTEN HOMEPAGE
	//Register siswa
	//localhost:10000/student : [{"id":1,"name":"budi","age":5}]
}
