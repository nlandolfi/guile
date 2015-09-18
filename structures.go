package guile

// --- Economic Interpretation {{{

type (
	Alternative Element

	Alternatives Set

	Preference BinaryRelation

	UtilityFunction func(Alternative) float64
)

// --- }}

// -- Preference Implementation {{{

func Rational(p Preference) bool {
	return WeakOrder(p)
}

// --- }}}

// --- Utility Binary Relation Implemenation {{{

func NewUtilityBinaryRelationOn(universe Set, fn UtilityFunction) BinaryRelation {
	return &utilityRelation{
		Alternatives: universe,
		utility:      fn,
	}
}

type utilityRelation struct {
	Alternatives
	utility UtilityFunction
}

func (ur *utilityRelation) Universe() Set {
	return ur.Alternatives
}

func (r *utilityRelation) ContainsRelation(e1 Element, e2 Element) bool {
	return r.utility(e1) >= r.utility(e2)
}

// --- }}}
