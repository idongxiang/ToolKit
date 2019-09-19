package main

import (
	"encoding/json"
	"fmt"
)

type SqlRequest struct {
	from           SelectFrom
	condition      SkyNetLog
	regexCondition RegexCondition
}

type QueryRequest struct {
	Query       string
	Description string
}

func queryRequest(body SqlRequest) QueryRequest {
	sql := sql(body.from, body.condition, body.regexCondition)
	request := QueryRequest{}
	request.Query = sql

	return request
}

func requestBody(body SqlRequest) string {
	request := queryRequest(body)
	bytes, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bytes)
}

func main() {
	log := SkyNetLog{
		appid:        "",
		logtime:      "",
		msg:          "",
		extrainfo:    "",
		module:       "com.xx.core",
		category:     "api",
		sub_category: "",
		priority:     "",
		ip:           "",
		filter1:      "",
		filter2:      "",
		contextid:    "",
		domainname:   "",
		dt:           "20190916",
		hour:         "",
	}

	reg := RegexCondition{
		regex:  "code\\\": \\\"(\\d)\\\"",
		target: "0",
	}

	from := SelectFrom{
		fields: []string{"count(*)"},
		table:  "skynet_all_101312",
	}

	request := SqlRequest{}
	request.from = from
	request.condition = log
	request.regexCondition = reg

	fmt.Println(requestBody(request))
}
