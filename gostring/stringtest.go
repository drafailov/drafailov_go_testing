//Package gostring contains function which sets the length ot random array of strings
package gostring

import (
	"math/rand"
)

// Str returns random strings with set length
func Str(leng int) [] string {
	var arr [94] string
	for i := 0; i < 94; i++ {
		arr[i] = string(i + 32)
	}
	b := make([]string, leng)
	for i := range b {
		b[i] = arr[rand.Intn(len(arr))]
	}
	return b
}
