package main

type SkyNetLog struct {
	Appid        string `json:"appid"`
	Logtime      string `json:"logtime"`
	Msg          string `json:"msg"`
	Extrainfo    string `json:"extrainfo"`
	Module       string `json:"module"`
	Category     string `json:"category"`
	Sub_category string `json:"sub_category"`
	Priority     string `json:"priority"`
	Ip           string `json:"ip"`
	Filter1      string `json:"filter1"`
	Filter2      string `json:"filter2"`
	Contextid    string `json:"contextid"`
	Domainname   string `json:"domainname"`
	Dt           string `json:"dt"`
	Hour         string `json:"hour"`
}

const LogDt string = "dt"
