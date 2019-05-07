package types

import "testing"

func TestSeqGetMapBenchmark(t *testing.T) {
	gc := true
	SeqGetMapBenchmark(1000, gc)
	SeqGetMapBenchmark(10000, gc)
	SeqGetMapBenchmark(100000, gc)
	SeqGetMapBenchmark(1000000, gc)
	SeqGetMapBenchmark(10000000, gc)
}
