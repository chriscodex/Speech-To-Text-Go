package main

import (
	"fmt"

	"github.com/ChrisCodeX/Speech-To-Text-Go.git/transcript"
	"github.com/ChrisCodeX/Speech-To-Text-Go.git/upload"
)

func main() {
	// Get Upload URL of audio
	uploadURL := upload.GetUploadURL()
	fmt.Println(uploadURL)

	transcriptId := transcript.GetIdTranscription(uploadURL)
	fmt.Println(transcriptId)
}
