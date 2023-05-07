package demos

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

// ❯ go test -run="none" -bench="BenchmarkMerge" -test.benchmem -count=5 . > old.txt
// here: replace the method called
// ❯ go test -run="none" -bench="BenchmarkMerge" -test.benchmem -count=5 . > new.txt
// ❯ benchstat old.txt new.txt
// name      old time/op    new time/op    delta
// Merge-12    2.47µs ± 5%    0.96µs ± 7%  -61.35%  (p=0.008 n=5+5)
//
// name      old alloc/op   new alloc/op   delta
// Merge-12    11.4kB ± 0%     2.7kB ± 0%  -76.37%  (p=0.008 n=5+5)
//
// name      old allocs/op  new allocs/op  delta
// Merge-12      17.0 ± 0%       1.0 ± 0%  -94.12%  (p=0.008 n=5+5)
func BenchmarkMerge(b *testing.B) {
	nums := genNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SplitAndMerge(nums)
		//ForLoopAndMerge(nums)
	}
}

func BenchmarkSplitAndMerge(b *testing.B) {
	nums := genNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SplitAndMerge(nums)
	}
}

func BenchmarkForLoopAndMerge(b *testing.B) {
	nums := genNums()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForLoopAndMerge(nums)
	}
}

func TestMerge(t *testing.T) {
	nums := genNums()
	a := SplitAndMerge(nums)
	b := ForLoopAndMerge(nums)
	if !assert.Equal(t, a, b) {
		t.Error("not equal")
	}
}

func genNums() []int {
	nums := make([]int, 0, 300)
	for i := 0; i < 300; i++ {
		nums = append(nums, rand.Intn(100))
	}
	return nums
}
