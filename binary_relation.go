package guile

type BinaryRelation interface {
	Universe() Set
	ContainsRelation(Element, Element) bool
}

type PhysicalBinaryRelation interface {
	BinaryRelation
	AddRelation(Element, Element)
	RemoveRelation(Element, Element)
}

func NewPhysicalBinaryRelationOn(universe Set) PhysicalBinaryRelation {
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

func Reflexive(b BinaryRelation) bool {
	for _, e := range b.Universe().Elements() {
		if !b.ContainsRelation(e, e) {
			return false
		}
	}

	return true
}

func Complete(b BinaryRelation) bool {
	elems := b.Universe().Elements()

	// n^2! yuck!
	for _, x := range elems {
		for _, y := range elems {
			if !(b.ContainsRelation(x, y) || b.ContainsRelation(y, x)) {
				return false
			}
		}
	}

	return true
}

func Transitive(b BinaryRelation) bool {
	if !Complete(b) {
		return false
	}

	elems := b.Universe().Elements()

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

func Symmetric(b BinaryRelation) bool {
	elems := b.Universe().Elements()

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

func AntiSymmetric(b BinaryRelation) bool {
	elems := b.Universe().Elements()
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

func WeakOrder(b BinaryRelation) bool {
	return Complete(b) && Transitive(b)
}

func StrictOrder(b BinaryRelation) bool {
	return WeakOrder(b) && AntiSymmetric(b)
}

// --- }}}
