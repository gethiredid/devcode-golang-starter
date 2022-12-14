package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get env
	PORT := getEnv("PORT", "3030")
	router := mux.NewRouter()

	// TODO: Uncomment baris kode dibawah ini
	// router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(map[string]string{"message": "Hello world"})
	// })

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
