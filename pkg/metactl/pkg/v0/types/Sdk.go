package types

type Sdk struct {
	Name         string                                         `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string                                         `json:"description,omitempty" yaml:"description,omitempty"`
	Init         func(*Sdk, map[string]interface{}) (err error) `json:"-" yaml:"-"`
	Reset        func() (err error)                             `json:"-" yaml:"-"`
	Tasks        []RenderTask                                   `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Dependencies []RenderTask                                   `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
}
