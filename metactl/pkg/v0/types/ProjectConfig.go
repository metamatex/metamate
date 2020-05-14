package types

type ProjectConfig struct {
	V0 V0Project `json:"v0,omitempty" yaml:"v0,omitempty"`
}

type V0Project struct {
	Gen Gen `json:"gen,omitempty" yaml:"gen,omitempty"`
}

type Gen struct {
	Tasks []RenderTask `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Sdks  []SdkConfig  `json:"sdks,omitempty" yaml:"sdks,omitempty"`
}

type SdkConfig struct {
	Name      string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Args      map[string]interface{} `json:"args,omitempty" yaml:"args,omitempty"`
	Endpoints []string               `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
}
