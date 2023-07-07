package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Err(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	r := mux.NewRouter()

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/webhook", handler).Methods(http.MethodPost)
	r.Use(mux.CORSMethodMiddleware(r))

	log.Println("Listening...")
	http.ListenAndServe(":8888", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// TEST: read json - taken from http endpt beat
		var jsonData = []json.RawMessage{}
		d := json.RawMessage{}
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			fmt.Println(err)
			http.Error(w, "error decoding resp obj", http.StatusBadRequest)
			return
		}

		jsonData = append(jsonData, d)
		n, err := json.Marshal(jsonData)
		if err != nil {
			Err(err)
		}

		// read out marshalled data from req body
		// fmt.Printf("Req body data type: %T\n", r.Body)
		fmt.Println("POST Was called!", string(n))
		// sendToWebhook(bytes.NewBuffer(n))

	} else {
		fmt.Println("That request is not allowed")
	}
}
