// Package benchmark contains test suits for Go embedded script engine.
package benchmark

import "context"

// VM virtual machine for embedded script.
type VM interface {
	// Match checks whether the obj parameter satisfies some rules.
	Match(ctx context.Context, obj interface{}) bool

	// Fib computes fibonacci number(n) then returns it.
	Fib(ctx context.Context, n int) int
}
