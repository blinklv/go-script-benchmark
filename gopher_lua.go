package benchmark

import (
	"context"
	"fmt"
	"net/url"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

// GopherLua lua virtual machine.
type GopherLua struct {
	L *lua.LState
}

// NewGopherLua creates a GopherLua instance.
func NewGopherLua(text string) (*GopherLua, error) {
	L := lua.NewState()
	L.SetGlobal("parse_query", L.NewFunction(gopherLuaParseQuery))
	if err := L.DoString(text); err != nil {
		return nil, fmt.Errorf("DoString failed: %v", err)
	}
	return &GopherLua{
		L: L,
	}, nil
}

// Match checks whether the obj parameter satisfies some rules.
func (vm *GopherLua) Match(_ context.Context, obj interface{}) bool {
	_ = vm.L.CallByParam(lua.P{
		Fn:   vm.L.GetGlobal("match"),
		NRet: 1,
	}, objToLTable(vm.L, obj))
	res := vm.L.Get(-1)
	vm.L.Pop(1)
	return res == lua.LTrue
}

// Fib computes fibonacci number(n) then returns it.
func (vm *GopherLua) Fib(_ context.Context, n int) int {
	_ = vm.L.CallByParam(lua.P{
		Fn:   vm.L.GetGlobal("fib"),
		NRet: 1,
	}, lua.LNumber(n))
	res := vm.L.Get(-1)
	vm.L.Pop(1)

	f, _ := res.(lua.LNumber)
	return int(f)
}

func gopherLuaParseQuery(L *lua.LState) int {
	query := L.ToString(1)
	if query == "" {
		return 0
	}

	vs, err := url.ParseQuery(query)
	if err != nil {
		return 0
	}

	table := L.NewTable()
	for k, v := range vs {
		ss := L.NewTable()
		for _, s := range v {
			ss.Append(lua.LString(s))
		}
		table.RawSetString(k, ss)
	}
	L.Push(table)
	return 1
}

func objToLTable(L *lua.LState, obj interface{}) *lua.LTable {
	var (
		v     = reflect.ValueOf(obj)
		t     = v.Type()
		table = L.NewTable()
	)
	for i := 0; i < v.NumField(); i++ {
		k := t.Field(i).Tag.Get("json")
		switch value := v.Field(i).Interface().(type) {
		case int:
			table.RawSetString(k, lua.LNumber(value))
		case float64:
			table.RawSetString(k, lua.LNumber(value))
		case string:
			table.RawSetString(k, lua.LString(value))
		}
	}

	return table
}
