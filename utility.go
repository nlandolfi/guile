package guile

type (
	Utility float64

	UtilityFunction func(Alternative) Utility
)

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

// --- ExpectedUtility over a lottery {{{

func ExpectedUtility(l Lottery, utility UtilityFunction) Utility {
	u := Utility(0)
	for o := range l.Outcomes() {
		u += utility(o) * Utility(l.ProbabilityOf(o))
	}
	return u
}

// --- }}}
