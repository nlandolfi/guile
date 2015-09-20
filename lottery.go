package guile

import "math/rand"

// --- Probability (Modeling Uncertainty) {{{
type Probability float64

func (p Probability) Valid() bool {
	return p >= 0 && p <= 1
}

const (
	Impossible Probability = 0
	Certain    Probability = 1
)

// --- }}}

type Lottery interface {
	Alternatives() Alternatives
	Outcomes() []Alternative
	AddOutcome(Alternative, Probability)
	ProbabilityOf(Alternative) Probability
}

// --- Lottery Constructors {{{

func NewLottery(a Alternatives) Lottery {
	return &lottery{
		alternatives: a,
		support:      make(map[Alternative]Probability),
	}
}

func NewUniformLottery(alternatives Alternatives) Lottery {
	l := NewLottery(alternatives)

	count := alternatives.Cardinality()
	p := Probability(1.0 / count)

	for a := range alternatives.Elements() {
		l.AddOutcome(a, p)
	}

	return l
}

func NewDegenerateLottery(alternatives Alternatives, a Alternative) Lottery {
	assert(alternatives.Contains(a), "set of alternatives doesn't contain degenerate alternatives")

	l := NewLottery(alternatives)

	l.AddOutcome(a, Certain)

	return l
}

/// --- }}}

// --- Lottery Implementation {{{

type lottery struct {
	alternatives Alternatives
	support      map[Alternative]Probability
}

func (l *lottery) Alternatives() Alternatives {
	return l.alternatives
}

func (l *lottery) Outcomes() []Alternative {
	as := make([]Alternative, 0)

	for a, _ := range l.support {
		as = append(as, a)
	}

	return as
}

func (l *lottery) AddOutcome(a Alternative, p Probability) {
	assert(Support(l) < 1, "lottery already fully supported")
	assert(Support(l)+p <= 1, "adding outcome would over-support")
	assert(p.Valid(), "invalid probability")

	l.support[a] = p
}

func (l *lottery) ProbabilityOf(a Alternative) Probability {
	p, ok := l.support[a]

	if ok {
		return p
	} else {
		return 0
	}
}

// --- }}}

// --- Properties {{{

func Cardinality(l Lottery) uint {
	return uint(len(l.Outcomes()))
}

func Support(l Lottery) Probability {
	p := Probability(0)

	for _, o := range l.Outcomes() {
		p += l.ProbabilityOf(o)
	}

	return p
}

func FullySupported(l Lottery) bool {
	return Support(l) == Probability(1)
}

func Degenerate(l Lottery) bool {
	return Cardinality(l) == 1 && l.ProbabilityOf(l.Outcomes()[0]) == Probability(1)
}

// --- }}}

// --- Composition {{{

func Compose(p, q Lottery, alpha Probability) Lottery {
	assert(FullySupported(p), "first lottery not fully supported")
	assert(FullySupported(q), "second lottery not fully supported")
	assert(Equivalent(p.Alternatives(), q.Alternatives()), "lottery alternatives must be equivalent")

	n := NewLottery(p.Alternatives())

	for a := range n.Alternatives().Elements() {
		p := alpha*p.ProbabilityOf(a) + (1-alpha)*q.ProbabilityOf(a)
		if p == Probability(0) {
			continue
		}

		n.AddOutcome(a, p)
	}

	return n
}

// --- }}}

// --- Simulation {{{

func Simulate(l Lottery) Alternative {
	assert(FullySupported(l), "lottery not fully supported")

	f := Probability(rand.Float64())
	p := Probability(0)

	var last Alternative
	for _, o := range l.Outcomes() {
		p += l.ProbabilityOf(o)
		last = o

		if f < p {
			return o
		}
	}

	return last
}

// --- }}}
