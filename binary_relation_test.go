package guile

import (
	"testing"

	"github.com/nlandolfi/set"
)

func u(a Alternative) Utility {
	switch a {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	default:
		return 0
	}
}

func TestUtilityBinaryRelation(t *testing.T) {
	s := set.New()

	s.Add("one")
	s.Add("two")
	s.Add("three")

	ub := NewUtilityBinaryRelationOn(s, u)

	if !set.Reflexive(ub) {
		t.Errorf("Our UtilityBinaryRelation should be reflexive")
	}

	if !set.Complete(ub) {
		t.Errorf("Our UtilityBinaryRelation should be complete")
	}

	if !set.Transitive(ub) {
		t.Errorf("Our UtilityBinaryRelation should be transitive")
	}

	if !Rational(Preference(ub)) {
		t.Errorf("Our UtilityBinaryRelation should be rational! -- von Neumann-Morgenstern")
	}
}
