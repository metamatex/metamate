package types

import "strings"

const (
	FORMAT_JSON = "json"
	FORMAT_YAML = "yaml"
	FORMAT_DEFAULT = FORMAT_YAML
)

type Config struct {
	Output string
	VerbosityLevel int
	OutputFormat string
	GlobalConfigPath string
	ProjectConfigPath string
	Addr string
	Token string
	NoColor bool
}

func (c Config) ReturnData() (b bool) {
	switch strings.ToLower(c.OutputFormat) {
	case FORMAT_YAML:
		b = true
	case FORMAT_JSON:
		b = true
	default:
		b = false
	}

	return
}