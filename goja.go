package benchmark

import "context"

// Goja javascript virtual machine.
type Goja struct {
}

// Match checks whether the obj parameter satisfies some rules.
func (vm *Goja) Match(ctx context.Context, obj interface{}) bool {
	return false
}

// Fib computes fibonacci number(n) then returns it.
func (vm *Goja) Fib(ctx context.Context, n int) int {
	return 0
}
