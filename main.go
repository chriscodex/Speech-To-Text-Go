package main

import (
	"fmt"
	"time"

	"github.com/ChrisCodeX/Speech-To-Text-Go.git/poll"
	"github.com/ChrisCodeX/Speech-To-Text-Go.git/transcript"
	"github.com/ChrisCodeX/Speech-To-Text-Go.git/upload"
)

func main() {
	uploadURL := upload.GetUploadURL()
	fmt.Println(uploadURL)

	transcriptId := transcript.GetIdTranscription(uploadURL)
	fmt.Println(transcriptId)

	time.Sleep(10 * time.Second)

	textTranscripted := poll.GetTextTranscripted(transcriptId)
	fmt.Println(textTranscripted)
}
