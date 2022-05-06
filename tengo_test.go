package benchmark

import "testing"

const tengoMatchScript = `
url := import("url")

match := func(obj) {
	if obj["name"] != "Tim" {
		return false
	}

	if obj["age"] < 18 {
		return false
	}

	if obj["gender"] == "female" {
		return false
	}

	if (obj.height < 165.0) {
		return false
	}

	grade := url.decode_query(obj["grade"])
	if grade["math"] == "B" {
		return false
	}

	return true
}

is_match := match(obj)
`

const tengoFibScript = `
fib := func(n) {
	a := 0
	b := 1
	for i := 1; i < n; i++ {
		t := a + b
		a = b
		b = t
	}
	return a
}

fib_result := fib(n)
`

func TestTengo(t *testing.T) {
	vm, err := NewTengo(tengoMatchScript, tengoFibScript)
	if err != nil {
		t.Fatalf("NewTengo failed: %v", err)
	}
	TestMatch(t, vm)
	TestFib(t, vm)
}

func BenchmarkTengoMatch(b *testing.B) {
	vm, err := NewTengo(tengoMatchScript, tengoFibScript)
	if err != nil {
		b.Fatalf("NewTengo failed: %v", err)
	}
	BenchmarkMatch(b, vm)
}

func BenchmarkTengoFib(b *testing.B) {
	vm, err := NewTengo(tengoMatchScript, tengoFibScript)
	if err != nil {
		b.Fatalf("NewTengo failed: %v", err)
	}
	BenchmarkFib(b, vm)
}
