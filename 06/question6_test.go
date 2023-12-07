package main

import (
	"testing"
	"time"
)

func TestPartOne(t *testing.T){
	ans := 288
	r := PartOne("test.txt")
	if r != ans{
		t.Errorf("Got %d; want %d", r, ans)
	} 
}

func TestPartTwo(t *testing.T){
	ans := 71503
	r := PartTwo("test.txt")
	if r != ans{
		t.Errorf("Got %d; want %d", r, ans)
	}
}

func TestRunAll(t *testing.T) {
	start1 := time.Now()
	p1 := PartOne("data.txt")
	elapsed1 := time.Since(start1)
	start2 := time.Now()
	p2 := PartTwo("data.txt")
	elapsed2 := time.Since(start2)
	t.Logf("ANSWER PartOne: %d....took %s\n", p1, elapsed1) //4568778
	t.Logf("ANSWER PartTwo: %d....took %s\n", p2, elapsed2) //28973936
}