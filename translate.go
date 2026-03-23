package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type TranslateRequest struct {
	Q      string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
}
type TranslateResponse struct {
	TranslatedText string `json:"translatedText"`
}

func translate(inputText string, targetLanguage string) (string, error) {
	url := "http://localhost:5000/translate"
	jsonPayload, err := json.Marshal(TranslateRequest{
		Q:      inputText,
		Source: "auto",
		Target: targetLanguage,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("translate API error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var output TranslateResponse
	jsonErr := json.Unmarshal(body, &output)
	if jsonErr != nil {
		return "", err
	}
	return output.TranslatedText, nil
}
