package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Person struct {
	Id        string  `json:"id, omitempty"`
	FirstName string  `json:"first_name, omitempty"`
	LastName  string  `json:"last_name, omitempty"`
	Address   Address `json:"address, omitempty"`
}

type Address struct {
	City string `json:"city, omitempty"`
	Code uint32 `json:"code, omitempty"`
}

var people []Person

func main() {
	people = append(people, Person{"1", "Cristiano", "Ronaldo", Address{"Palermo", 56042}})
	people = append(people, Person{"2", "John", "James", Address{"Reggio Calabria", 21264}})
	people = append(people, Person{"3", "Asaroth", "Dimitri", Address{"Roma", 85032}})
	people = append(people, Person{"4", "Scan", "Gengie", Address{"Milano", 11811}})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")

	log.Println(http.ListenAndServe(":8000", router))
}

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	for _, value := range people {
		if value.Id == params["id"] {
			json.NewEncoder(writer).Encode(value)
			return
		}
	}

	json.NewEncoder(writer).Encode(Person{})
}
func GetPeople(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(people)
}
