//Package gomath contains GCD & LCM mathematical functions
package gomath

//GCD returns greatest common divisor of two integer numbers
func GCD(a, b int) int {
	var k int
	for b != 0 {
		k = a % b
		a = b
		b = k
	}
	return a
}

//LCM returns least common multiple of two integer numbers
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
