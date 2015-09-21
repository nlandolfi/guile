package guile

// --- Economic Interpretation {{{

type (
	Alternative Element

	Alternatives Set

	Preference BinaryRelation

	PreferenceProfile []Preference

	SocialWelfareFunction func(PreferenceProfile) Preference
)

// --- }}}

// -- Preference Implementation {{{

func Rational(p Preference) bool {
	return WeakOrder(p)
}

func ComposablePreferences(prefs []Preference) bool {
	br := make([]BinaryRelation, len(prefs))

	for i := range prefs {
		br[i] = BinaryRelation(prefs[i])
	}

	return ComposableRelations(br)
}

func ProfileUniverse(p PreferenceProfile) Alternatives {
	assert(len(p) > 0, "preference profile empty")
	assert(ComposablePreferences(p), "preferneces not defined over same universe")

	return p[0].Universe()
}

// CountPreferenceOf returns the number of people who prefer x to y in the preference profile pp
func CountPreferenceOf(x, y Alternative, pp PreferenceProfile) float64 {
	return CountWeightedPreferenceOf(x, y, pp, func(i uint) float64 { return 1 })
}

// CountWeightedPreferenceOf returns the weigth of people who prefer x, y in preference profile pp
func CountWeightedPreferenceOf(x, y Alternative, pp PreferenceProfile, w func(uint) float64) float64 {
	c := 0.0

	for i, p := range pp {
		if p.ContainsRelation(x, y) {
			c += w(uint(i))
		}
	}

	return c
}

func BordaCount(x Alternative, p Preference) uint {
	c := uint(0)

	for e := range p.Universe().Elements() {
		if p.ContainsRelation(x, e) {
			c += 1
		}
	}

	return c
}

func ProfileBordaCount(x Alternative, pp PreferenceProfile) uint {
	c := uint(0)

	for _, p := range pp {
		c += BordaCount(x, p)
	}

	return c
}

// --- }}}

// --- Social Welfare Function Implementations {{{

func PairwiseMajority(pp PreferenceProfile) Preference {
	u := ProfileUniverse(pp)

	return NewFunctionBinaryRelation(u, func(x, y Element) bool {
		return CountPreferenceOf(x, y, pp) >= CountPreferenceOf(y, x, pp)
	})
}

func BordaCounting(pp PreferenceProfile) Preference {
	u := ProfileUniverse(pp)

	return NewFunctionBinaryRelation(u, func(x, y Element) bool {
		return ProfileBordaCount(x, pp) >= ProfileBordaCount(y, pp)
	})
}

func Dictatorship(pp PreferenceProfile, individual uint) Preference {
	return pp[individual]
}

func AntiDictatorship(pp PreferenceProfile, individual uint) Preference {
	return Reverse(pp[individual])
}

func Constant(pp PreferenceProfile, actual Preference) Preference {
	// interesting how explicit it become that a constant one depends not on the PreferenceProfile
	return actual
}

func WeightedMajority(pp PreferenceProfile, w func(uint) float64) Preference {
	u := ProfileUniverse(pp)

	return NewFunctionBinaryRelation(u, func(x, y Element) bool {
		return CountWeightedPreferenceOf(x, y, pp, w) >= CountWeightedPreferenceOf(y, x, pp, w)
	})
}

// --- }}}
