package testing

import (
	"fmt"
	"sync"
	"testing"
)

type TB interface {
	Errorf(format string, args ...interface{})
	Run(name string, f func(t TB))
	Parallel()
}

type StdT struct {
	t *testing.T
}

func (t *StdT) Parallel() {
	t.t.Parallel()
}

func (t *StdT) Errorf(format string, args ...interface{}) {
	t.t.Errorf(format, args...)
}

func (t *StdT) Run(name string, f func(t TB)) {
	t.t.Run(name, func(t *testing.T) {
		f(&StdT{t: t})
	})
}

type T struct {
	Name string `yaml:"name,omitempty"`
	Errors []string `yaml:"errors,omitempty"`
	Children []*T `yaml:"children,omitempty"`
	F func(t TB) `yaml:"-"`
	Wg sync.WaitGroup `yaml:"-"`
}

func NewT(name string, f func(t TB)) (t *T) {
	return &T{Name: name, F: f}
}

func (t *T) Parallel() {
}

func (t *T) Errorf(format string, args ...interface{}) {
	t.Errors = append(t.Errors, fmt.Sprintf(format, args...))
}

func (t *T) Run(name string, f func(t TB)) {
	t0 := NewT(t.Name + "/" + name, f)
	t.Children = append(t.Children, t0)
	t.Wg.Add(1)
	go func() {
		t0.F(t0)
		t0.Wg.Wait()
		t.Wg.Done()
	}()
}
