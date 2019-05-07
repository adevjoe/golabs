package testing

import "testing"

func TestPrintSize(t *testing.T) {
	PrintSize("int8", int8(0))
}

func BenchmarkPrintSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintSize("int8", int8(0))
	}
}
