package main

import (
	"testing"
)

func TestApp(t *testing.T) {
	// DATA: WANT
	var want [15]string = [15]string{"0", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "Fizz Buzz"}

	// DATA: GOT
	var got [15]string
	for n := 1; n <= 15; n++ {
		got[n-1] = getPrint(n)
	}

	// CHECK
	if want != got {
		t.Errorf("\n[ERROR] getPrint(n int):\n  want = %s\n  got  = %s", want, got)
	}
}
