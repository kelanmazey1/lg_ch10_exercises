// Package for the exercises from chapter 10 of Learing Go.
package lgch10exercises

import "golang.org/x/exp/constraints"

// Type to group Integer and Float [https://golang.org/x/exp/constraints]
type Number interface {
	constraints.Integer | constraints.Float
}

// Adds  2 integers together and returns result
//
// See: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
