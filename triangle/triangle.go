// Package triangle contains a function to check what type of triangle
package triangle

// type of triangle
type Kind string

const (
	NaT = "NaT"
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
)

// KindFromSides says what triangle something is by it's length of sides
func KindFromSides(a, b, c float64) Kind {
	switch {
	case a <= 0 || b <= 0 || c <= 0 || a >= b+c || b >= c+a || c >= b+a:
		return NaT
	case a+b+c != a+b+c:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b && c != a || b == c && a != b || a == c && b != a:
		return Iso
	case a != b && b != c:
		return Sca
	default:
		return NaT
	}
}
