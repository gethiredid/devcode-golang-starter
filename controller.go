package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Contact struct {
	Id           int64  `gorm:"primaryKey" json:"id,omitempty"`
	Full_name    string `gorm:"type:varchar(255)" json:"full_name,omitempty"`
	Phone_number string `gorm:"type:varchar(255)" json:"phone_number,omitempty"`
	Email        string `gorm:"type:varchar(255)" json:"email,omitempty"`
}

type ResItems struct {
	Status  string    `json:"status"`
	Message string    `json:"message,omitempty"`
	Items   []Contact `json:"data"`
}

type ResItem struct {
	Status    string  `json:"status"`
	Message   string  `json:"message,omitempty"`
	Item      Contact `json:"data,omitempty"`
	DeletedId int64   `json:"deletedId,omitempty"`
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

func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	// TODO: edit data (full_name/email/phone_number) pada database berdasarkan id nya

	var contact Contact

	res := ResItem{
		Status:  "Success",
		Message: "Contact updated",
		Item:    contact,
	}

	response, _ := json.Marshal(res)

	w.Write(response)
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var contact Contact

	// TODO: hapus data pada database berdasarkan id nya

	res := ResItem{
		Status:    "Success",
		Message:   "Contact deleted",
		DeletedId: id,
	}

	response, _ := json.Marshal(res)

	w.Write(response)

}
