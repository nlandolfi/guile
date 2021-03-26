package main

import (
	"log"

	"github.com/nlandolfi/guile"
	"github.com/nlandolfi/set"
)

func id(a guile.Alternative) guile.Utility {
	return guile.Utility(a.(int))
}

func sq(a guile.Alternative) guile.Utility {
	return guile.Utility(a.(int)) * guile.Utility(a.(int))
}

func main() {
	l := guile.NewUniformLottery(set.WithElements(1, 4, 5))

	log.Print(guile.ExpectedUtility(l, id))
	log.Print(guile.ExpectedUtility(l, sq))

	l2 := guile.NewUniformLottery(set.WithElements(4))

	log.Print(guile.ExpectedUtility(l2, id))
	log.Print(guile.ExpectedUtility(l2, sq))

	l3 := guile.NewUniformLottery(set.WithElements(3, 4))

	log.Print(guile.ExpectedUtility(l3, id))
	log.Print(guile.ExpectedUtility(l3, sq))
}
