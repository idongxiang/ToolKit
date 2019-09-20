package main

func main() {
	request := SqlRequest{}

	from := SelectFrom{
		Fields: []string{"count(*)"},
		Table:  "skynet_all_101312",
	}

	condition := SkyNetLog{
		appid:        "",
		logtime:      "",
		msg:          "",
		extrainfo:    "",
		module:       "com.ly.train.third.service.kyfw12306.core",
		category:     "ApiManager",
		sub_category: "",
		priority:     "",
		ip:           "",
		filter1:      "getStationList",
		filter2:      "",
		contextid:    "",
		domainname:   "",
		dt:           "20190919",
		hour:         "",
	}

	regexCondition := RegexCondition{}
	regexCondition.Regex = ""
	regexCondition.Target = ""

	request.From = from
	request.Condition = condition
	request.RegexCondition = regexCondition

	doSqlQuery(request)
}
