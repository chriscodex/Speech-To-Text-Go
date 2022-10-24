package main

import (
	"log"
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

	const AUDIO_URL = ""
	const TRANSCRIPT_URL = "https://api.assemblyai.com/v2/transcript"

}
