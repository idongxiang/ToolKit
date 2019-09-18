package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const QueryUrl string = "http://query"

func main() {
	fmt.Println("Hello, 世界")
	request := SkyNetLog{
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
		dt:           "",
		hour:         "",
	}

	fmt.Println(request.appid)

	sql := "{\"query\":\"SELECT COUNT(*) FROM skynet_all_101312 " +
		"WHERE \\\"dt\\\" IN ('20190911') " +
		"AND \\\"module\\\" = 'com.ly.train.third.service.kyfw12306.core' " +
		"AND \\\"category\\\" = 'ApiManager' " +
		"AND \\\"filter1\\\" = 'getStationList' " +
		"AND \\\"REGEXP_EXTRACT\\\"(msg,'code\\\": \\\\\\\"(\\\\d)\\\\\\\"', 1) = '0'\"}"

	req, e := http.NewRequest(http.MethodPost, QueryUrl, strings.NewReader(sql))
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Authentication", "6399921d6c002bd6a85c809ee74519ba")

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	defer res.Body.Close()

	body, e := ioutil.ReadAll(res.Body)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Println(string(body))
}
