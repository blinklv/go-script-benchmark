package benchmark

import "testing"

const ottoScript = `
function match(obj) {
	if (obj.name !== "Tim") {
		return false;
	}

	if (obj.age < 18) {
		return false;
	}

	if (obj.gender === "female") {
		return false;
	}

	if (obj.height < 165.0) {
		return false;
	}

	var grade = parse_query(obj.grade);
	if (grade.math[0] === "B") {
		return false;
	}

	return true;
}

function fib(n) {
	var a = 0, b= 1;
	for (var i = 1; i < n; ++i) {
		var t = a + b;
		a = b;
		b = t;
	}

	return a;
}
`

func TestOtto(t *testing.T) {
	vm, err := NewOtto(ottoScript)
	if err != nil {
		t.Fatalf("NewOtto failed: %v", err)
	}
	TestMatch(t, vm)
	TestFib(t, vm)
}

func BenchmarkOttoMatch(b *testing.B) {
	vm, err := NewOtto(ottoScript)
	if err != nil {
		b.Fatalf("NewOtto failed: %v", err)
	}
	BenchmarkMatch(b, vm)
}

func BenchmarkOttoFib(b *testing.B) {
	vm, err := NewOtto(ottoScript)
	if err != nil {
		b.Fatalf("NewOtto failed: %v", err)
	}
	BenchmarkFib(b, vm)
}
