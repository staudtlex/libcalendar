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
// The following Go code is translated from the Lisp code discussed in:
// - Dershowitz, Nachum, and Edward Reingold. 1990. "Calendrical
//   Calculations", Software - Practice and Experience, 20 (9), 899-928.
//   https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.17.4274
// - Reingold, Edward, Nachum Dershowitz, and Stewart Clamen. 1993.
//   "Calendrical Calculations, II: Three Historical Calendars",
//   Software - Practice & Experience, 23 (4), 383-404.
//   https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.13.9215

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

// ratio creates a rational number from its arguemnts x. If no argument x
// is passed, `ratio` returns a *big.Rat whose value is set to zero. Called with
// a single argument, `ratio` returns a *big.Rat whose value is set to the
// argument's value. Otherwise, `ratio` returns a *big.Rat with its numerator set
// to the first argument's value, and the denominator set to the second. Any
// additional argument is ignored. If the second argument's value == 0, ratio
// panics.
func ratio(x float64, y float64) *big.Rat {
	return big.NewRat(int64(x), int64(y))
}

// add computes the sum of the rational numbers specified in x.
func add(x ...*big.Rat) *big.Rat {
	sum := &big.Rat{}
	for i := range x {
		sum.Add(sum, x[i])
	}
	return sum
}

// sub computes the difference of the rational numbers x and y.
func sub(x, y *big.Rat) *big.Rat {
	return (&big.Rat{}).Sub(x, y)
}

// mult returns the product of the rational numbers specified in x.
func mult(x ...*big.Rat) *big.Rat {
	prod := (&big.Rat{}).SetInt64(1)
	for i := range x {
		prod.Mul(prod, x[i])
	}
	return prod
}

// div returns the quotient of two rational numbers x and y.
func div(x, y *big.Rat) *big.Rat {
	return (&big.Rat{}).Quo(x, y)
}

// quotient computes the quotient of two rational numbers, truncated
// towards zero.
func quotient(x, y *big.Rat) float64 {
	result, _ := div(x, y).Float64()
	return math.Floor(result)
}

// floorf returns the greatest integer less than or equal to x. Note that the
// returned value is a float64.
func floorf(x *big.Rat) float64 {
	val, _ := x.Float64()
	return math.Floor(val)
}

// floorf returns the greatest integer less than or equal to x.
func floor(x *big.Rat) *big.Rat {
	return big.NewRat(int64(floorf(x)), 1)
}

// modr computes the positive (rational) remainder of a mod b.
func modr(a *big.Rat, b *big.Rat) *big.Rat {
	r := sub(a, mult(b, floor(div(a, b))))
	if r.Cmp(big.NewRat(0, 1)) < 0 {
		return add(r, b)
	} else {
		return r
	}
}
