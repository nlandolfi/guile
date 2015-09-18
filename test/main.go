package main

import (
	"log"

	"github.com/nlandolfi/guile"
)

func main() {
	s := guile.NewSet()

	s.Add("one")
	s.Add("two")
	s.Add("three")

	b := guile.NewPhysicalBinaryRelationOn(s)

	b.AddRelation("one", "two")

	ub := guile.NewUtilityBinaryRelationOn(s, u)

	log.Print(guile.Reflexive(ub))
	log.Print(guile.Complete(ub))
	log.Print(guile.Transitive(ub))

	log.Print(guile.Rational(guile.Preference(ub)))
}

func u(a guile.Alternative) float64 {
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
