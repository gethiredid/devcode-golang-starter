package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Contact struct {
	Id           int    `json:"id"`
	Full_name    string `json:"full_name"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
}

type ResItems struct {
	Status  string    `json:"status"`
	Message string    `json:"message,omitempty"`
	Items   []Contact `json:"data"`
}

type ResItem struct {
	Status  string  `json:"status"`
	Message string  `json:"message,omitempty"`
	Item    Contact `json:"data"`
}

func main() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	if port == "" {
		port = "3030"
	}

	var contacts []Contact

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{"message": "Hello world"})
	})

	// get all contacts
	router.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		res := ResItems{
			Status: "Success",
			Items:  contacts,
		}

		response, _ := json.Marshal(res)

		w.Write(response)
	}).Methods("GET")

	// create contacts
	router.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var newContact Contact
		err := json.NewDecoder(r.Body).Decode(&newContact)
		newContact.Id = len(contacts) + 1

		if err == nil {
			contacts = append(contacts, newContact)

			res := ResItem{
				Status:  "Success",
				Message: "Contact created",
				Item:    newContact,
			}

			response, _ := json.Marshal(res)

			w.Write(response)
		}

	}).Methods("POST")

	log.Println("API is running on port 3030")
	http.ListenAndServe(":"+port, router)
}
