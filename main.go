package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

var students []Student

const (
	PORT int = 4444
	TAG string = "/api/v1/students"
)

func main(){
	http.HandleFunc(fmt.Sprintf("%s", TAG), getData)
	http.HandleFunc(fmt.Sprintf("%s/add", TAG), addData)

	log.Printf("listening on http://0.0.0.0:%d", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

func getData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func addData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var newStudent Student
	newStudent.Id = len(students) + 1
	newStudent.Name = "Tambahhhhh"
	students = append(students, newStudent)
	log.Println(r)
	json.NewEncoder(w).Encode(newStudent)
}