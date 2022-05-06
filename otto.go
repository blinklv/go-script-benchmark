package benchmark

import (
	"context"
	"fmt"
	"net/url"

	"github.com/robertkrimen/otto"
)

// Otto javascript virtual machine.
type Otto struct {
	o *otto.Otto
}

func NewOtto(text string) (*Otto, error) {
	o := otto.New()
	if err := o.Set("parse_query", ottoParseQuery); err != nil {
		return nil, fmt.Errorf("set parse_query failed: %v", err)
	}
	_, err := o.Run(text)
	if err != nil {
		return nil, fmt.Errorf("run script failed: %v", err)
	}
	return &Otto{o: o}, nil
}

func (vm *Otto) Match(_ context.Context, obj interface{}) bool {
	arg, err := vm.o.ToValue(obj)
	if err != nil {
		panic("ToValue(obj) failed")
	}
	res, err := vm.o.Call("match", otto.NullValue(), arg)
	if err != nil {
		panic("otto match panic")
	}
	b, err := res.ToBoolean()
	if err != nil {
		panic("result to bool failed")
	}
	return b
}

func (vm *Otto) Fib(_ context.Context, n int) int {
	arg, err := vm.o.ToValue(n)
	if err != nil {
		panic("ToValue(obj) failed")
	}
	res, err := vm.o.Call("fib", otto.NullValue(), arg)
	if err != nil {
		panic("otto fib panic")
	}
	fn, err := res.ToInteger()
	if err != nil {
		panic("result to int failed")
	}
	return int(fn)
}

func ottoParseQuery(fcall otto.FunctionCall) otto.Value {
	query := fcall.Argument(0).String()
	if query == "" {
		return otto.UndefinedValue()
	}

	vs, err := url.ParseQuery(query)
	if err != nil {
		return otto.UndefinedValue()
	}

	m := make(map[string]interface{})
	for k, v := range vs {
		m[k] = v
	}

	value, err := fcall.Otto.ToValue(m)
	if err != nil {
		return otto.UndefinedValue()
	}

	return value
}
