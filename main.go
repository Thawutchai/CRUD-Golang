package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Object event
type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

//assige Object to allEvents
type allEvents []event

//set value to Object
var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
	{
		ID:          "2",
		Title:       "Advanced Golang Concepts",
		Description: "Com and learn about the advanced concepts that will make you an expert in GO",
	},
}

// create API, and route
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents)
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------- Check API Welcome ---------")
	fmt.Fprint(w)
	fmt.Fprint(w, "Welcome home!")
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------- Create Event ---------")
	var newEvent event

	//r.body = get body from postman
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

//get from id
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------- Get from ID ---------")
	//get parameter name = id
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

//get all
func getAllEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------- Get All ---------")
	json.NewEncoder(w).Encode(events)
}

//update event
func updateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------- Update by ID ---------")
	eventID := mux.Vars(r)["id"]
	var updateEvent event

	//read data from r.body use ioutil[import]
	reqBody, err := ioutil.ReadAll(r.Body)
	//nil == function empty
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and descripton only in order to update")
	}
	json.Unmarshal(reqBody, &updateEvent)

	//foreach update
	for i, singleEvent := range events {
		//check is update from ID
		if singleEvent.ID == eventID {
			singleEvent.Title = updateEvent.Title
			singleEvent.Description = updateEvent.Description
			events = append(events[:i], singleEvent)
			fmt.Println(singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

//delete
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------- Delete by ID ---------")
	eventID := mux.Vars(r)["id"]
	fmt.Println("ID : [" + eventID + "]")
	for i, singEvent := range events {
		if singEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, "The event sith ID %v has been delete successfully : [", eventID+"]")
		}
	}

}
