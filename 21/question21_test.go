package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	ansTest := 16
	rTest := PartOne("test.txt", 6)
	if rTest != ansTest {
		t.Errorf("Got %d; want %d", rTest, ansTest)
	}
	
	// ansFull := 406849
	// rFull := PartOne("data.txt")
	// if rFull != ansFull {
	// 	t.Errorf("Got %d; want %d", rFull, ansFull)
	// }
}

func TestPartTwo(t *testing.T) {
	// var ansTest1 int64 = 167409079868000
	// rTest1 := PartTwo("test.txt")
	// if rTest1 != ansTest1 {
	// 	t.Errorf("Got %d; want %d", rTest1, ansTest1)
	// }
	// var ansTest2 int64 = 0
	// rTest2 := PartTwo("test2.txt")
	// if rTest2 != ansTest2 {
	// 	t.Errorf("Got %d; want %d", rTest2, ansTest2)
	// }
	
	// ansFull := 247153
	// rFull := PartTwo("data.txt")
	// if rFull != ansFull {
	// 	t.Errorf("Got %d; want %d", rFull, ansFull)
	// }
}