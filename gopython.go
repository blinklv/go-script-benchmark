package benchmark

import (
	"context"
	"fmt"
	"net/url"
	"reflect"
	"runtime"

	_ "github.com/go-python/gpython/modules"
	"github.com/go-python/gpython/py"
)

// GoPython python virtual machine.
type GoPython struct {
	ctx py.Context
	mod *py.Module
}

// NewGoPython creates g GoPython instance.
func NewGoPython(text string) (*GoPython, error) {
	ctx := py.NewContext(py.DefaultContextOpts())
	mod, err := ctx.ModuleInit(&py.ModuleImpl{
		CodeSrc: text,
		Info: py.ModuleInfo{
			Name: "benchmark",
			Doc:  "embedded python benchmark module",
		},
		Methods: []*py.Method{
			py.MustNewMethod("parse_query", gopythonParseQuery, 0, ""),
		},
		Globals: py.StringDict{
			"PY_VERSION": py.String("Python 3.4 (github.com/go-python/gpython)"),
			"GO_VERSION": py.String(fmt.Sprintf("%s on %s %s", runtime.Version(), runtime.GOOS, runtime.GOARCH)),
			"MYLIB_VERS": py.String("Vacation 1.0 by Fletch F. Fletcher"),
		},
		OnContextClosed: func(instance *py.Module) {
			fmt.Print("<<< host py.Context of py.Module instance closing >>>\n+++\n")
		},
	})
	if err != nil {
		return nil, err
	}

	return &GoPython{
		ctx: ctx,
		mod: mod,
	}, nil
}

// Match checks whether the obj parameter satisfies some rules.
func (vm *GoPython) Match(_ context.Context, obj interface{}) bool {
	res, err := vm.mod.Call("match", py.Tuple{objToPyDict(obj)}, nil)
	if err != nil {
		panic(err)
	}
	b, _ := res.(py.Bool)
	return bool(b)
}

// Fib computes fibonacci number(n) then returns it.
func (vm *GoPython) Fib(_ context.Context, n int) int {
	res, err := vm.mod.Call("fib", py.Tuple{py.Int(n)}, nil)
	if err != nil {
		panic(err)
	}
	i, _ := res.(py.Int)
	return int(i)
}

func gopythonParseQuery(_ py.Object, args py.Tuple) (py.Object, error) {
	var query py.String
	if err := py.LoadTuple(args, []interface{}{&query}); err != nil {
		return nil, err
	}

	vs, err := url.ParseQuery(string(query))
	if err != nil {
		return nil, err
	}

	dict := py.NewStringDict()
	for k, v := range vs {
		list := py.NewList()
		for _, s := range v {
			list.Items = append(list.Items, py.String(s))
		}
		dict[k] = list
	}

	return dict, nil
}

func objToPyDict(obj interface{}) py.StringDict {
	var (
		v    = reflect.ValueOf(obj)
		t    = v.Type()
		dict = py.NewStringDict()
	)

	for i := 0; i < v.NumField(); i++ {
		k := t.Field(i).Tag.Get("json")
		switch value := v.Field(i).Interface().(type) {
		case int:
			dict[k] = py.Int(int64(value))
		case float64:
			dict[k] = py.Float(value)
		case string:
			dict[k] = py.String(value)
		}
	}
	return dict
}
