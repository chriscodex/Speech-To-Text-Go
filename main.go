package main

import (
	"fmt"
	"time"

	"github.com/ChrisCodeX/Speech-To-Text-Go.git/poll"
	"github.com/ChrisCodeX/Speech-To-Text-Go.git/transcript"
	"github.com/ChrisCodeX/Speech-To-Text-Go.git/upload"
)

const timeToWait = 10
const audioName = "record1.m4a"

func main() {
	// Get Upload URL of audio
	uploadURL := upload.GetUploadURL(audioName)
	fmt.Println(uploadURL)

	// Get the Id of transcription
	transcriptId := transcript.GetIdTranscription(uploadURL)
	fmt.Println(transcriptId)

	time.Sleep(timeToWait * time.Second)

	// Get text transcripted
	textTranscripted := poll.GetTextTranscripted(transcriptId)
	fmt.Println(textTranscripted)
}
