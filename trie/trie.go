package trie

type trieChild map[string]*Trie

type Trie struct {
	child trieChild
	data  string
}

func New() *Trie {
	return &Trie{
		child: make(trieChild),
	}
}

func (t *Trie) Add(host []string, data string) {
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

func (t *Trie) Matches(host []string, out []string) []string {
	// host reach left boundary
	if len(host) == 0 {
		// the tree has reached a leaf node
		// trie: a.com
		// host: a.com
		if len(t.child) == 0 {
			return t.getData(out)
		}
		// The tree has not reached a leaf node
		// trie: a.com
		// host: xx.a.com
		return out
	}

	key := host[len(host)-1]
	left := host[:len(host)-1]

	child, exists := t.child[key]
	if exists {
		return child.Matches(left, out)
	}

	// the tree has reached a leaf node
	// trie: foo.com
	// host: *.foo.com
	if len(t.child) == 0 {
		return out
	}

	// trie: foo.com
	// host: *.com
	if key == "*" {
		return t.getData(out)
	}

	// trie: *.com
	// host: foo.com
	if _, ok := t.child["*"]; ok {
		return t.getData(out)
	}

	return out
}

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
	if t.data != "" {
		out = append(out, t.data)
	}
	if len(t.child) == 0 {
		return out
	}

	for _, child := range t.child {
		out = child.getData(out)
	}
	return out
}
