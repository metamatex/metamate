package _go

import (
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
)

const (
	TaskUtils    = "TaskUtils"
	TaskUtilsPtr = "TaskUtilsPtr"
)

func init() {
	tasks[TaskUtils] = types.RenderTask{
		TemplateData: &goUtilsTpl,
		Out:          ptr.String("utils/utils_.go"),
	}

	tasks[TaskUtilsPtr] = types.RenderTask{
		TemplateData: &goUtilsPtrTpl,
		Out:          ptr.String("utils/ptr/ptr_.go"),
	}
}

var goUtilsTpl = `package utils
{{ $package := index .Data "package" }}
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
}`

var goUtilsPtrTpl = `package ptr

func String(s string) (*string) {
	return &s
}

func Uint32(i uint32) (*uint32) {
	return &i
}

func Int32(i int32) (*int32) {
	return &i
}

func Bool(b bool) (*bool) {
	return &b
}

func Float64(f float64) (*float64) {
	return &f
}`
