package main

type SkyNetLog struct {
	appid        string
	logtime      string
	msg          string
	extrainfo    string
	module       string
	category     string
	sub_category string
	priority     string
	ip           string
	filter1      string
	filter2      string
	contextid    string
	domainname   string
	dt           string
	hour         string
}

const LogDt string = "dt"