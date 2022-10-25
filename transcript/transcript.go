package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	const AUDIO_URL = "https://cdn.assemblyai.com/upload/019c9d26-42fe-4efa-8264-f84b7ae1df26"
	const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"

	// Prepare json data
	values := map[string]string{"audio_url": AUDIO_URL}
	jsonData, _ := json.Marshal(values)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", TRANSCRIPT_URL, bytes.NewBuffer(jsonData))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", API_KEY)

	res, _ := client.Do(req)

	defer res.Body.Close()

	// Decode json and store it in a map
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	// Print the id of the transcribed audio
	fmt.Println(result["id"])
}
