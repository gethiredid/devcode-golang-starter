package main

import (
	"encoding/json"
	"net/http"
)

type Contact struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	Full_name    string `gorm:"type:varchar(255)" json:"full_name"`
	Phone_number string `gorm:"type:varchar(255)" json:"phone_number"`
	Email        string `gorm:"type:varchar(255)" json:"email"`
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

var DB = db()

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var newContact Contact
	json.NewDecoder(r.Body).Decode(&newContact)

	// TODO: simpan data dari request body kedalam database

	res := ResItem{
		Status:  "Success",
		Message: "Contact created",
		Item:    newContact,
	}

	response, _ := json.Marshal(res)

	w.Write(response)

}

func getAllContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var contacts []Contact

	// TODO: ambil semua data kontak dari database

	res := ResItems{
		Status: "Success",
		Items:  contacts,
	}

	response, _ := json.Marshal(res)
	w.Write(response)
}
