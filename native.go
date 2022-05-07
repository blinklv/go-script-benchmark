package benchmark

import (
	"context"
	"net/url"
)

// Native golang native virtual machine.
type Native struct{}

// Match checks whether the obj parameter satisfies some rules.
func (vm *Native) Match(_ context.Context, obj interface{}) bool {
	m, ok := obj.(map[string]interface{})
	if !ok {
		return false
	}

	if name, _ := m["name"].(string); name != "Tim" {
		return false
	}

	if age, _ := m["age"].(int); age < 18 {
		return false
	}

	if gender, _ := m["gender"].(string); gender == "female" {
		return false
	}

	if height, _ := m["height"].(float64); height < 165.0 {
		return false
	}

	if grade, ok := m["grade"].(string); ok {
		if vs, _ := url.ParseQuery(grade); vs.Get("math") == "B" {
			return false
		}
	}

	return true
}

// Fib computes fibonacci number(n) then returns it.
func (vm *Native) Fib(_ context.Context, n int) int {
	var a, b = 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
