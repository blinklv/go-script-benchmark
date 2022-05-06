package benchmark

import (
	"testing"
)

const gopherLuaScript = `
function match(obj)
    if (obj["name"] ~= "Tim")
    then
        return false
    end

    if (obj["age"] < 18)
    then
        return false
    end

    if (obj["gender"] == "female")
    then
        return false
    end

    local grade = parse_query(obj["grade"])
    if (grade["math"][1] == "B")
    then
        return false
    end

    return true;
end

function fib(n)
    local a, b = 0, 1
    for i = 1, n-1 do
        a, b = b, a + b
    end
    return a
end
`

func TestGopherLua(t *testing.T) {
	vm, err := NewGopherLua(gopherLuaScript)
	if err != nil {
		t.Fatalf("NewGopherLua failed: %v", err)
	}
	TestMatch(t, vm)
	TestFib(t, vm)
}

func BenchmarkGopherLuaMatch(b *testing.B) {
	vm, err := NewGopherLua(gopherLuaScript)
	if err != nil {
		b.Fatalf("NewGopherLua failed: %v", err)
	}
	BenchmarkMatch(b, vm)
}

func BenchmarkGopherLuaFib(b *testing.B) {
	vm, err := NewGopherLua(gopherLuaScript)
	if err != nil {
		b.Fatalf("NewGopherLua failed: %v", err)
	}
	BenchmarkFib(b, vm)
}
