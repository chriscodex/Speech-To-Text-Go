package main

import (
	"fmt"

	"github.com/ChrisCodeX/Speech-To-Text-Go.git/upload"
)

func main() {
	uploadURL := upload.GetUploadURL()
	fmt.Println(uploadURL)
}
