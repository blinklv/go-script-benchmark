package benchmark

import (
	"testing"
)

const gojaScript = `
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
	let a = 0, b= 1;

	for (let i = 1; i < n; ++i) {
		let t = a + b;
		a = b;
		b = t;
	}

	return a;
}

`

func TestGoja(t *testing.T) {
	vm, err := NewGoja(gojaScript)
	if err != nil {
		t.Fatalf("NewGoja failed: %v", err)
	}
	TestMatch(t, vm)
	TestFib(t, vm)
}

func BenchmarkGojaMatch(b *testing.B) {
	vm, err := NewGoja(gojaScript)
	if err != nil {
		b.Fatalf("NewGoja failed: %v", err)
	}
	BenchmarkMatch(b, vm)
}

func BenchmarkGojaFib(b *testing.B) {
	vm, err := NewGoja(gojaScript)
	if err != nil {
		b.Fatalf("NewGoja failed: %v", err)
	}
	BenchmarkFib(b, vm)
}
