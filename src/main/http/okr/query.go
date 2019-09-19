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

func __main() {
	log := SkyNetLog{
		appid:        "skynet_all_101312",
		logtime:      "",
		msg:          "",
		extrainfo:    "",
		module:       "",
		category:     "",
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
		target: "1",
	}

	from := SelectFrom{
		fields: []string{"logtime", "msg"},
		table:  "skynet_all_101312",
	}

	request := SqlRequest{}
	request.from = from
	request.condition = log
	request.regexCondition = reg

	fmt.Println(requestBody(request))
}
