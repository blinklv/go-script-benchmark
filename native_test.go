package benchmark

import "testing"

func TestNative(t *testing.T) {
	vm := &Native{}
	TestMatch(t, vm)
	TestFib(t, vm)
}

func BenchmarkNativeMatch(b *testing.B) {
	vm := &Native{}
	BenchmarkMatch(b, vm)
}

func BenchmarkNativeFib(b *testing.B) {
	vm := &Native{}
	BenchmarkFib(b, vm)
}
