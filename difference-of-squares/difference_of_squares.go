// Package diffsquares provides functions to perform calculations on the first N natural numbers for a given n.
package diffsquares

// SquareOfSum sums the first n elements (from 1 to n) and returns the result to the power of 2.
func SquareOfSum(n int) int {
	result := (n * (n + 1)) / 2
	return result * result
}

// SumOfSquares calculates the power of two for the first n elements (from 1 to n) and then sums them all.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between SquareOfSum and SumOfSquares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
