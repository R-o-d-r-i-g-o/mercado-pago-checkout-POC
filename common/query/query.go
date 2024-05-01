package query

import (
	"fmt"
	"reflect"
	"strings"
)

type Query string

func New(query string) Query {
	return Query(query)
}

func (q Query) Replace(args ...any) Query {
	needReplaceAll := len(args) == 1

	for _, arg := range args {
		kt := reflect.TypeOf(arg)

		if kt.Kind() == reflect.String {
			arg = "'" + strings.Replace(arg.(string), "'", "\\'", -1) + "'"
		}

		if needReplaceAll {
			q = Query(strings.ReplaceAll(q.String(), "$query", String(arg)))
		} else {
			q = Query(strings.Replace(q.String(), "$query", String(arg), 1))
		}

	}

	return q
}

func String(arg any) string {
	return fmt.Sprintf("%v", arg)
}

func (q Query) String() string {
	return string(q)
}
