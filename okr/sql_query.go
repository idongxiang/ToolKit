package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type QueryRequest struct {
	Query       string `json:"query"`
	Description string `json:"description"`
}

type SqlRequest struct {
	Select         Select
	Condition      interface{}
	RegexCondition RegexCondition
}

func doSqlQuery(request SqlRequest) string {
	body := getSqlBody(request)
	fmt.Println(body)
	res := doPostQuery(body)
	fmt.Println(res)
	return res
}

func getSqlBody(body SqlRequest) string {
	sql := sql(body.Select, body.Condition, body.RegexCondition)
	request := QueryRequest{}
	request.Query = sql

	bytes, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bytes)
}

func doPostQuery(requestBody string) string {
	fmt.Println(config.Sql.Url)
	req, e := http.NewRequest(http.MethodPost, config.Sql.Url, strings.NewReader(requestBody))
	if e != nil {
		fmt.Println(e.Error())
		return ""
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Authentication", config.Sql.Authentication)

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
