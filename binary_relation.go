package guile

type BinaryRelation interface {
	Universe() Set

	AddRelation(Element, Element)
	RemoveRelation(Element, Element)
	ContainsRelation(Element, Element) bool

	Reflexive() bool
	Complete() bool
	Transitive() bool
	Symmetric() bool
	AntiSymmetric() bool
	WeakOrder() bool
	StrictOrder() bool
}

func NewBinaryRelationOn(universe Set) BinaryRelation {
	return &binaryRelation{
		universe:  universe,
		relations: make(map[Element]map[Element]bool),
	}
}

// --- Binary Relation Implementation {{{

type binaryRelation struct {
	universe  Set
	relations map[Element]map[Element]bool
}

func (b *binaryRelation) Universe() Set {
	return b.universe
}

// --- Relation Management {{{

func (b *binaryRelation) AddRelation(e1, e2 Element) {
	assert(b.universe.Contains(e1), "(*binaryRelation).AddRelation: element 1 is not contained in universe")
	assert(b.universe.Contains(e2), "(*binaryRelation).AddRelation: element 2 is not contained in universe")

	var bucket map[Element]bool
	var exists bool

	// Add Normal Relation
	bucket, exists = b.relations[e1]

	if !exists {
		bucket = map[Element]bool{e2: true}
	} else {
		bucket[e2] = true
	}

	b.relations[e1] = bucket
}

func (b *binaryRelation) RemoveRelation(e1, e2 Element) {
	assert(b.universe.Contains(e1), "(*binaryRelation).AddRelation: element 1 is not contained in universe")
	assert(b.universe.Contains(e2), "(*binaryRelation).AddRelation: element 2 is not contained in universe")

	if bucket, exists := b.relations[e1]; exists {
		if _, exists := bucket[e2]; exists {
			delete(bucket, e2)
		}
	}
}

func (b *binaryRelation) ContainsRelation(e1, e2 Element) bool {
	assert(b.universe.Contains(e1), "(*binaryRelation).AddRelation: element 1 is not contained in universe")
	assert(b.universe.Contains(e2), "(*binaryRelation).AddRelation: element 2 is not contained in universe")

	if bucket, exists := b.relations[e1]; exists {
		if _, defined := bucket[e2]; defined {
			return true
		}
	}

	return false
}

// --- }}}

// --- Properties {{{
func (b *binaryRelation) Reflexive() bool {
	for _, e := range b.universe.Elements() {
		if !b.ContainsRelation(e, e) {
			return false
		}
	}

	return true
}

func (b *binaryRelation) Complete() bool {
	elems := b.universe.Elements()

	// n^2! yuck!
	for _, e := range elems {
		for _, eprime := range elems {
			if !b.ContainsRelation(e, eprime) {
				return false
			}
		}
	}

	return true
}

func (b *binaryRelation) Transitive() bool {
	if !b.Complete() {
		return false
	}

	elems := b.universe.Elements()

	for _, x := range elems {
		for _, y := range elems {
			for _, z := range elems {
				if b.ContainsRelation(x, y) && b.ContainsRelation(y, z) {
					if !b.ContainsRelation(x, z) {
						return false
					}
				}
			}
		}
	}

	return true
}

func (b *binaryRelation) Symmetric() bool {
	elems := b.universe.Elements()
	for _, x := range elems {
		for _, y := range elems {
			if b.ContainsRelation(x, y) {
				if !b.ContainsRelation(x, y) {
					return false
				}
			}
		}
	}

	return true
}

func (b *binaryRelation) AntiSymmetric() bool {
	elems := b.universe.Elements()
	for _, x := range elems {
		for _, y := range elems {
			if b.ContainsRelation(x, y) && b.ContainsRelation(y, x) {
				if x != y {
					return false
				}
			}
		}
	}

	return true
}

// --- }}}

func (b *binaryRelation) WeakOrder() bool {
	return b.Complete() && b.Transitive()
}

func (b *binaryRelation) StrictOrder() bool {
	return b.WeakOrder() && b.AntiSymmetric()
}

// --- }}}
