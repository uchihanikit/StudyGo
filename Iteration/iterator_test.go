package iteration

import "testing"

func TestIterator(t *testing.T) {
	got := iterator("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iterator("a")
	}
}
