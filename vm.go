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

// TestMatch test VM.Match method.
func TestMatch(t *testing.T, vm VM) {
	ctx := context.Background()
	obj := map[string]interface{}{
		"name":   "Tim",
		"age":    27,
		"gender": "male",
		"height": 180.5,
		"grade":  "math=A&physics=S&english=B&history=C",
	}
	if !vm.Match(ctx, obj) {
		t.Fatalf("vm.Match should should be true")
	}
}

// TestMatch test VM.Fib method.
func TestFib(t *testing.T, vm VM) {
	ctx := context.Background()
	if n := vm.Fib(ctx, 35); n != 5702887 {
		t.Fatalf("vm.Fib(35) = %d is not equal to 5702887", n)
	}
}

// BenchmarkMatch benchmark VM.Match method.
func BenchmarkMatch(b *testing.B, vm VM) {
	ctx := context.Background()
	obj := map[string]interface{}{
		"name":   "Tim",
		"age":    27,
		"gender": "male",
		"height": 180.5,
		"grade":  "math=A&physics=S&english=B&history=C",
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
