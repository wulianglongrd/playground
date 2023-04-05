package string

import "testing"

/*
```shell
$ go version
go version go1.20.2 darwin/amd64
$ go test -run="none" -bench="BenchmarkContact" -test.benchmem .
goos: darwin
goarch: amd64
pkg: github.com/wulianglongrd/playground/string
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkContact/contact:both-short-12         	58856082	        19.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkContact/contact:short-long-12         	25647933	        42.78 ns/op	      80 B/op	       1 allocs/op
BenchmarkContact/contact:long-short-12         	28128960	        42.46 ns/op	      80 B/op	       1 allocs/op
BenchmarkContact/contact:both-long-12          	24325188	        49.18 ns/op	     144 B/op	       1 allocs/op
BenchmarkContact/builder:both-short-12         	35477517	        33.42 ns/op	      32 B/op	       1 allocs/op
BenchmarkContact/builder:short-long-12         	29323489	        39.96 ns/op	      80 B/op	       1 allocs/op
BenchmarkContact/builder:long-short-12         	29771395	        40.21 ns/op	      80 B/op	       1 allocs/op
BenchmarkContact/builder:both-long-12          	24687818	        47.19 ns/op	     144 B/op	       1 allocs/op
PASS
ok  	github.com/wulianglongrd/playground/string	11.626s
```
*/
func BenchmarkContact(b *testing.B) {
	short := "short-string-15"
	long := "long-string-common-plat-public-passport-userinfo-ul-offline-l-70"

	// contact
	b.Run("contact:both-short", func(b *testing.B) {
		benchmarkContact(b, short, short)
	})
	b.Run("contact:short-long", func(b *testing.B) {
		benchmarkContact(b, short, long)
	})
	b.Run("contact:long-short", func(b *testing.B) {
		benchmarkContact(b, long, short)
	})
	b.Run("contact:both-long", func(b *testing.B) {
		benchmarkContact(b, long, long)
	})

	// builder
	b.Run("builder:both-short", func(b *testing.B) {
		benchmarkBuilder(b, short, short)
	})
	b.Run("builder:short-long", func(b *testing.B) {
		benchmarkBuilder(b, short, long)
	})
	b.Run("builder:long-short", func(b *testing.B) {
		benchmarkBuilder(b, long, short)
	})
	b.Run("builder:both-long", func(b *testing.B) {
		benchmarkBuilder(b, long, long)
	})
}

func benchmarkContact(b *testing.B, s1, s2 string) {
	for i := 0; i < b.N; i++ {
		Contact(s1, s2)
	}
}

func benchmarkBuilder(b *testing.B, s1, s2 string) {
	for i := 0; i < b.N; i++ {
		Builder(s1, s2)
	}
}
