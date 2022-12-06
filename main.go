package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	log.Println("API is running on port 3030")
	http.ListenAndServe(":"+port, router)
}
