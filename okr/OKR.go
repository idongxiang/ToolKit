package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const CasePath string = "okr/case/"

func main() {
	files, e := ioutil.ReadDir(CasePath)
	if e != nil {
		fmt.Println(e)
		return
	}

	ch := make(chan string, len(files))
	for _, f := range files {
		go chDoCase(CasePath+f.Name(), ch)
	}

	for i := 0; i < len(files); i++ {
		fmt.Println(<-ch)
	}
}

func chDoCase(file string, ch chan string) {
	ch <- doCase(file)
}

func doCase(file string) string {
	b, e := ioutil.ReadFile(file)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	request := SqlRequest{}
	e = json.Unmarshal(b, &request)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	b, e = json.Marshal(request.Condition)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	log := SkyNetLog{}
	e = json.Unmarshal(b, &log)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	request.Condition = log
	return doSqlQuery(request)
}

func _main() {
	r := SqlRequest{}

	s := Select{
		Fields: []string{"count(*)"},
		Table:  "skynet_all_101312",
	}

	condition := SkyNetLog{
		Appid:        "",
		Logtime:      "",
		Msg:          "",
		Extrainfo:    "",
		Module:       "com.ly.train.third.service.kyfw12306.core",
		Category:     "ApiManager",
		Sub_category: "",
		Priority:     "",
		Ip:           "",
		Filter1:      "getStationList",
		Filter2:      "",
		Contextid:    "",
		Domainname:   "",
		Dt:           "20190919",
		Hour:         "",
	}

	regexCondition := RegexCondition{}
	regexCondition.Regex = ""
	regexCondition.Target = ""

	r.Select = s
	r.Condition = condition
	r.RegexCondition = regexCondition

	doSqlQuery(r)
}
