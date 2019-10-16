package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const DataFile string = ResourcePath + "data.csv"
const FailLog string = ResourcePath + "fail.log"

func main() {
	file, e := os.Open(DataFile)
	if e != nil {
		fmt.Println(e)
		return
	}
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		sendCard(record[0], record[1])
	}
}

type CardRequest struct {
	OpenId      string `json:"openId"`
	UnionId     string `json:"unionId"`
	ActivityId  string `json:"activityId"`
	SourceValue string `json:"sourceValue"`
	Source      string `json:"Source"`
	Operator    string `json:"operator"`
}

type CardResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const SUCCESS string = "0000"

func sendCard(unionId string, openId string) {
	r := CardRequest{
		OpenId:      openId,
		UnionId:     unionId,
		ActivityId:  "475",
		SourceValue: "脚本批量发券",
		Source:      "7",
		Operator:    "兼职-杨勇壮",
	}
	bytes, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		appendError(unionId, openId, err.Error())
		return
	}
	result, err := doPostQuery(string(bytes))
	if err != nil {
		fmt.Println(err)
		appendError(unionId, openId, err.Error())
		return
	}
	response := CardResponse{}
	err = json.Unmarshal([]byte(result), &response)
	if err != nil {
		fmt.Println(err)
		appendError(unionId, openId, err.Error())
		return
	}
	if !strings.EqualFold(SUCCESS, response.Code) {
		appendError(unionId, openId, response.Message)
		return
	}
	fmt.Println(unionId + "," + openId)
}

func appendError(unionId string, openId string, msg string) {
	f, err := os.OpenFile(FailLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(unionId + "," + openId + "," + msg + "\n")); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
