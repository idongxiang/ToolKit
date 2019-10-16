package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func doPostQuery(requestBody string) string {
	fmt.Println(config.Url)
	req, e := http.NewRequest(http.MethodPost, config.Url, strings.NewReader(requestBody))
	if e != nil {
		fmt.Println(e.Error())
		return ""
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Authentication", config.Authentication)

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		fmt.Println(e.Error())
		return ""
	}

	defer res.Body.Close()
	body, e := ioutil.ReadAll(res.Body)

	if e != nil {
		fmt.Println(e.Error())
		return ""
	}

	return string(body)
}

