// Package benchmark contains test suits for Go embedded script engine.
package benchmark

import (
	"context"
	"testing"
)

// VM virtual machine for embedded script.
type VM interface {
	// Match checks whether the obj parameter satisfies some rules.
	Match(ctx context.Context, obj interface{}) bool

	// Fib computes fibonacci number(n) then returns it.
	Fib(ctx context.Context, n int) int
}

// BenchmarkMatch benchmark VM.Match method.
func BenchmarkMatch(b *testing.B, vm VM) {
	ctx := context.Background()
	obj := struct {
		Name   string  `json:"name"`
		Age    int     `json:"age"`
		Gender string  `json:"gender"`
		Height float64 `json:"height"`
		Grade  string  `json:"grade"`
	}{
		Name:   "Tim",
		Age:    27,
		Gender: "male",
		Height: 180.5,
		Grade:  "math=A&physics=S&english=B&history=C",
	}

	for n := 0; n < b.N; n++ {
		_ = vm.Match(ctx, obj)
	}
}

// BenchmarkFib benchmark VM.Fib method.
func BenchmarkFib(b *testing.B, vm VM) {
	var ctx = context.Background()
	for n := 0; n < b.N; n++ {
		_ = vm.Fib(ctx, 35)
	}
}
