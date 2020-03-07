package types

type DataReport struct {
	MessageReport `yaml:",omitempty,inline" json:",omitempty"`
	Data          interface{} `yaml:",omitempty" json:"data,omitempty"`
}
