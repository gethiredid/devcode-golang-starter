package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get env
	PORT := getEnv("PORT", "3030")
	MYSQL_USER := getEnv("MYSQL_USER", "root")
	MYSQL_PASSWORD := getEnv("MYSQL_PASSWORD", "")
	MYSQL_HOST := getEnv("MYSQL_HOST", "127.0.0.1")
	MYSQL_PORT := getEnv("MYSQL_PORT", "3306")
	MYSQL_DBNAME := getEnv("MYSQL_DBNAME", "dbname")

	var DB *gorm.DB
	conString := MYSQL_USER + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DBNAME
	db, err := gorm.Open(mysql.Open(conString))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Contact{})
	DB = db

	router := mux.NewRouter()

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

	log.Println("API is running on port " + PORT)
	http.ListenAndServe(":"+PORT, router)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}
	return defaultValue
}
