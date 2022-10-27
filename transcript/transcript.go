package transcript

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"

func GetIdTranscription(audioURL string) interface{} {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Import API KEY from .env file
	API_KEY := os.Getenv("API_KEY")

	// Prepare json data
	values := map[string]string{"audio_url": audioURL}
	jsonData, _ := json.Marshal(values)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", TRANSCRIPT_URL, bytes.NewBuffer(jsonData))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", API_KEY)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	// Decode json and store it in a map
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	return result["id"]
}
