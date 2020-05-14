package types

type SdkGenerator struct {
	Name         string                                     `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string                                     `json:"description,omitempty" yaml:"description,omitempty"`
	Init         func(*SdkGenerator, SdkConfig) (err error) `json:"-" yaml:"-"`
	Reset        func(SdkConfig) (err error)                `json:"-" yaml:"-"`
	Tasks        []RenderTask                               `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Dependencies []RenderTask                               `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
}
