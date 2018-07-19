package arrayutil

//Sort sort an array using quicksort algorithm
func Sort(arr []int) []int {
	piv := len(arr) - 1
	quickSort(0, piv, arr)
	return arr
}

func quickSort(a, piv int, arr [] int) int {
	b := a

	for i := a; i < piv; i++ {
		if arr[i] < arr[piv] {
			if i != b {
				arr[i], arr[b] = arr[b], arr[i]
			}
			b++
		}
	}

	arr[b], arr[piv] = arr[piv], arr[b]

	if b-1 > a {
		quickSort(a, b-1, arr)
	}
	if b+1 < piv {
		quickSort(b+1, piv, arr)
	}

	return b
}

// Reverse reverses an array elements in place
func Reverse(arr []int) [] int {
	mid := len(arr) / 2
	for i, j := 0, len(arr)-1; i < mid; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
