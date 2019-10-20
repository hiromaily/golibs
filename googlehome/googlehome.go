package googlehome

import (
	"net/http"
	"net/url"
	"strings"
)

func post(reqURL string, body *strings.Reader) (int, error) {

	req, err := http.NewRequest(
		"POST",
		reqURL,
		body,
	)
	if err != nil {
		return 0, err
	}

	req.Close = true

	// header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

// SendMessage is to post message to google home server
func SendMessage(reqURL, text string) (int, error) {
	//application/x-www-form-urlencoded
	values := url.Values{}
	values.Set("text", text)

	return post(reqURL, strings.NewReader(values.Encode()))
}
