package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type lingvaResponse struct {
	Translation string  `json:"translation"`
	Error       *string `json:"error"`
}

func lingva(c config, text string) (translation string, err error) {

	// https://github.com/thedaviddelta/lingva-translate#rest-api-v1
	requestUrl, _ := url.JoinPath(c.lingvaDomain, "api", "v1", c.source, c.target, url.PathEscape(text))
	requestUrl = "https://" + requestUrl
	resp, err := http.Get(requestUrl)
	if err != nil {
		return "", fmt.Errorf("cannot perform request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read response body: %w", err)
	}

	var result lingvaResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse json response. reading body: %v, error: %w", body, err)
	}

	if result.Error != nil {
		return "", fmt.Errorf("api error: %s", *result.Error)
	}
	return result.Translation, nil
}
