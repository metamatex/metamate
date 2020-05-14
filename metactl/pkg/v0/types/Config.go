package types

import "strings"

const (
	FormatJson    = "json"
	FormatYaml    = "yaml"
	FormatDefault = FormatYaml
)

type GlobalArgs struct {
	Output            string
	VerbosityLevel    int
	OutputFormat      string
	GlobalConfigPath  string
	ProjectConfigPath string
	Addr              string
	Token             string
	NoColor           bool
}

func (a GlobalArgs) ReturnData() (b bool) {
	switch strings.ToLower(a.OutputFormat) {
	case FormatYaml:
		b = true
	case FormatJson:
		b = true
	default:
		b = false
	}

	return
}
