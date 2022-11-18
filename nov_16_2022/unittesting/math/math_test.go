package main

import (
	"testing"
)

/* SIMPLE TEST */
func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

// func TestSubtract(t *testing.T) {
// 	got := Subtract(10, 5)
// 	want := 5

// 	if got != want {
// 		t.Errorf("got %d, wanted %d", got, want)
// 	}
// }

/* TABLE DRIVEN TEST */
// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'

// type addTest struct {
// 	arg1, arg2, expected int
// }

// type subtractTest struct {
// 	arg1, arg2, expected int
// }

// var addTests = []addTest{
// 	addTest{2, 3, 5},
// 	addTest{4, 8, 12},
// 	addTest{6, 9, 15},
// 	addTest{3, 10, 13},
// }

// var subtractTests = []subtractTest{
// 	subtractTest{10, 2, 8},
// 	subtractTest{5, 2, 3},
// 	subtractTest{8, 1, 7},
// 	subtractTest{9, 6, 3},
// }

// func TestAdd(t *testing.T) {

// 	for _, test := range addTests {
// 		if output := Add(test.arg1, test.arg2); output != test.expected {
// 			t.Errorf("Output %d not equal to expected %d", output, test.expected)
// 		}
// 	}
// }

// func TestSubtract(t *testing.T) {

// 	for _, test := range subtractTests {
// 		if output := Subtract(test.arg1, test.arg2); output != test.expected {
// 			t.Errorf("Output %d not equal to expected %d", output, test.expected)
// 		}
// 	}
// }

/*BENCHMARK TESTING*/
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

// func ExampleAdd() {
// 	fmt.Println(Add(4, 6))
// 	// Output: 10
// }
