package common

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println(err)
		return "400", errors.New("http error")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if resp.StatusCode == 200 {
		return string(body), nil
	}
	return "400", err
}
