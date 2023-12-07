package main

import (
	"testing"
)

func TestPartOne(t *testing.T){
	ans := 4361
	r := PartOne("test.txt")
	if r != ans{
		t.Errorf("Got %d; want %d", r, ans)
	} 
}

func TestPartTwo(t *testing.T){
	ans := 0
	r := PartTwo("test.txt")
	if r != ans{
		t.Errorf("Got %d; want %d", r, ans)
	}
}

func TestRunAll(t *testing.T) {
	t.Logf("ANSWER PartOne: %d\n", PartOne("data.txt")) //52974
	t.Logf("ANSWER PartTwo: %d\n", PartTwo("data.txt")) //53340
}