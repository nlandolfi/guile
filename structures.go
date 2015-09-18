package guile

func assert(flag bool, s string) {
	if !flag {
		panic(s)
	}
}

// --- Set Interface {{{

type Element interface{}

type VirtualSet interface {
	Contains(Element) bool
}

type Set interface {
	VirtualSet

	Add(Element) bool
	Remove(Element) bool
	Cardinality() uint

	Elements() []Element
}

type Order int

const (
	Left Order = iota
	Right
)

type BinaryRelation interface {
	Universe() Set

	AddRelation(Element, Element, Order)
	RemoveRelation(Element, Element)
	ContainsRelation(Element, Element) bool
	OrderOf(Element, Element) Order

	Reflexive() bool
	Complete() bool
	Transitive() bool
	Symmetric() bool
	AntiSymmetric() bool
	WeakOrder() bool
	StrictOrder() bool
}

// --- }}}

// --- Set Implementation {{{

func NewSet() Set {
	return &set{
		elements: make(map[Element]bool),
		count:    0,
	}
}

type set struct {
	elements map[Element]bool
	count    uint
}

func (s *set) Add(e Element) bool {
	contains := s.Contains(e)
	if !contains {
		s.elements[e] = true
		s.count += 1
	}
	return contains
}

func (s *set) Remove(e Element) bool {
	contains := s.Contains(e)
	if contains {
		delete(s.elements, e)
		s.count -= 1
	}
	return contains
}

func (s *set) Contains(e Element) bool {
	_, ok := s.elements[e]
	return ok
}

func (s *set) Cardinality() uint {
	return s.count
}

func (s *set) Elements() []Element {
	e := make([]Element, 0)
	for k, _ := range s.elements {
		e = append(e, k)
	}
	return e
}

// --- }}}

// --- Binary Relation Implementation {{{

func Inverse(o Order) Order {
	switch o {
	case Left:
		return Right
	case Right:
		return Left
	}

	return -1
}

type binaryRelation struct {
	universe  Set
	relations map[Element]map[Element]Order
}

func NewBinaryRelationOn(universe Set) BinaryRelation {
	return &binaryRelation{
		universe:  universe,
		relations: make(map[Element]map[Element]Order),
	}
}

func (b *binaryRelation) Universe() Set {
	return b.universe
}

func (b *binaryRelation) AddRelation(e1, e2 Element, o Order) {
	assert(b.universe.Contains(e1), "(*binaryRelation).AddRelation: element 1 is not contained in universe")
	assert(b.universe.Contains(e2), "(*binaryRelation).AddRelation: element 2 is not contained in universe")

	var bucket map[Element]Order
	var exists bool

	// Add Normal Relation
	bucket, exists = b.relations[e1]

	if !exists {
		bucket = map[Element]Order{e2: o}
	} else {
		bucket[e2] = o
	}

	b.relations[e1] = bucket

	// Add Inverse Relation
	bucket, exists = b.relations[e2]

	if !exists {
		bucket = map[Element]Order{e1: Inverse(o)}
	} else {
		bucket[e1] = Inverse(o)
	}

	b.relations[e2] = bucket
}

func (b *binaryRelation) RemoveRelation(e1, e2 Element) {
	assert(b.universe.Contains(e1), "(*binaryRelation).AddRelation: element 1 is not contained in universe")
	assert(b.universe.Contains(e2), "(*binaryRelation).AddRelation: element 2 is not contained in universe")

	if bucket, exists := b.relations[e1]; exists {
		if _, exists := bucket[e2]; exists {
			delete(bucket, e2)
		}
	}

	if bucket, exists := b.relations[e2]; exists {
		if _, exists := bucket[e1]; exists {
			delete(bucket, e1)
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

func (b *binaryRelation) OrderOf(e1, e2 Element) Order {
	assert(b.ContainsRelation(e1, e2), "(*binaryRelation).OrderOf: relation not defined")

	return b.relations[e1][e2]
}

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
				if b.OrderOf(x, y) == Right && b.OrderOf(y, x) == Right {
					if b.OrderOf(x, z) != Right {
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
				if !(b.ContainsRelation(x, y) && b.OrderOf(x, y) == b.OrderOf(y, x)) {
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
			if b.ContainsRelation(x, y) && b.ContainsRelation(y, x) && b.OrderOf(x, y) == Inverse(b.OrderOf(y, x)) {
				if x != y {
					return false
				}
			}
		}
	}

	return true
}

func (b *binaryRelation) WeakOrder() bool {
	return b.Complete() && b.Transitive()
}

func (b *binaryRelation) StrictOrder() bool {
	return b.WeakOrder() && b.AntiSymmetric()
}

// --- }}}

// Math Econ Structures
type (
	Alternative Element

	Alternatives Set

	Preference interface {
		BinaryRelation
		Rational() bool
	}
)

type Relation struct {
	First, Second Alternative
}

type Relations map[Relation]bool

type RelationTree map[Alternative][]Alternative
