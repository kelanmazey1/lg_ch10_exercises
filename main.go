// Package for the exercises from chapter 10 of Learing Go.
package lgch10exercises

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

//
// [Adds]: https://www.mathsisfun.com/numbers/addition.html 2 integers together and returns result
func Add[T Number](a, b T) T {
	return a + b
}
