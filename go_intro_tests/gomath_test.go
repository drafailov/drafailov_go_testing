package go_intro_tests

import (
	"testing"

	"github.com/drafailov/dimitar_rafailov_go_testing/gomath"
)

func TestGCD(t *testing.T) {

	var GCDTest = []struct {
		firstNumber  int
		secondNumber int
		expectedGCD  int
	}{
		{0, 5, 5},
		{1, 0, 1},
		{1, 1, 1},
		{54, 24, 6},
		{1024, 256, 256},
		{235819, 16066466, 23},
		{134324237, 1242342342, 1},
	}

	for _, e := range GCDTest {
		receivedGCD := gomath.GCD(e.firstNumber, e.secondNumber)
		if receivedGCD != e.expectedGCD {
			t.Errorf("GCD(%v, %v): receivedGCD (%v); expectedGCD (%v)",
				e.firstNumber, e.secondNumber, receivedGCD, e.expectedGCD)
		}
	}
}

func TestLCM(t *testing.T) {

	var LCMTest = []struct {
		firstNumber  int
		secondNumber int
		expectedLCM  int
	}{
		{1, 1, 1},
		{4, 6, 12},
		{12, 18, 36},
		{327, 12, 1308},
		{6986, 17905118, 17905118},
		{638128, 62262, 19865562768},
	}

	for _, e := range LCMTest {
		receivedLCM := gomath.LCM(e.firstNumber, e.secondNumber)
		if receivedLCM != e.expectedLCM {
			t.Errorf("LCM(%v, %v): received = %v; expected %v",
				e.firstNumber, e.secondNumber, receivedLCM, e.expectedLCM)
		}
	}
}
