// Package http is just sample
package http

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	lg "github.com/hiromaily/golibs/log"
	//u "github.com/hiromaily/golibs/utils"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

//https://gowalker.org/github.com/parnurzeal/gorequest
//https://godoc.org/github.com/parnurzeal/gorequest

// create signature for authentication
func createSignature(data []byte, secret string) string {
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write(data)
	//signature := url.QueryEscape(base64.StdEncoding.EncodeToString(hash.Sum(nil)))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return signature
}

// debug for http request
func debugHTTPRequest(data []byte, err error) {
	if err == nil {
		//log.Debug(fmt.Sprintf("dump of http request\n%s", data))
		lg.Debugf("dump of http request : %s", data)
	} else {
		//log.Fatal(fmt.Sprintf("%v", err))
		lg.Fatal(err)
	}
}

// set HTTP headers with content length
func setHTTPHeadersWithContentLength(req *http.Request, contentLength string) {
	//TODO:What is different between Header.Set and Header.Add？ Try to chenge Set to Add.
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Content-Length", contentLength)
	req.Header.Add("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Add("Connection", "close")
	//req.Header.Del("Accept-Encoding")
}

// set HTTP headers
func setHTTPHeaders(req *http.Request) {
	req.Header.Add("Cookie", "xxxxx_csrf_token=FoZUJzY0xWTmZKaW9oZ1k")
	req.Header.Add("Origin", "https://xxxx.com")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "ja,en-US;q=0.8,en;q=0.6,de;q=0.4,nl;q=0.2")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Referer", "https://xxxx.com/entry")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
}

func createClient() *http.Client {
	//gzip
	tr := &http.Transport{
		DisableCompression: false,
	}
	return &http.Client{Transport: tr}
}

func createSendData() url.Values {
	data := make(url.Values)

	data.Add("param1", "1")
	data.Add("param2", "string something")

	return data
}

// HandleResponse is to handle response
func HandleResponse(resp *http.Response) string {
	//resp.StatusCode
	//resp.Header

	//OK
	//lg.Debugf("resp.StatusCode: %d", resp.StatusCode)

	//HTTP Response headers
	//contentType := resp.Header.Get("Content-Type")
	//lg.Debugf("contentType: %s", contentType)

	defer resp.Body.Close()

	//handling response
	resbody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return ""
	}
	//lg.Debug("response body: %s", string(resbody))

	return string(resbody)
}

// PostRequest is to request by POST using http.NewRequest
func PostRequest(reqURL string, bytesMessage []byte) (int, string, error) {

	client := createClient()

	req, err := http.NewRequest(
		"POST",
		reqURL,
		bytes.NewBuffer(bytesMessage),
	)
	if err != nil {
		return 0, "", err
	}

	//Set Http Headers
	setHTTPHeadersWithContentLength(req, strconv.Itoa(len(bytesMessage)))

	//debug http request
	//byte, err := httputil.DumpRequestOut(req, true)
	debugHTTPRequest(httputil.DumpRequestOut(req, true))

	req.Close = true //これを追加したらリクエスト可能な数が増えた。

	//HTTP request
	resp, err := client.Do(req)

	if err != nil {
		//err e.g.
		//error: xxx socket: too many open files
		//-> $ ulimit -n 2048
		return resp.StatusCode, "", err
	}

	//handle response
	return resp.StatusCode, HandleResponse(resp), nil
}

// GetRequestWithData is to request by GET with data using http.NewRequest
func GetRequestWithData(reqURL string) (int, string, error) {

	client := &http.Client{}

	data := createSendData()

	req, err := http.NewRequest(
		"GET",
		reqURL,
		bytes.NewBuffer([]byte(data.Encode())),
	)

	if err != nil {
		return 500, "", err
	}

	//Set Http Headers
	setHTTPHeaders(req)

	req.Close = true //これを追加したらリクエスト可能な数が増えた。

	//HTTP request
	resp, err := client.Do(req)

	if err != nil {
		//return resp.StatusCode, "", err
		return 500, "", err
	}

	//handle response
	return resp.StatusCode, HandleResponse(resp), nil
}

// GetRequestSimple is to send request by Get (just sample code)
func GetRequestSimple(url string) (int, string, error) {

	//HTTP Request
	response, err := http.Get(url)

	if err != nil {
		return 500, "", err
	}

	//fmt.Println("status:", response.Status)
	//200 OK

	preStatus := strings.Split(response.Status, " ")
	status, _ := strconv.Atoi(preStatus[0])

	//read response body(byte)
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return 500, "", err
	}

	//show body
	return status, string(body), nil
}

//-----------------------------------------------------------------------------
// gorequest
//-----------------------------------------------------------------------------

// Get is get method using gorequest package
func Get(url string) {

	// gorequest
	request := gorequest.New()
	resp, body, errs := request.Get(url).End()
	fmt.Println(resp, body, errs)
	//
	//resp, body, errs := gorequest.New().Get(url).End()

}
