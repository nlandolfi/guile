package main

import (
	"log"

	"github.com/nlandolfi/guile"
)

func main() {
	s := guile.NewSetWithElements([]guile.Element{"pushups", "reading", "sleep"})

	health := guile.NewPhysicalBinaryRelationOn(s)
	health.AddRelation("sleep", "pushups")
	health.AddRelation("pushups", "reading")

	philosophy := guile.NewPhysicalBinaryRelationOn(s)
	philosophy.AddRelation("reading", "pushups")
	philosophy.AddRelation("pushups", "sleep")

	humanity := guile.NewPhysicalBinaryRelationOn(s)
	humanity.AddRelation("reading", "sleep")
	humanity.AddRelation("sleep", "pushups")

	elos := guile.PreferenceProfile{health, philosophy, humanity}

	r := guile.PairwiseMajority(elos)
	log.Print(guile.MostPreferred(r))

	b := guile.BordaCounting(elos)
	log.Print(guile.MostPreferred(b))
}
