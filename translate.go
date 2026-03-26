package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func translate(inputText string, targetLanguage string) (string, error) {
	endpoint := fmt.Sprintf(
		"https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=%s&dt=t&q=%s",
		targetLanguage,
		url.QueryEscape(inputText),
	)

	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result []interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if len(result) == 0 {
		return "", fmt.Errorf("empty translation result")
	}

	translated := result[0].([]interface{})[0].([]interface{})[0].(string)
	return translated, nil
}
