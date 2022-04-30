package benchmark

import (
	"context"
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
	benchmarkMatch(b, vm)
}

func BenchmarkGojaFib(b *testing.B) {
	vm, err := NewGoja(gojaScript)
	if err != nil {
		b.Fatalf("NewGoja failed: %v", err)
	}
	benchmarkFib(b, vm)
}

func benchmarkMatch(b *testing.B, vm VM) {
	var (
		ctx = context.Background()
		obj = struct {
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
	)
	for n := 0; n < b.N; n++ {
		_ = vm.Match(ctx, obj)
	}
}

func benchmarkFib(b *testing.B, vm VM) {
	var ctx = context.Background()
	for n := 0; n < b.N; n++ {
		_ = vm.Fib(ctx, 35)
	}
}
