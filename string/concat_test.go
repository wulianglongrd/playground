package string

import "testing"

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
