package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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

	var parts []string
	for i, items := range result[0].([]any) {
		if items == nil {
			continue
		}
		fmt.Println(i, items)
		arr, ok := items.([]any)
		if !ok || len(arr) == 0 {
			continue
		}
		fmt.Println(i, arr)
		text, ok := arr[0].(string)
		if !ok {
			continue
		}
		fmt.Print(i, text)

		parts = append(parts, text)
		fmt.Println(parts)
	}
	return strings.Join(parts, " "), nil
}
