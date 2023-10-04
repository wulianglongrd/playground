package trie

type trieChild map[string]*Trie

type Trie struct {
	child trieChild
	data  *string
}

func New() *Trie {
	return &Trie{
		child: make(trieChild),
	}
}

func (t *Trie) Add(host []string, data *string) {
	if len(host) == 0 {
		// note: this is not necessarily a leaf node.
		t.data = data
		return
	}

	key := host[len(host)-1]
	left := host[:len(host)-1]

	child, ok := t.child[key]
	if !ok {
		child = New()
		t.child[key] = child
	}

	child.Add(left, data)
}

// Matches implement `Matches` semantic matching
// https://github.com/istio/istio/blob/master/pkg/config/host/name.go#L37
func (t *Trie) Matches(host []string, out []string) []string {
	// host reach left boundary
	if len(host) == 0 {
		// the tree has reached a leaf node.
		// there means: trie equals host (match),
		// example1, trie: a.com, host: a.com
		// example2, trie: *.a.com, host: *.a.com
		if len(t.child) == 0 {
			return t.getData(out)
		}
		// The tree has not reached a leaf node.
		// there means: trie is longer than host (not match),
		// example, trie: {anything}.a.com, host: a.com
		return out
	}

	key := host[len(host)-1]
	left := host[:len(host)-1]

	child, exists := t.child[key]
	if exists {
		return child.Matches(left, out)
	}

	// the tree has reached a leaf node.
	// there means: trie is shorter than host (not match),
	// example, trie: foo.com, host: {anything}.foo.com
	if len(t.child) == 0 {
		return out
	}

	// match
	// example, trie: {anything-but-not-*}.com, host: *.com
	if key == "*" {
		return t.getData(out)
	}

	// match
	// example, trie: *.com, host: {anything-but-not-*}.com
	if _, ok := t.child["*"]; ok {
		return t.getData(out)
	}

	// not match
	// example, trie: a.com, host: b.com
	return out
}

// SubsetOf implement `SubsetOf` semantic matching
// https://github.com/istio/istio/blob/master/pkg/config/host/name.go#L64
func (t *Trie) SubsetOf(host []string, out []string) []string {
	if len(host) == 0 {
		if len(t.child) == 0 {
			return t.getData(out)
		}
		return out
	}

	key := host[len(host)-1]
	left := host[:len(host)-1]

	child, exists := t.child[key]
	if exists {
		return child.SubsetOf(left, out)
	}

	if len(t.child) == 0 {
		return out
	}

	if key == "*" {
		return t.getData(out)
	}

	return out
}

func (t *Trie) getData(out []string) []string {
	if t.data != nil {
		out = append(out, *t.data)
	}
	if len(t.child) == 0 {
		return out
	}

	for _, child := range t.child {
		out = child.getData(out)
	}
	return out
}
