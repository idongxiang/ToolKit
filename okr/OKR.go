package main

const CasePath string = "okr/case/"

func main() {
	//doRealChCase(CasePath)
	doLatestWeekCase(config.Case)
}

func _test() {
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
	regexCondition.Operator = ""

	r.Select = s
	r.Condition = condition
	r.RegexCondition = regexCondition

	doSqlQuery(r)
}
