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
	var newContact Contact
	json.NewDecoder(r.Body).Decode(&newContact)

	// TODO: validasi data terlebih dahulu sebelum menambah data kedalam database

	if err := DB.Create(&newContact).Error; err == nil {
		res := ResItem{
			Status:  "Success",
			Message: "Contact created",
			Item:    newContact,
		}

		response, _ := json.Marshal(res)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func getAllContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var contacts []Contact

	if err := DB.Find(&contacts).Error; err == nil {
		res := ResItems{
			Status: "Success",
			Items:  contacts,
		}

		response, _ := json.Marshal(res)
		w.Write(response)
	}
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	// TODO: validasi data terlebih dahulu sebelum mengedit data yang ada dalam database

	var updateData Contact
	json.NewDecoder(r.Body).Decode(&updateData)

	// update contact
	DB.Where("id = ?", id).Updates(&updateData)
	updateData.Id = id

	res := ResItem{
		Status:  "Success",
		Message: "Contact updated",
		Item:    updateData,
	}

	response, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	// TODO: validasi data terlebih dahulu sebelum mengahapus data yang ada dalam database

	var removeContact Contact
	DB.Delete(&removeContact, id)

	res := ResItem{
		Status:    "Success",
		Message:   "Contact deleted",
		DeletedId: id,
	}

	response, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
