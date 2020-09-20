package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Animal struct {
	ID       string    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Species  string    `json:"species,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type Location struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
}

var animals []Animal

func getAnimals(value http.ResponseWriter, res *http.Request) {
	json.NewEncoder(value).Encode(animals)
}
func getAnimal(value http.ResponseWriter, res *http.Request) {
	params := mux.Vars(res)

	for _, item := range animals {
		if item.ID == params["id"] {
			json.NewEncoder(value).Encode(item)
			return
		}
	}
	json.NewEncoder(value).Encode(&Animal{})
}
func CreateAnimal(value http.ResponseWriter, res *http.Request) {
	params := mux.Vars(res)
	var animal Animal
	_ = json.NewDecoder(res.Body).Decode(&animal)
	animal.ID = params["id"]
	animals = append(animals, animal)
	json.NewEncoder(value).Encode(animals)

}
func DeleteAnimal(value http.ResponseWriter, res *http.Request) {
	params := mux.Vars(res)
	for index, item := range animals {
		if item.ID == params["id"] {
			animals = append(animals[:index], animals[index+1:]...)
			break
		}
		json.NewEncoder(value).Encode(animals)
	}

}
func main() {
	//Enrutador
	router := mux.NewRouter()

	// data
	animals = append(animals, Animal{ID: "1", Name: "Koda", Species: "Bear", Location: &Location{Country: "USA", State: "Alaska"}})
	animals = append(animals, Animal{ID: "2", Name: "Bucky", Species: "Bear"})
	// endpoints
	router.HandleFunc("/animal", getAnimals).Methods("GET")
	router.HandleFunc("/animal/{id}", getAnimal).Methods("GET")
	router.HandleFunc("/animal/{id}", CreateAnimal).Methods("POST")
	router.HandleFunc("/animal/{id}", DeleteAnimal).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
