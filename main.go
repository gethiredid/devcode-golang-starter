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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var newContact Contact
		err := json.NewDecoder(r.Body).Decode(&newContact)

		if err == nil {
			if err := DB.Create(&newContact).Error; err == nil {
				res := ResItem{
					Status:  "Success",
					Message: "Contact created",
					Item:    newContact,
				}

				response, _ := json.Marshal(res)

				w.Write(response)
			}

		}

	}).Methods("POST")

	// update contact
	router.HandleFunc("/contacts/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)

		var contact Contact
		json.NewDecoder(r.Body).Decode(&contact)

		DB.Where("id = ?", id).Updates(&contact)

		contact.Id = id

		res := ResItem{
			Status:  "Success",
			Message: "Contact updated",
			Item:    contact,
		}

		response, _ := json.Marshal(res)

		w.Write(response)

	}).Methods("PUT")

	// delete contact
	router.HandleFunc("/contacts/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		// input := map[string]string{"id": ""}
		var contact Contact

		json.NewDecoder(r.Body).Decode(&contact)

		DB.Delete(&contact, id)

		res := ResItem{
			Status:    "Success",
			Message:   "Contact deleted",
			DeletedId: id,
		}

		response, _ := json.Marshal(res)

		w.Write(response)

	}).Methods("DELETE")

	log.Println("API is running on port 3030")
	http.ListenAndServe(":"+port, router)
}
