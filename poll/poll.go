package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Import API KEY from .env file
	API_KEY := os.Getenv("API_KEY")

	const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"
	const TRANSCRIPT_ID = "r427robbnf-6649-4765-b5db-c0ef3fb13d7f"
	const POLLING_URL = TRANSCRIPT_URL + "/" + TRANSCRIPT_ID

	client := &http.Client{}

	req, _ := http.NewRequest("GET", POLLING_URL, nil)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", API_KEY)

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

}
