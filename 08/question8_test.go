package main

import (
	"testing"
	"time"
)

func TestPartOne(t *testing.T){
	ans := 2
	r := PartOne("test.txt")
	if r != ans{
		t.Errorf("Got %d; want %d", r, ans)
	} 
	ans2 := 6
	r2 := PartOne("test2.txt")
	if r2 != ans2{
		t.Errorf("Got %d; want %d", r2, ans2)
	} 
}

func TestPartTwo(t *testing.T){
	ans := 6
	r := PartTwo("test3.txt")
	if r != uint64(ans){
		t.Errorf("Got %d; want %d", r, ans)
	}
}

func TestRunAll(t *testing.T) {
	start1 := time.Now()
	p1 := PartOne("data.txt")
	elapsed1 := time.Since(start1)
	t.Logf("ANSWER PartOne: %d....took %s\n", p1, elapsed1) //13019
	start2 := time.Now()
	p2 := PartTwo("data.txt")
	elapsed2 := time.Since(start2)
	t.Logf("ANSWER PartTwo: %d....took %s\n", p2, elapsed2) //13524038372771
}