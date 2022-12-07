package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func main() {
	var DB *gorm.DB
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/go_contacts"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Contact{})
	DB = db

	port := os.Getenv("PORT")
	router := mux.NewRouter()

	if port == "" {
		port = "3030"
	}

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{"message": "Hello world"})
	})

	// get all contacts
	router.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
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
	}).Methods("GET")

	// create contacts
	router.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {

		// validate body request
		var newContact Contact
		json.NewDecoder(r.Body).Decode(&newContact)
		if len(newContact.Email) <= 0 || len(newContact.Full_name) <= 0 || len(newContact.Phone_number) <= 0 {
			res := map[string]string{"message": "full_name, phone_number, and email is required", "status": "Failed"}
			response, _ := json.Marshal(res)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		// check if data already exist
		var contact Contact
		DB.Where("email = ?", newContact.Email).Where("full_name = ?", newContact.Full_name).Where("phone_number", newContact.Phone_number).First(&contact)
		if len(contact.Email) > 0 || len(contact.Full_name) > 0 || len(contact.Phone_number) > 0 {
			res := map[string]string{"message": "full_name, phone_number, and email is duplicate", "status": "Failed"}
			response, _ := json.Marshal(res)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

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

	}).Methods("POST")

	// update contact
	router.HandleFunc("/contacts/{id}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)

		// validate body request
		var updateData Contact
		json.NewDecoder(r.Body).Decode(&updateData)
		if len(updateData.Email) <= 0 && len(updateData.Full_name) <= 0 && len(updateData.Phone_number) <= 0 {
			res := map[string]string{"message": "no contact updated", "status": "Failed"}
			response, _ := json.Marshal(res)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

		// check if contact not found
		var contact Contact
		if err := DB.First(&contact, id).Error; err != nil {
			message := "Contact with id " + strconv.Itoa(int(id)) + " is not found"
			res := map[string]string{"message": message, "status": "Failed"}
			response, _ := json.Marshal(res)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

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

	}).Methods("PUT")

	// delete contact
	router.HandleFunc("/contacts/{id}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		// check if contact not found
		var contact Contact
		if err := DB.First(&contact, id).Error; err != nil {
			message := "Contact with id " + strconv.Itoa(int(id)) + " is not found"
			res := map[string]string{"message": message, "status": "Failed"}
			response, _ := json.Marshal(res)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(response)
			return
		}

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

	}).Methods("DELETE")

	log.Println("API is running on port 3030")
	http.ListenAndServe(":"+port, router)
}
