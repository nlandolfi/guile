package guile

// An Element can be any type. It should be noted
// however that the type of various Elements
// added to a set should probably not change. Though
// you can do this if you want.
type Element interface{}

// A VirtualSet is an abstraction over a physical set.
// Think of this as a set which is infinite and membership
// is defined by a predicate. i.e., the set of primes
type VirtualSet interface {
	Contains(Element) bool
}

// A Set represents the classical collection of elements
// as established in set theory. It of course inherits the
// from the VirtualSet interface; specifically, the Contains
// function. But because it is physical and can be manipulated
// you can also add, remove and ask for all the elements of the
// set. Additionally it is countably finite, so you can ask for
// the Cardinality.
type Set interface {
	VirtualSet

	Add(Element) bool
	Remove(Element) bool
	Cardinality() uint

	Elements() []Element
}

// NewSet constructs a new Set object, using guile's internal
// set structure.
func NewSet() Set {
	return &set{
		elements: make(map[Element]bool),
		count:    0,
	}
}

// NewSetWithElements constructs a new Set object,
// contains the elements in the slice, 'elements'
func NewSetWithElements(elements []Element) Set {
	// maybe optimize?
	s := NewSet()

	for i := range elements {
		s.Add(elements[i])
	}

	return s
}

func Equivalent(s1, s2 Set) bool {
	for _, e := range s1.Elements() {
		if !s2.Contains(e) {
			return false
		}
	}

	for _, e := range s2.Elements() {
		if !s1.Contains(e) {
			return false
		}
	}

	return true
}

// --- Set Implementation {{{

// set is guile's internal implementation of a set
type set struct {
	elements map[Element]bool
	count    uint
}

// Add will include Element, e, as a member of the set.
// If e is already a member of the set Add still works.
// Add returns a boolean. If the Element, e, was already
// a member of the Set Add returns true, else it is false
func (s *set) Add(e Element) bool {
	contains := s.Contains(e)
	if !contains {
		s.elements[e] = true
		s.count += 1
	}
	return contains
}

// Remove will exclude an Element, e, as a member of the set.
// If e is not a member of the set, Remove still works, but it
// will return false. If e was a member which was removed,
// Remove will return true.
func (s *set) Remove(e Element) bool {
	contains := s.Contains(e)
	if contains {
		delete(s.elements, e)
		s.count -= 1
	}
	return contains
}

// Contains returns a flag determining whether an Element, e
// is a member of the set
func (s *set) Contains(e Element) bool {
	_, ok := s.elements[e]
	return ok
}

// Cardinality returns the size of the set.
// Suppose set S, Cardinality(S) ≡ |S| (as expected)
func (s *set) Cardinality() uint {
	return s.count
}

// Elements returns a slice of the elements contained in this
// set. This slice is not the internal reprentation and therefore
// can be mutated.
func (s *set) Elements() []Element {
	e := make([]Element, 0)
	for k, _ := range s.elements {
		e = append(e, k)
	}
	return e
}

// --- }}}
