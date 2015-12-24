package main

import (
	"log"

	"github.com/nlandolfi/guile"
	"github.com/nlandolfi/set"
	"github.com/nlandolfi/set/relation"
)

func main() {
	s := set.With([]set.Element{"pushups", "reading", "sleep"})

	health := relation.New(s)
	health.AddRelation("sleep", "pushups")
	health.AddRelation("pushups", "reading")

	philosophy := relation.New(s)
	philosophy.AddRelation("reading", "pushups")
	philosophy.AddRelation("pushups", "sleep")

	humanity := relation.New(s)
	humanity.AddRelation("reading", "sleep")
	humanity.AddRelation("sleep", "pushups")

	elos := guile.PreferenceProfile{health, philosophy, humanity}

	r := guile.PairwiseMajority(elos)
	log.Print(guile.MostPreferred(r))

	b := guile.BordaCounting(elos)
	log.Print(guile.MostPreferred(b))
}
