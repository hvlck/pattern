package pattern

import (
	"os"
	"testing"
)

func loadFile(location string) []byte {
	f, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}

	return f
}

func TestSame(t *testing.T) {
	one, two := loadFile("./examples/one.html"), loadFile("./examples/two.html")
	nodes, score, err := Compare(one, two)
	if err != nil {
		t.Fatalf("error comparing files: %v", err)
	}

	if len(nodes) != 0 || score != 1 {
		t.Fatalf("files not equal")
	}
}

func TestDifferent(t *testing.T) {
	one, two := loadFile("./examples/one.html"), loadFile("./examples/three.html")
	nodes, score, err := Compare(one, two)
	if err != nil {
		t.Fatalf("error comparing files: %v", err)
	}

	if len(nodes) != 0 || score != 1 {
		t.Fatalf("files not equal")
	}
}

func BenchmarkSame(b *testing.B) {
	b.StopTimer()
	one, two := loadFile("./examples/one.html"), loadFile("./examples/two.html")
	b.StartTimer()
	nodes, score, err := Compare(one, two)
	if err != nil {
		b.Fatalf("error comparing files: %v", err)
	}

	if len(nodes) != 0 || score != 1 {
		b.Fatalf("files not equal")
	}
}
