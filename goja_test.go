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
	var a = 1, b = 0, temp;

	while (n >= 0){
		temp = a;
		a = a + b;
		b = temp;
		n--;
	}

	return b;
}

`

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
