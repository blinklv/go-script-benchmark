package benchmark

import "testing"

const gopythonScript = `
def match(obj):
	if obj["name"] != "Tim":
		return False

	if obj["age"] < 18:
		return False

	if obj["gender"] == "female":
		return False

	grade = parse_query(obj["grade"])
	if grade["math"][0] == "B":
		return True

	return True

def fib(n):
	a, b = 0, 1
	for i in range(1, n):
		a, b = b, a + b
	return a
`

func BenchmarkGoPythonMatch(b *testing.B) {
	vm, err := NewGoPython(gopythonScript)
	if err != nil {
		b.Fatalf("NewGoPython failed: %v", err)
	}
	BenchmarkMatch(b, vm)
}

func BenchmarkGoPythonFib(b *testing.B) {
	vm, err := NewGoPython(gopythonScript)
	if err != nil {
		b.Fatalf("NewGoPython failed: %v", err)
	}
	BenchmarkFib(b, vm)
}
