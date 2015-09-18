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

	b := guile.NewBinaryRelationOn(s)

	b.AddRelation("one", "two", guile.Left)
	log.Print(b.ContainsRelation("one", "two"))
	log.Print(b.ContainsRelation("two", "one"))
}
