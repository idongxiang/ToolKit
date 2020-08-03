package main

import (
	"fmt"
	"reflect"
	"strings"
)

const LogDt string = "dt"

func _main() {
	log := SkyNetLog{
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
		dt:           "20190916",
		hour:         "",
	}

	reg := RegexCondition{
		regex:  "code\\\": \\\"(\\d)\\\"",
		target: "1",
	}

	reg = RegexCondition{
		regex:  "",
		target: "1",
	}

	from := SelectFrom{
		fields: []string{"logtime", "msg"},
		table:  "skynet_all_101312",
	}

	sql := sql(from, log, reg)
	fmt.Println(sql)
}

type RegexCondition struct {
	regex  string
	target string
}

type SelectFrom struct {
	fields []string
	table  string
}

func sql(from SelectFrom, condition interface{}, regex RegexCondition) string {
	sql := strings.Builder{}
	sql.WriteString(_select(from.fields))
	sql.WriteString(_from(from.table))
	sql.WriteString(_where(condition, regex))
	return sql.String()
}

func _select(fields []string) string {
	statement := strings.Builder{}
	statement.WriteString("select")
	for _, field := range fields {
		statement.WriteString(" ")
		statement.WriteString(field)
		statement.WriteString(",")
	}
	return strings.TrimRight(statement.String(), ",")
}

func _from(table string) string {
	return " from " + table
}

func _where(condition interface{}, regex RegexCondition) string {
	sql := strings.Builder{}
	statement := _condition(condition)
	if len(statement) == 0 {
		return ""
	}
	sql.WriteString(statement)
	handleReg(regex, &sql)
	return sql.String()
}

func _condition(condition interface{}) string {
	t := reflect.TypeOf(condition)
	v := reflect.ValueOf(condition)
	statement := strings.Builder{}
	first := true
	for i := 0; i < t.NumField(); i++ {
		fn := t.Field(i).Name
		fv := v.Field(i).String()
		if len(fv) == 0 {
			continue
		}
		if first {
			handleWhere(fn, &statement)
			if fn == LogDt {
				handleIn(fv, &statement)
			} else {
				handleEq(fv, &statement)
			}
			first = false
		} else {
			handleAnd(fn, &statement)
			if fn == LogDt {
				handleIn(fv, &statement)
			} else {
				handleEq(fv, &statement)
			}
		}
	}

	return statement.String()
}

func handleWhere(f string, b *strings.Builder) {
	b.WriteString(" where ")
	b.WriteString(f)
}

func handleAnd(f string, b *strings.Builder) {
	b.WriteString(" and ")
	b.WriteString(f)
}

func handleIn(v string, b *strings.Builder) {
	b.WriteString(" in ('")
	b.WriteString(v)
	b.WriteString("')")
}

func handleEq(v string, b *strings.Builder) {
	b.WriteString(" = '")
	b.WriteString(v)
	b.WriteString("'")
}

func handleReg(reg RegexCondition, b *strings.Builder) {
	if len(reg.regex) == 0 {
		return
	}
	b.WriteString(" and REGEXP_EXTRACT(msg, '")
	b.WriteString(reg.regex)
	b.WriteString("', 1) = '")
	b.WriteString(reg.target)
	b.WriteString("'")
}
