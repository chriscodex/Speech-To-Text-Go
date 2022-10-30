# Speech to Text with Go ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ChrisCodeX/Speech-To-Text-Go)

This repository contains a speech-to-text audio converter in English, using the AssemblyAI API.

---

## **Table of Contents** 📖  
1. [Pre-Requirements](#pre-requirements-)
2. [Installation](#installation-)
3. [HTTP Endpoints](#http-endpoints-desktop_computer)
4. [WebSocket](#websocket-)

---
## **Pre-Requirements** 📋  
To have access to the API key that AssemblyAI provides for free, it is only necessary to register and the key will be generated automatically. You can do it from this link.

---

## **Installation** 🔧 
- The project uses external dependencies, so once the project is cloned you must run the command:
```
go mod tidy
```
- Create the .env file in the "/" path of the project and include the API key provided by AssemblyAI as follows:
```
API_KEY=Your_api_key
```
---  

## **Conversion** :desktop_computer: 
To convert your audio to text, add your audio to the "/upload/audio" folder, then in main.go modify these 2 variables with the data from your uploaded file:
```
const timeToWait = 10
const audioName = "record1.m4a"
```
Note: 
- timeToWait is the waiting time for the API to do the conversion, usually it is 30% of the total duration of the audio, so modify it according to the duration of your audio.
- audioName is the name of your audio, you should consider putting the format of your audio.
Supported formats: .3ga, .8svx, .aac, .ac3, .aif, .aiff, .alac, .amr, .ape, .au, .dss, .flac, .flv, .m4a, .m4b, .m4p, .m4r, .mp3, .mpga, .ogg, .oga, .mogg, .opus, .qcp, .tta, .voc, .wav, .wma, .wv