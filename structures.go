package guile

// Math Econ Structures
type (
	Alternative Element

	Alternatives Set

	Preference interface {
		BinaryRelation
		Rational() bool
	}
)

// -- Preference Implementation {{{

type preference binaryRelation

func (p *preference) Rational() bool {
	return (*binaryRelation)(p).WeakOrder()
}

// --- }}}
