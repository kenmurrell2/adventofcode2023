package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Converter struct {
	dest []int
	src []int
	length []int
}

func (c *Converter) Convert(num int) int {
	for i, _ := range(c.dest) {
		if num >= c.src[i] && num < c.src[i] + c.length[i]{
			return c.dest[i] + num - c.src[i] 
		} 
	}

	return num
}

func ParseInput(path string, runtwo bool) ([]int, []Converter) { //stupid but whatever
	file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	var seeds []int
	var cons []Converter
	
	var t_dest []int
	var t_src []int
	var t_length []int

    for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "seeds:") {
			if runtwo {
				ParseSeeds2(text, &seeds)
			} else {
				ParseSeeds(text, &seeds)
			}
			continue
		}
		if strings.Contains(text, "-to-") {
			continue
		}
		if len(strings.Trim(text, "\n")) == 0{
			if len(t_dest) == 0 {
				continue
			}
			CreateAndReset(&cons, &t_dest, &t_src, &t_length)
		} else {
			numbers := strings.Split(text, " ")
			d, _ := strconv.Atoi(numbers[0])
			t_dest = append(t_dest, d)
			s, _ := strconv.Atoi(numbers[1])
			t_src = append(t_src, s)
			l, _ := strconv.Atoi(numbers[2])
			t_length = append(t_length, l)
		}
    }
	if len(t_dest) > 0 {
		CreateAndReset(&cons, &t_dest, &t_src, &t_length)
	}
	return seeds, cons
}

func CreateAndReset(cons *[]Converter, t_dest *[]int, t_src *[]int, t_length *[]int) {
	d1 := make([]int, len(*t_dest))
	s1 := make([]int, len(*t_src))
	l1 := make([]int, len(*t_length))

	copy(d1, *t_dest)
	copy(s1, *t_src)
	copy(l1, *t_length)
	*cons = append(*cons, Converter{
		dest: d1,
		src: s1,
		length: l1,
	})
	*t_dest = (*t_dest)[:0]
	*t_src = (*t_src)[:0]
	*t_length = (*t_length)[:0]
}

func ParseSeeds2(line string, seeds *[]int){
	starting := -1
	cline := strings.Split(line, ": ")
	numbers := strings.Split(cline[1], " ")
	for _, n := range(numbers) {
		integer, _ := strconv.Atoi(n)
		if starting < 0 {
			starting = integer
		} else {
			for i:=0; i<integer; i++ {
				*seeds = append(*seeds, starting + i)
			} 
			starting = -1
		}
	}
}

func ParseSeeds(line string, seeds *[]int){
	cline := strings.Split(line, ": ")
	numbers := strings.Split(cline[1], " ")
	for _, n := range(numbers) {
		integer, _ := strconv.Atoi(n)
		*seeds = append(*seeds, integer)
	}
}

func PartOne(path string) int {
	least := 0
	var num int
	seeds, cons := ParseInput(path, false)
	for _, s := range(seeds){
		num = s
		for _, c := range(cons){
			num = c.Convert(num)
		}
		if num < least || least == 0 {
			least = num
		}
	}
	return least
}

func PartTwo(path string) int {
	least := 0
	seeds, cons := ParseInput(path, true)
	back := make(chan int)
	done := make(chan bool)
	for _, s := range(seeds){
		go PartTwoRun(done, back, cons, s)
	}
	go func() {
		for i:=0; i< len(seeds); i++{
			<-done
		}
		close(back)
	}()
	
	for e := range(back) {
		if e < least || least == 0 {
			least = e
			fmt.Println(least)
		}
	}
	return least
}

func PartTwoRun(done chan<- bool, back chan<- int, cons []Converter, seed int) {
	num := seed
	for _, c := range(cons){
		num = c.Convert(num)
	}
	back <-num
	done <-true
}

func main() {
	x := PartTwo("data.txt")
	fmt.Println("ANSWER")
	fmt.Println(x)
}