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
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func SendMessage(url, text string) (int, error) {
	//var url = 'https://xxxxx.ngrok.io/google-home-notifier';
	//var urlFetchOption = {
	//	'method' : 'post',
	//	'contentType' : 'application/x-www-form-urlencoded',
	//	'payload' : { 'text' : text}
	//};

	//application/x-www-form-urlencoded
	values := url.Values{}
	values.Set("text", text)

	return post(url, strings.NewReader(values.Encode()))
}
