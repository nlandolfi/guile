package guile

import "testing"

func PreferenceOverLotteryExample(t *testing.T) {
	s := NewSetWithElements([]Element{0, 1, 2, 3, 4})

	l := NewUniformLottery(s)

	l2 := NewLottery(s)
	l.AddOutcome(2, .5)
	l.AddOutcome(3, .5)

	l3 := NewDegenerateLottery(s, 4)

	// E(l)  => .2*0 + .2*10 + .2*20 + .2*30 + .2*30 = 18
	// E(l2) => .5*20 + .5*30 = 35
	// E(l3) => 1*40 = 40

	lotteries := NewSetWithElements([]Element{l, l2, l3})

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
