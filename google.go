package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const googleBaseUrl = "https://translate.googleapis.com/translate_a/single"

func google(c config, text string) (translation string, err error) {
	params := url.Values{}
	params.Add("client", "gtx")
	params.Add("dt", "t")
	params.Add("q", text)
	params.Add("sl", c.source)
	params.Add("tl", c.target)

	requestFullUrl := fmt.Sprintf("%s?%s", googleBaseUrl, params.Encode())
	resp, err := http.Get(requestFullUrl)
	if err != nil {
		return "", fmt.Errorf("cannot perform request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read response body: %w", err)
	}

	var result []any
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse json response. reading body: %v, error: %w", body, err)
	}

	a := result[0].([]any)
	var output strings.Builder
	for _, t := range a {
		t := t.([]any)
		str := t[0].(string)
		output.WriteString(str)
	}
	return output.String(), nil
}
