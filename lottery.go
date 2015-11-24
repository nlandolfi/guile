package guile

import "github.com/nlandolfi/prob"

type Lottery prob.DiscreteDistribution

// --- Lottery Constructors {{{

func NewLottery(a Alternatives) Lottery {
	return prob.NewDiscreteDistribution(a)
}

func NewUniformLottery(alternatives Alternatives) Lottery {
	return prob.NewUniformDiscrete(alternatives)
}

func NewDegenerateLottery(alternatives Alternatives, a Alternative) Lottery {
	assert(alternatives.Contains(a), "set of alternatives doesn't contain degenerate alternatives")

	l := NewLottery(alternatives)

	l.AddOutcome(a, prob.Certain)

	return l
}

/// --- }}}

// --- Lottery Implementation {{{

// assert is a helper function to provide
// moderate runtime type checking on the Element interface
func assert(flag bool, s string) {
	if !flag {
		panic(s)
	}
}

// --- }}}

// --- Properties {{{

func Cardinality(l Lottery) uint {
	return prob.Cardinality(l)
}

func Support(l Lottery) prob.Probability {
	return prob.Support(l)
}

func FullySupported(l Lottery) bool {
	return prob.FullySupported(l)
}

func Degenerate(l Lottery) bool {
	return prob.Degenerate(l)
}

// --- }}}

// --- Composition {{{

func Compose(p, q Lottery, alpha prob.Probability) Lottery {
	return prob.Compose(p, q, alpha)
}

// --- }}}

// --- Simulation {{{

func Simulate(l Lottery) Alternative {
	return prob.Simulate(l)
}

// --- }}}
