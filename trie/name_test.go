// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trie

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

// test cases copy from istio source code, updated and covered trie implementation
// https://github.com/istio/istio/blob/master/pkg/config/host/name_test.go

func TestNameMatches(t *testing.T) {
	tests := []struct {
		name string
		a, b Name
		out  bool
	}{
		{"empty", "", "", true},
		{"first empty", "", "foo.com", false},
		{"second empty", "foo.com", "", false},

		{
			"non-wildcard domain",
			"foo.com", "foo.com", true,
		},
		{
			"non-wildcard domain",
			"bar.com", "foo.com", false,
		},
		{
			"non-wildcard domain - order doesn't matter",
			"foo.com", "bar.com", false,
		},

		{
			"domain does not match subdomain",
			"bar.foo.com", "foo.com", false,
		},
		{
			"domain does not match subdomain - order doesn't matter",
			"foo.com", "bar.foo.com", false,
		},

		{
			"wildcard matches subdomains",
			"*.com", "foo.com", true,
		},
		{
			"wildcard matches subdomains",
			"*.com", "bar.com", true,
		},
		{
			"wildcard matches subdomains",
			"*.foo.com", "bar.foo.com", true,
		},

		{"wildcard matches anything", "*", "foo.com", true},
		{"wildcard matches anything", "*", "*.com", true},
		{"wildcard matches anything", "*", "com", true},
		{"wildcard matches anything", "*", "*", true},
		{"wildcard matches anything", "*", "", true},

		{"wildcarded domain matches wildcarded subdomain", "*.com", "*.foo.com", true},
		{"wildcarded sub-domain does not match domain", "foo.com", "*.foo.com", false},
		{"wildcarded sub-domain does not match domain - order doesn't matter", "*.foo.com", "foo.com", false},

		{"long wildcard does not match short host", "*.foo.bar.baz", "baz", false},
		{"long wildcard does not match short host - order doesn't matter", "baz", "*.foo.bar.baz", false},
		{"long wildcard matches short wildcard", "*.foo.bar.baz", "*.baz", true},
		{"long name matches short wildcard", "foo.bar.baz", "*.baz", true},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, tt.name), func(t *testing.T) {
			if tt.out != tt.a.Matches(tt.b) {
				t.Fatalf("%q.Matches(%q) = %t wanted %t", tt.a, tt.b, !tt.out, tt.out)
			}

			// test trie implementation
			trie := New()
			sa := string(tt.a)
			trie.Add(toSlice(tt.a), &sa)

			got := make([]string, 0)
			got = trie.Matches(toSlice(tt.b), got)
			if tt.out != (len(got) > 0) {
				t.Fatalf("trie: %q.Matches(%q) = %t wanted %t", tt.a, tt.b, !tt.out, tt.out)
			}
			if len(got) > 0 && got[0] != string(tt.a) {
				t.Fatalf("trie: %q.Matches(%q) = %s wanted %s", tt.a, tt.b, got[0], tt.a)
			}
		})
	}
}

func TestNameSubsetOf(t *testing.T) {
	tests := []struct {
		name string
		a, b Name
		out  bool
	}{
		{"empty", "", "", true},
		{"first empty", "", "foo.com", false},
		{"second empty", "foo.com", "", false},

		{
			"non-wildcard domain",
			"foo.com", "foo.com", true,
		},
		{
			"non-wildcard domain",
			"bar.com", "foo.com", false,
		},
		{
			"non-wildcard domain - order doesn't matter",
			"foo.com", "bar.com", false,
		},

		{
			"domain does not match subdomain",
			"bar.foo.com", "foo.com", false,
		},
		{
			"domain does not match subdomain - order doesn't matter",
			"foo.com", "bar.foo.com", false,
		},

		{
			"wildcard matches subdomains",
			"foo.com", "*.com", true,
		},
		{
			"wildcard matches subdomains",
			"bar.com", "*.com", true,
		},
		{
			"wildcard matches subdomains",
			"bar.foo.com", "*.foo.com", true,
		},

		{"wildcard matches anything", "foo.com", "*", true},
		{"wildcard matches anything", "*.com", "*", true},
		{"wildcard matches anything", "com", "*", true},
		{"wildcard matches anything", "*", "*", true},
		{"wildcard matches anything", "", "*", true},

		{"wildcarded domain matches wildcarded subdomain", "*.foo.com", "*.com", true},
		{"wildcarded sub-domain does not match domain", "*.foo.com", "foo.com", false},

		{"long wildcard does not match short host", "*.foo.bar.baz", "baz", false},
		{"long name matches short wildcard", "foo.bar.baz", "*.baz", true},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("[%d] %s", idx, tt.name), func(t *testing.T) {
			if tt.out != tt.a.SubsetOf(tt.b) {
				t.Fatalf("%q.SubsetOf(%q) = %t wanted %t", tt.a, tt.b, !tt.out, tt.out)
			}

			// test trie implementation
			trie := New()
			sa := string(tt.a)
			trie.Add(toSlice(tt.a), &sa)

			got := make([]string, 0)
			got = trie.SubsetOf(toSlice(tt.b), got)
			if tt.out != (len(got) > 0) {
				t.Fatalf("trie: %q.Matches(%q) = %t wanted %t", tt.a, tt.b, !tt.out, tt.out)
			}
			if len(got) > 0 && got[0] != string(tt.a) {
				t.Fatalf("trie: %q.Matches(%q) = %s wanted %s", tt.a, tt.b, got[0], tt.a)
			}
		})
	}
}

/*
Note: the benchmark test results are closely related to the number of "vsHosts" and also related to the "differentiation of egressHosts".

❯ go test -run="none" -bench="BenchmarkMatches" -test.benchmem -count=1
goos: darwin
goarch: amd64
pkg: github.com/wulianglongrd/playground/trie
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMatches/old_matches-12         	  799670	      1267 ns/op	       0 B/op	       0 allocs/op
BenchmarkMatches/trie_matches-12        	 2562811	       467.8 ns/op	      16 B/op	       1 allocs/op
BenchmarkMatches/old_subsetOf-12        	  943035	      1294 ns/op	       0 B/op	       0 allocs/op
BenchmarkMatches/trie_subsetOf-12       	 2725290	       435.2 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/wulianglongrd/playground/trie	5.935s
*/
func BenchmarkMatches(b *testing.B) {
	egressHosts := []string{
		"v1.productpage.cluster.local.svc",
		"v1.reviews.cluster.local.svc",
		"v2.reviews.cluster.local.svc",
		"v3.reviews.cluster.local.svc",
		"v1.details.cluster.local.svc",
		"v1.ratings.cluster.local.svc",
		"*.wildcard.com",
	}
	vsHosts := make([]string, 0)
	vsHosts = append(vsHosts, egressHosts...)
	for i := 0; i < 3; i++ {
		for _, h := range egressHosts {
			if strings.HasPrefix(h, "*") {
				vsHosts = append(vsHosts, fmt.Sprintf("*.%d%s", i, h[1:]))
				continue
			}
			vsHosts = append(vsHosts, fmt.Sprintf("%d.%s", i, h))
		}
	}
	egressHosts = append(egressHosts, "notexists.com", "*.notexists.com")

	// shuffle the slice
	rand.Shuffle(len(vsHosts), func(i, j int) {
		vsHosts[i], vsHosts[j] = vsHosts[j], vsHosts[i]
	})
	rand.Shuffle(len(egressHosts), func(i, j int) {
		egressHosts[i], egressHosts[j] = egressHosts[j], egressHosts[i]
	})

	b.ResetTimer()
	b.Run("old matches", func(b *testing.B) {
		for _, eh := range egressHosts {
			for i := 0; i < b.N; i++ {
				for _, vh := range vsHosts {
					Name(vh).Matches(Name(eh))
				}
			}
		}
	})

	b.Run("trie matches", func(b *testing.B) {
		// build trie
		tr := New()
		for _, vh := range vsHosts {
			tr.Add(strings.Split(vh, "."), &vh)
		}

		// test matches
		for _, eh := range egressHosts {
			a2 := strings.Split(eh, ".")
			out := make([]string, 0)
			for i := 0; i < b.N; i++ {
				out = out[:0]
				tr.Matches(a2, out)
			}
		}
	})

	b.Run("old subsetOf", func(b *testing.B) {
		for _, eh := range egressHosts {
			for i := 0; i < b.N; i++ {
				for _, vh := range vsHosts {
					Name(vh).SubsetOf(Name(eh))
				}
			}
		}
	})

	b.Run("trie subsetOf", func(b *testing.B) {
		// build trie
		tr := New()
		for _, vh := range vsHosts {
			tr.Add(strings.Split(vh, "."), &vh)
		}

		// test subsetOf matches
		for _, eh := range egressHosts {
			a2 := strings.Split(eh, ".")
			out := make([]string, 0)
			for i := 0; i < b.N; i++ {
				out = out[:0]
				tr.SubsetOf(a2, out)
			}
		}
	})
}

func toSlice(a Name) []string {
	return strings.Split(string(a), ".")
}
