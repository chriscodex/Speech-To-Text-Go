package upload

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const UPLOAD_URL = "https://api.assemblyai.com/v2/upload"

func GetUploadURL(audioName string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Import API KEY from .env file
	API_KEY := os.Getenv("API_KEY")

	// Load file
	nameAudio := "upload/audio/" + audioName
	data, err := ioutil.ReadFile(nameAudio)

	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP client and set header
	client := http.Client{}

	req, _ := http.NewRequest("POST", UPLOAD_URL, bytes.NewBuffer(data))

	req.Header.Set("authorization", API_KEY)

	// Request
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	// decode json and store it in a map
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	// Return the upload_url
	return result["upload_url"].(string)
}
