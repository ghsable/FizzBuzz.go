package main

import (
	"testing"
	"reflect"
)

func TestApp(t *testing.T) {
	// DATA: WANT
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "Fizz Buzz"}

	// DATA: GOT
	var got []string
	for n := 1; n <= 15; n++ {
		got = append(got, getPrint(n))
	}

	// CHECK
	if !reflect.DeepEqual(want, got) {
		t.Errorf("\n[ERROR] getPrint(n int):\n  want = %s\n  got  = %s", want, got)
	}
}
