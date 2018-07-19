//Package arrayutil contains functions with arrays
package arrayutil

import (
	"fmt"
	"math"
)

//MinElement returns copy of min element from array
func MinElement(arr []int) int {
	min := math.MaxInt32
	for _, e := range arr {
		if min > e {
			min = e
		}
	}
	return min
}

//Sum returns sum of array elements
func Sum(arr []int) int {
	sum := 0
	for _, e := range arr {
		sum += e
	}
	return sum
}

//Print prints elements of the array on the screen
func PrintA(arr []int) {
	fmt.Print("[")
	for _, e := range arr {
		fmt.Printf("%v ", e)
	}
	fmt.Println("]")
}
