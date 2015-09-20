package guile

// --- Economic Interpretation {{{

type (
	Alternative Element

	Alternatives Set

	Preference BinaryRelation
)

// --- }}

// -- Preference Implementation {{{

func Rational(p Preference) bool {
	return WeakOrder(p)
}

// --- }}}
