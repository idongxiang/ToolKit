package main

import (
	"reflect"
	"strings"
)

type RegexCondition struct {
	Regex  string
	Target string
}

type Select struct {
	Fields []string
	Table  string
}

func sql(s Select, condition interface{}, regex RegexCondition) string {
	sql := strings.Builder{}
	sql.WriteString(_select(s.Fields))
	sql.WriteString(_from(s.Table))
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
			if strings.EqualFold(fn, LogDt) {
				handleIn(fv, &statement)
			} else {
				handleEq(fv, &statement)
			}
			first = false
		} else {
			handleAnd(fn, &statement)
			if strings.EqualFold(fn, LogDt) {
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
	b.WriteString(strings.ToLower(f))
}

func handleAnd(f string, b *strings.Builder) {
	b.WriteString(" and ")
	b.WriteString(strings.ToLower(f))
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
	if len(reg.Regex) == 0 {
		return
	}
	b.WriteString(" and REGEXP_EXTRACT(msg, '")
	b.WriteString(reg.Regex)
	b.WriteString("', 1) = '")
	b.WriteString(reg.Target)
	b.WriteString("'")
}
