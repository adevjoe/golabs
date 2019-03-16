package bench

import "testing"

func TestPrintSize(t *testing.T) {

}

func BenchmarkPrintSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintSize("int8", int8(0))
	}
}
