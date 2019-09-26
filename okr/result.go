package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type QueryResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  Result `json:"result"`
	Data    string `json:"data"`
}

type Result struct {
	Headers []string        `json:"headers"`
	Results [][]interface{} `json:"results"`
	Error   string          `json:"error"`
}

func mergeCountResults(results map[string][]string) {
	for k, v := range results {
		count := 0
		for _, r := range v {
			qr := QueryResult{}
			e := json.Unmarshal([]byte(r), &qr)
			if e != nil {
				fmt.Println("服务器错误:" + e.Error() + ",r=" + r)
				return
			}

			if qr.Code != 200 {
				fmt.Println("响应错误:" + ",r=" + r)
				continue
			}

			if len(qr.Result.Error) != 0 {
				fmt.Println("响应错误:2,r=" + r)
				continue
			}

			if len(qr.Result.Results) > 1 {
				fmt.Println("响应结果错误1,r=" + r)
				return
			}
			if len(qr.Result.Results[0]) > 1 {
				fmt.Println("响应结果错误2,r=" + r)
				return
			}

			_, ok := qr.Result.Results[0][0].(float64)
			if ok {
				count += int(qr.Result.Results[0][0].(float64))
			} else {
				i, e := strconv.Atoi(qr.Result.Results[0][0].(string))
				if e != nil {
					fmt.Println("响应结果错误3,r=" + e.Error() + ",r=" + r)
					return
				}
				count += i
			}
		}
		fmt.Println("日期 = " + k + ", 数量 =" + strconv.Itoa(count))
	}
}
