package guile

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

func NewSet() Set {
	return &set{
		elements: make(map[Element]bool),
		count:    0,
	}
}

// --- Set Implementation {{{

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

func assert(flag bool, s string) {
	if !flag {
		panic(s)
	}
}

// --- }}}
