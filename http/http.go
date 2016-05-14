package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Http Request
func RequestHttp(url string) (int, string, error) {

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
	return status, string(body)
}
