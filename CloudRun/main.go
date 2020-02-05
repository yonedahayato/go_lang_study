package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", LoggingHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Printf("Defaulting to port %s\n", port)
	}
	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type JSONMessage struct {
	Message interface{} `json:"message"`
}

func LoggingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-Access-Token, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var m JSONMessage
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		fmt.Printf("json.NewDecoder: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if m.Message == nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	j, _ := json.Marshal(m)
	fmt.Println(string(j))
}
