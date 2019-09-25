package main

import (
	"reflect"
	"strings"
)

type RegexCondition struct {
	Regex    string
	Target   string
	Operator string
}

type ExpressCondition struct {
	Express  string
	Target   string
	Operator string
}

type RangeCondition struct {
	Field    string
	Target   string
	Operator string
}

type Select struct {
	Fields []string
	Table  string
}

func sql(s Select, condition interface{}, regex RegexCondition, express ExpressCondition, rc []RangeCondition) string {
	sql := strings.Builder{}
	sql.WriteString(_select(s.Fields))
	sql.WriteString(_from(s.Table))
	sql.WriteString(_where(condition, regex, express, rc))
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

func _where(condition interface{}, regex RegexCondition, express ExpressCondition, rc []RangeCondition) string {
	sql := strings.Builder{}
	statement := _condition(condition)
	if len(statement) == 0 {
		return ""
	}
	sql.WriteString(statement)
	handleRange(rc, &sql)
	handleReg(regex, &sql)
	handleExpress(express, &sql)
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
			if skyNetIn(fn) {
				handleIn(fv, &statement)
			} else {
				handleEq(fv, &statement)
			}
			first = false
		} else {
			handleAnd(fn, &statement)
			if skyNetIn(fn) {
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
	b.WriteString("', 1) ")
	if len(reg.Operator) > 0 {
		b.WriteString(reg.Operator)
	} else {
		b.WriteString("=")
	}
	b.WriteString(" '")
	b.WriteString(reg.Target)
	b.WriteString("'")
}

func handleExpress(express ExpressCondition, b *strings.Builder) {
	if len(express.Express) == 0 {
		return
	}
	b.WriteString(" and ")
	b.WriteString(express.Express)
	if len(express.Operator) > 0 {
		b.WriteString(express.Operator)
	} else {
		b.WriteString(" =")
	}
	b.WriteString(" '")
	b.WriteString(express.Target)
	b.WriteString("'")
}

func handleRange(rc []RangeCondition, b *strings.Builder) {
	if len(rc) == 0 {
		return
	}
	for _, r := range rc {
		b.WriteString(" and ")
		b.WriteString(r.Field)
		b.WriteString(" ")
		b.WriteString(r.Operator)
		b.WriteString(" '")
		b.WriteString(r.Target)
		b.WriteString("'")
	}
}
