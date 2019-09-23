package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func doLatestWeekCase(path string) {
	files, e := ioutil.ReadDir(path)
	if e != nil {
		fmt.Println(e)
		return
	}

	var cases []string
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		cases = append(cases, f.Name())
	}

	if len(cases) == 0 {
		fmt.Println("no case files")
		return
	}

	for _, c := range cases {
		requests := getLatestWeekCase(path + c)
		for _, r := range requests {
			doSqlQuery(r)
		}
	}
}

func doRealChCase(path string) {
	files, e := ioutil.ReadDir(path)
	if e != nil {
		fmt.Println(e)
		return
	}

	var cases []string
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		cases = append(cases, f.Name())
	}

	if len(cases) == 0 {
		fmt.Println("no case files")
		return
	}

	ch := make(chan string, len(cases))
	for _, c := range cases {
		chDoCase(path+c, ch)
	}

	for i := 0; i < len(files); i++ {
		fmt.Println(<-ch)
	}
}

func chDoCase(file string, ch chan string) {
	ch <- doSqlQuery(*getCase(file))
}

func getCase(file string) *SqlRequest {
	b, e := ioutil.ReadFile(file)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	request := SqlRequest{}
	e = json.Unmarshal(b, &request)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	b, e = json.Marshal(request.Condition)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	log := SkyNetLog{}
	e = json.Unmarshal(b, &log)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	request.Condition = log

	return &request
}

func getLatestWeekCase(file string) []SqlRequest {
	request := *getCase(file)
	log := request.Condition.(SkyNetLog)

	dt, e := time.ParseInLocation("20060102", log.Dt, time.Local)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	var requests []SqlRequest
	for i := 0; i < 7; i++ {
		t := &log
		t.Dt = dt.AddDate(0, 0, -i).Format("20060102")
		r := &request
		r.Condition = *t
		requests = append(requests, *r)
	}

	return requests
}
