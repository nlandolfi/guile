package guile

import (
	"math"
	"testing"

	"github.com/nlandolfi/set"
)

var epsilon = 0.000001

func equiv(f1, f2 float64) bool {
	return math.Abs(f1-f2) < epsilon
}

func PreferenceOverLotteryExample(t *testing.T) {
	s := set.WithElements(0, 1, 2, 3, 4)

	l := NewUniformLottery(s)

	l2 := NewLottery(s)
	l.AddOutcome(2, .5)
	l.AddOutcome(3, .5)

	l3 := NewDegenerateLottery(s, 4)

	// E(l)  => .2*0 + .2*10 + .2*20 + .2*30 + .2*30 = 18
	// E(l2) => .5*20 + .5*30 = 35
	// E(l3) => 1*40 = 40

	lotteries := set.WithElements(l, l2, l3)

	b := NewUtilityBinaryRelationOn(lotteries, func(a Alternative) Utility {
		return ExpectedUtility(a.(Lottery), func(a Alternative) Utility {
			return Utility(a.(int) * 10)
		})
	})

	if !Rational(b) {
		t.Errorf("expected the relation to be rational")
	}

	if !b.ContainsRelation(l3, l2) {
		t.Errorf("expected to prefer l3 to l2")
	}

	if !b.ContainsRelation(l2, l) {
		t.Errorf("expected to prefer l2 to l1")
	}
}

func TestLotteryUtility(t *testing.T) {
	l := NewLottery(set.New())

	l.AddOutcome("A", 0.20)
	l.AddOutcome("B", 0.30)
	l.AddOutcome("C", 0.40)
	l.AddOutcome("D", 0.10)

	utilityFn := func(a Alternative) Utility {
		switch a {
		case "A":
			return 1
		case "B":
			return 3
		case "C":
			return 5
		case "D":
			return 2
		}
		panic("not recognized")
	}

	if f := ExpectedUtility(l, utilityFn); !equiv(float64(f), 3.3) {
		t.Fatalf("Expected utility of 3.3, but got: %f", f)
	}
}

func TestLotteryOfLotteryUtility(t *testing.T) {
	l2 := NewUniformLottery(set.WithElements("X", "Y"))
	l1 := NewUniformLottery(set.WithElements("A", l2))

	utilityFn := func(a Alternative) Utility {
		switch a {
		case "A":
			return 4
		default:
			return ExpectedUtility(a.(Lottery), func(a Alternative) Utility {
				switch a {
				case "X":
					return 4
				case "Y":
					return 8
				default:
					panic("unknown")
				}
			})
		}
	}

	if f := ExpectedUtility(l1, utilityFn); !equiv(float64(f), 5.0) {
		t.Fatalf("Expected utility of 4.5, but got: %f", f)
	}
}
