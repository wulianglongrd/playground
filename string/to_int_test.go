package string

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

/**
```shell
$ go version
go version go1.20.2 darwin/amd64
$ go test -run="none" -bench="BenchmarkConvertToInt" -test.benchmem .
goos: darwin
goarch: amd64
pkg: github.com/wulianglongrd/playground/string
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkConvertToInt/small-ParseInt-12         	96055290	        12.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvertToInt/small-Atoi-12             	236436679	         5.048 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvertToInt/small-Sscan-12            	 3183345	       371.5 ns/op	      80 B/op	       4 allocs/op
BenchmarkConvertToInt/big-ParseInt-12           	44240553	        27.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvertToInt/big-Atoi-12               	42489405	        28.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvertToInt/big-Sscan-12              	 1544557	       769.8 ns/op	      96 B/op	       4 allocs/op
PASS
ok  	github.com/wulianglongrd/playground/string	10.256s
```
**/

func BenchmarkConvertToInt(b *testing.B) {
	smallNum := "8080"
	bigNum := strconv.Itoa(math.MaxInt - 100)

	b.Run("small-ParseInt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.ParseInt(smallNum, 10, 64)
		}
	})
	b.Run("small-Atoi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.Atoi(smallNum)
		}
	})
	b.Run("small-Sscan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var o int
			_, _ = fmt.Sscanf(smallNum, "%d", &o)
		}
	})

	b.Run("big-ParseInt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.ParseInt(bigNum, 10, 64)
		}
	})
	b.Run("big-Atoi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.Atoi(bigNum)
		}
	})
	b.Run("big-Sscan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var o int
			_, _ = fmt.Sscanf(bigNum, "%d", &o)
		}
	})
}
