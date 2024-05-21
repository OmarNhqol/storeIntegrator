package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", webhookHandler)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

type Data struct {
	AccessToken  string `json:"access_token"`
	Expires      int64  `json:"expires"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

// RequestBody represents the structure of the incoming request body.
type RequestBody struct {
	Event     string `json:"event"`
	Merchant  int64  `json:"merchant"`
	CreatedAt string `json:"created_at"`
	Data      Data   `json:"data"`
}

// webhookHandler handles incoming POST requests to the /webhook endpoint.
func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// Check that the request method is POST.
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into the RequestBody struct.
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Print the access token to the console.
	fmt.Println("Access Token:", reqBody.Data.AccessToken)

	// Send a response back to the client.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Access token received"))
}
