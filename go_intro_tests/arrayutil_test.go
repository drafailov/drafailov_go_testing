package go_intro_tests

import (
	"testing"

	"github.com/drafailov/dimitar_rafailov_go_testing/arrayutil"
)

func TestMinElement(t *testing.T) {

	var MinElementTest = []struct {
		in       []  int
		expected int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6}, 0},
		{[]int{1, 2, 3, 4, 5, 6, -6}, -6},
		{[]int{10, 10, 10, 10, 10}, 10},
		{[]int{-1, -2, -3, -4, -5, -6, -7, -190, 0}, -190},
		{[]int{14440, 23232, 12123, 31234, 177235, 99996, 809996, 42309992,}, 12123},
	}

	for _, e := range MinElementTest {
		received := arrayutil.MinElement(e.in)
		if received != e.expected {
			t.Errorf("MinElement %v: received = %v; expected %v",
				e.in, received, e.expected)
		}
	}
}

func TestSum(t *testing.T) {

	var SumTest = []struct {
		in       []  int
		expected int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6}, 21},
		{[]int{1, 2, 3, 4, 5, 6, -6}, 15},
		{[]int{10, 10, 10, 10, 10}, 50},
		{[]int{-1, -2, -3, -4, -5, -6, -7, -190, 0}, -218},
		{[]int{14440, 23232, 12123, 31234, 177235, 99996, 809996, 42309992}, 43478248},
	}

	for _, e := range SumTest {
		received := arrayutil.Sum(e.in)
		if received != e.expected {
			t.Errorf("Sum %v: received = %v; expected %v",
				e.in, received, e.expected)
		}
	}
}

func TestSort(t *testing.T) {

	var SortTest = []struct {
		in       []  int
		expected [] int
	}{
		{[]int{0}, []int{0}},
		{[]int{0, 1, 2, 3, 4, 5, 6}, []int{0, 1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6, -6}, []int{-6, 1, 2, 3, 4, 5, 6}},
		{[]int{10, 10, 10, 10, 10}, []int{10, 10, 10, 10, 10}},
		{[]int{-1, -2, -3, -4, -5, -6, -7, -190, 0}, []int{-190, -7, -6, -5, -4, -3, -2, -1, 0}},
		{[]int{14440, 23232, 12123, 31234, 177235, 99996, 809996, 42309992}, []int{12123, 14440, 23232, 31234, 99996, 177235, 809996, 42309992}},
	}

	for _, e := range SortTest {
		received := arrayutil.Sort(e.in)
		if len(received) != len(e.expected) {
			t.Errorf("received = %v; expected %v",
				received, e.expected)
		}
		for i := range e.expected {
			if received[i] != e.expected[i] {
				t.Fatalf("received = %v; expected %v",
					received, e.expected)
			}

		}
	}
}

func TestReverse(t *testing.T) {

	var ReverseTest = []struct {
		in       []  int
		expected [] int
	}{
		{[]int{0}, []int{0}},
		{[]int{0, 1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1, 0}},
		{[]int{1, 2, 3, 4, 5, 6, -6}, []int{-6, 6, 5, 4, 3, 2, 1}},
		{[]int{10, 10, 10, 10, 10}, []int{10, 10, 10, 10, 10}},
		{[]int{-1, -2, -3, -4, -5, -6, -7, -190, 0}, []int{0, -190, -7, -6, -5, -4, -3, -2, -1}},
		{[]int{14440, 23232, 12123, 31234, 177235, 99996, 809996, 42309992}, []int{42309992, 809996, 99996, 177235, 31234, 12123, 23232, 14440}},
	}

	for _, e := range ReverseTest {
		received := arrayutil.Reverse(e.in)
		if len(received) != len(e.expected) {
			t.Errorf("received: %v; expected: %v",
				received, e.expected)
		}
		for i := range e.expected {
			if received[i] != e.expected[i] {
				t.Fatalf("received: %v; expected: %v",
					received, e.expected)
			}

		}
	}
}
