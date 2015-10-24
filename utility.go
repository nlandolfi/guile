package guile

import "github.com/nlandolfi/set"

type (
	Utility float64

	UtilityFunction func(Alternative) Utility
)

// --- Utility Binary Relation Implemenation {{{

func NewUtilityBinaryRelationOn(universe set.Interface, fn UtilityFunction) set.BinaryRelation {
	return &utilityRelation{
		Alternatives: universe,
		utility:      fn,
	}
}

type utilityRelation struct {
	Alternatives
	utility UtilityFunction
}

func (ur *utilityRelation) Universe() set.Interface {
	return ur.Alternatives
}

func (r *utilityRelation) ContainsRelation(e1, e2 set.Element) bool {
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
