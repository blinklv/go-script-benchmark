package benchmark

import (
	"context"
	"net/url"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

// Tengo virtual machine.
type Tengo struct {
	match *tengo.Compiled
	fib   *tengo.Compiled
}

// NewTengo creates a Tengo instance.
func NewTengo(matchSrc, fibSrc string) (*Tengo, error) {
	libs := stdlib.GetModuleMap(stdlib.AllModuleNames()...)
	libs.AddBuiltinModule("url", urlModule)

	matchScript := tengo.NewScript([]byte(matchSrc))
	matchScript.SetImports(libs)
	matchCompiled, err := matchScript.Compile()
	if err != nil {
		return nil, err
	}

	fibScript := tengo.NewScript([]byte(fibSrc))
	fibScript.SetImports(libs)
	fibCompiled, err := fibScript.Compile()
	if err != nil {
		return nil, err
	}

	return &Tengo{
		match: matchCompiled,
		fib:   fibCompiled,
	}, nil
}

// Match checks whether the obj parameter satisfies some rules.
func (vm *Tengo) Match(ctx context.Context, obj interface{}) bool {
	_ = vm.match.Set("obj", obj)
	if err := vm.match.RunContext(ctx); err != nil {
		return false
	}
	return vm.match.Get("is_match").Bool()
}

// Match checks whether the obj parameter satisfies some rules.
func (vm *Tengo) Fib(ctx context.Context, n int) int {
	_ = vm.fib.Set("n", n)
	if err := vm.fib.RunContext(ctx); err != nil {
		return -1
	}
	return vm.fib.Get("fib_result").Int()
}

var urlModule = map[string]tengo.Object{
	"decode_query": &tengo.UserFunction{
		Name:  "decode_query",
		Value: urlDecodeQuery,
	}, // decode_query(str) => map
	"encode_query": &tengo.UserFunction{
		Name:  "encode_query",
		Value: urlEncodeQuery,
	}, // encode_query(map) => str
}

// urlDecodeQuery is Tengo version of 'url.ParseQuery' function.
func urlDecodeQuery(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 1 {
		return nil, tengo.ErrWrongNumArguments
	}

	qstr, ok := tengo.ToString(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	qs, err := url.ParseQuery(qstr)
	if err != nil {
		return tengoError(err), nil
	}

	m := &tengo.Map{Value: make(map[string]tengo.Object)}
	for k, q := range qs {
		// In most cases, the parameters shouldn't be duplicated. So
		// we only gets the first value associated with the given key.
		m.Value[k] = &tengo.String{Value: q[0]}
	}

	return m, nil
}

// urlEncodeQuery is Tengo version of 'url.Values.Encode' method.
func urlEncodeQuery(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 1 {
		return nil, tengo.ErrWrongNumArguments
	}

	m, ok := args[0].(*tengo.Map)
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "map(compatible)",
			Found:    args[0].TypeName(),
		}
	}

	vs := url.Values{}
	for k, v := range m.Value {
		// Skip non-string values.
		if v, ok := tengo.ToString(v); ok {
			vs.Add(k, v)
		}
	}

	return &tengo.String{Value: vs.Encode()}, nil
}

// tengoError converts an Go error to the corresponded Tengo error.
func tengoError(err error) tengo.Object {
	if err == nil {
		return tengo.TrueValue
	}
	return &tengo.Error{Value: &tengo.String{Value: err.Error()}}
}
