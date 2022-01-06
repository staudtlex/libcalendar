// Copyright (C) 2021  Alexander Staudt
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// The Go code for the functions "amod" and "sum" is translated from the
// Lisp code discussed in:
// - Dershowitz, Nachum, and Edward Reingold. 1990. "Calendrical
//   Calculations", Software---Practice and Experience, 20 (9), 899--928.
// - Reingold, Edward, Nachum Dershowitz, and Stewart Clamen. 1993.
//   "Calendrical Calculations, II: Three Historical Calendars",
//   Software---Practice & Experience, 23 (4), 383--404.

package libcalendar

import (
	"math"
	"math/big"
)

// mod computes the positive remainder of a mod b
func mod(a float64, b float64) float64 {
	r := a - b*math.Floor(a/b)
	if r < 0 {
		r = r + b
	}
	return r
}

// amod computes the adjusted positive remainder Mod(a-1, b) + 1
func amod(a float64, b float64) float64 {
	return mod(a-1, b) + 1
}

// sum computes the sum \Sigma_{i\ge k, p(i)} f(i), as long as the condition
// p(i) is true.
func sum(f func(float64) float64, k float64, p func(float64) bool) float64 {
	sum := 0.0
	for i := k; p(i); i++ {
		sum += f(i)
	}
	return sum
}

// member returns true if x is an element of slice s, and false otherwise.
func member(x float64, s []float64) bool {
	for i := 0; i < len(s); i++ {
		if x == s[i] {
			return true
		}
	}
	return false
}

// Convenience functions for working with *big.Rat

// Ratio creates a rational number from its arguemnts x. If no argument x
// is passed, `Ratio` returns a *big.Rat whose value is set to zero. Called with
// a single argument, `Ratio` returns a *big.Rat whose value is set to the
// argument's value. Otherwise, `Ratio` returns a *big.Rat with its numerator set
// to the first argument's value, and the denominator set to the second. Any
// additional argument is ignored. If the second argument's value == 0, Ratio
// panics.
func Ratio(x ...float64) *big.Rat {
	switch len(x) {
	case 0:
		return (&big.Rat{}).SetInt64(0)
	case 1:
		return (&big.Rat{}).SetInt64(int64(x[0]))
	default:
		return big.NewRat(int64(x[0]), int64(x[1]))
	}
}

// Add computes the sum of the rational numbers specified in x.
func Add(x ...*big.Rat) *big.Rat {
	sum := &big.Rat{}
	for i := range x {
		sum.Add(sum, x[i])
	}
	return sum
}

// Sub computes the difference of the rational numbers x and y.
func Sub(x, y *big.Rat) *big.Rat {
	return (&big.Rat{}).Sub(x, y)
}

// Mult returns the product of the rational numbers specified in x.
func Mult(x ...*big.Rat) *big.Rat {
	prod := (&big.Rat{}).SetInt64(1)
	for i := range x {
		prod.Mul(prod, x[i])
	}
	return prod
}

// Div returns the quotient of two rational numbers x and y.
func Div(x, y *big.Rat) *big.Rat {
	return (&big.Rat{}).Quo(x, y)
}

// Quotient computes the Quotient of two rational numbers, truncated
// towards zero.
func Quotient(x, y *big.Rat) float64 {
	result, _ := Div(x, y).Float64()
	return math.Floor(result)
}

// Floorf returns the greatest integer less than or equal to x. Note that the
// returned value is a float64.
func Floorf(x *big.Rat) float64 {
	val, _ := x.Float64()
	return math.Floor(val)
}

// floorf returns the greatest integer less than or equal to x.
func Floor(x *big.Rat) *big.Rat {
	return Ratio(Floorf(x))
}

// Modr computes the positive (rational) remainder of a mod b.
func Modr(a *big.Rat, b *big.Rat) *big.Rat {
	r := Sub(a, Mult(b, Floor(Div(a, b))))
	if r.Cmp(Ratio(0)) < 0 {
		return Add(r, b)
	} else {
		return r
	}
}
