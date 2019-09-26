package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func doCaseByConfigDays(path string) *map[string][]string {
	files, e := ioutil.ReadDir(path)
	if e != nil {
		fmt.Println(e)
		return nil
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
		return nil
	}

	results := make(map[string][]string)
	for _, c := range cases {
		requests := getCasesByConfigDays(path + c)
		for _, r := range requests {
			result := doSqlQuery(r)
			values := results[r.Condition.(SkyNetLog).Dt]
			values = append(values, result)
			results[r.Condition.(SkyNetLog).Dt] = values
		}
	}

	for k, v := range results {
		fmt.Println("k = " + k + ", v =" + strconv.Itoa(len(v)))
	}

	return &results
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

func getCasesByConfigDays(file string) []SqlRequest {
	request := *getCase(file)
	log := request.Condition.(SkyNetLog)

	dt, e := time.ParseInLocation("20060102", log.Dt, time.Local)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	var requests []SqlRequest
	for i := 0; i < config.Days; i++ {
		if strings.EqualFold(config.Split, "hour") {
			splitDayByHours(i, log, dt, request, &requests)
		} else {
			requests = append(requests, request)
		}
	}

	return requests
}

func splitDayByHours(day int, log SkyNetLog, dt time.Time, request SqlRequest, requests *[]SqlRequest) {
	for j := 0; j < 24; j++ {
		t := &log
		dt := dt.AddDate(0, 0, -day)
		t.Dt = dt.Format("20060102")
		r := &request
		r.Condition = *t

		var lowerHour time.Duration
		var e error
		if j != 0 {
			lowerHour, e = time.ParseDuration(strconv.Itoa(j) + "h")
			if e != nil {
				fmt.Println(e)
				return
			}
		}

		upperHour, e := time.ParseDuration(strconv.Itoa(j+1) + "h")
		if e != nil {
			fmt.Println(e)
			return
		}

		var lowerLogTime string
		if j == 0 {
			lowerLogTime = dt.Format("2006-01-02 15:04:05")
		} else {
			lowerLogTime = dt.Add(lowerHour).Format("2006-01-02 15:04:05")
		}
		lower := RangeCondition{}
		lower.Field = LogTime
		lower.Operator = ">="
		lower.Target = lowerLogTime

		upperLogTime := dt.Add(upperHour).Format("2006-01-02 15:04:05")
		upper := RangeCondition{}
		upper.Field = LogTime
		upper.Operator = "<"
		upper.Target = upperLogTime

		rcs := []RangeCondition{lower, upper}

		r.RangeCondition = rcs
		*requests = append(*requests, *r)
	}
}
