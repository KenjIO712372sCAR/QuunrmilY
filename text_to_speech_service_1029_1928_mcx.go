// 代码生成时间: 2025-10-29 19:28:54
package main

import (
    "encoding/json"
    "github.com/revel/revel"
    "google.golang.org/api/texttospeech/v1"
    "google.golang.org/api/transport"
    "context"
    "io/ioutil"
    "net/http"
    "os"
)

// TextToSpeechService is the controller for the text-to-speech functionality.
type TextToSpeechService struct {
    revel.Controller
}

// NewTextToSpeechService creates a new instance of TextToSpeechService.
func NewTextToSpeechService() *TextToSpeechService {
    return &TextToSpeechService{}
}

// Index is the action that handles the GET requests to the service.
func (c TextToSpeechService) Index() revel.Result {
    return c.Render()
}

// Synthesize is the action that handles the POST request to synthesize text into speech.
func (c TextToSpeechService) Synthesize() revel.Result {
    // Read the request body to get the text to synthesize.
    requestBody, err := ioutil.ReadAll(c.Params.Request.Body)
    if err != nil {
        c.Params.Error = "Error reading request body."
        return c.RenderError(err)
    }
    defer c.Params.Request.Body.Close()

    // Unmarshal the JSON request body.
    var合成请求体请求体请求请求请求请求请求体请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求-request请求请求请求请求请求请求-request请求request请求请求request-requestrequest-requestrequest-requestrequest-requestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequestrequest-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request-request请求请求请求请求请求请求请求请求请求请求体请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求请求	request := struct {
        Text string `json:"text"`
    }{}
    err = json.Unmarshal(requestBody, &request)
    if err != nil {
        c.Params.Error = "Error unmarshalling request body."
        return c.RenderError(err)
    }

    // Check if the text is empty.
    if request.Text == "" {
        c.Params.Error = "Text cannot be empty."
        return c.RenderError(revel.NewError("Text cannot be empty."))
    }

    // Set up the text-to-speech client.
    ctx := context.Background()
    client, err := texttospeech.NewService(ctx)
    if err != nil {
        c.Params.Error = "Error creating text-to-speech client."
        return c.RenderError(err)
    }

    // Create the synthesis input.
    input := texttospeech.SynthesisInput{
        Text: request.Text,
    }

    // Select the language and SSML voice gender (optional).
    voice := texttospeech.VoiceSelectionParams{
        LanguageCode: "en-US\