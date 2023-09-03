package common

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpPost(url string, data string) (string, error) {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	resp.Header.Set("charset", "utf-8")
	resp.Header.Set("Content-Type", "text/html")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "", errors.New("http send err")
	}

	fmt.Println(string(body))
	if resp.StatusCode == 200 {
		return string(body), nil
	} else {
		return "", errors.New("http send err")
	}

}

func HttpPostForm() {
	resp, err := http.PostForm("https://www.denlery.top/api/v1/login",
		url.Values{"username": {"auto"}, "password": {"auto123123"}})
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))

}

func HttpPostJson() {
	jsonStr := []byte(`{ "username": "auto", "password": "auto123123" }`)
	url := "https://www.denlery.top/api/v1/login"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	statuscode := resp.StatusCode
	hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(statuscode)
	fmt.Println(hea)

}
