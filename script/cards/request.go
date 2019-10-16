package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func doPostQuery(requestBody string) (string, error) {
	fmt.Println(config.Url)
	req, e := http.NewRequest(http.MethodPost, config.Url, strings.NewReader(requestBody))
	if e != nil {
		return "", e
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Authorization", config.Authentication)
	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		return "", e
	}

	defer res.Body.Close()
	body, e := ioutil.ReadAll(res.Body)

	if e != nil {
		fmt.Println(e.Error())
		return "", e
	}

	return string(body), nil
}
