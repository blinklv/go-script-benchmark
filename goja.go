package benchmark

import (
	"context"
	"net/url"

	"github.com/dop251/goja"
)

// Goja javascript virtual machine.
type Goja struct {
	r     *goja.Runtime
	match func(interface{}) bool
	fib   func(int) int
}

// NewGoja creates a Goja instance.
func NewGoja(text string) (*Goja, error) {
	r := goja.New()
	r.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	if _, err := r.RunString(text); err != nil {
		return nil, err
	}

	var match func(interface{}) bool
	if err := r.ExportTo(r.Get("match"), &match); err != nil {
		return nil, err
	}

	var fib func(int) int
	if err := r.ExportTo(r.Get("fib"), &fib); err != nil {
		return nil, err
	}

	if err := r.Set("parse_query", gojaParseQuery); err != nil {
		return nil, err
	}

	return &Goja{
		r:     r,
		match: match,
		fib:   fib,
	}, nil
}

// Match checks whether the obj parameter satisfies some rules.
func (vm *Goja) Match(_ context.Context, obj interface{}) bool {
	return vm.match(obj)
}

// Fib computes fibonacci number(n) then returns it.
func (vm *Goja) Fib(_ context.Context, n int) int {
	return vm.fib(n)
}

func gojaParseQuery(fcall goja.FunctionCall, r *goja.Runtime) goja.Value {
	query := fcall.Argument(0).String()
	if query == "" {
		return goja.Undefined()
	}

	vs, err := url.ParseQuery(query)
	if err != nil {
		return goja.Undefined()
	}

	m := make(map[string]interface{})
	for k, v := range vs {
		m[k] = v
	}
	return r.ToValue(m)
}
