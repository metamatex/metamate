// generated by metactl sdk gen 
package mql

import (
	"gopkg.in/yaml.v2"
	"regexp"
)

func Sprint(i interface{}) (string) {
	b, err := yaml.Marshal(i)
	if err != nil {
		return ""
	}

	re := regexp.MustCompile("(?m)[\r\n]+^.*xxx_unrecognized.*$")
	res := re.ReplaceAll(b, []byte{})

	re = regexp.MustCompile("(?m)[\r\n]+^.*: null.*$")
	res = re.ReplaceAll(res, []byte{})

	re = regexp.MustCompile("(?m)[\r\n]+^.*: \\[\\].*$")
	res = re.ReplaceAll(res, []byte{})

	return string(res)
}

func Print(i interface{}) {
	println(Sprint(i))
}