package utils

import (
	"github.com/jinzhu/inflection"
	"strings"
)

func Title(ss []string) (ss0 []string) {
	for _, s := range ss {
		ss0 = append(ss0, strings.Title(s))
	}

	return
}

func Concat(ss []string) (s string) {
	for _, v := range ss {
		s += v
	}

	return
}

func Plural(s string) (string) {
	return inflection.Plural(s)
}
