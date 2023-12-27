package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Point struct {
	row byte
	col byte
}

type Empty struct {}

func ParseInput(path string) (*[][]bool, *Point) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var grid [][]bool
	row := 0
	var start Point
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		lineSplit := strings.Split(text, "")
		var gridline []bool
		for col, char := range lineSplit {
			if char == "." {
				gridline = append(gridline, true)
			} else if char == "#" {
				gridline = append(gridline, false)
			} else if char == "S" {
				gridline = append(gridline, true)
				start = Point{byte(row), byte(col)}
			}
		}
		grid = append(grid, gridline)
		row += 1
	}
	return &grid, &start
}



func PartOne(path string, steps int) int {
	grid, start := ParseInput(path)
	visit := make(map[Point]Empty)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(1)
	go run(grid, start, &visit, steps, &wg, &mu)
	wg.Wait()
	return len(visit)
}

func run(grid *[][]bool, p *Point, visit *map[Point]Empty, remaining int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	dirs := []Point {
		Point{p.row+1, p.col},
		Point{p.row-1, p.col},
		Point{p.row, p.col-1},
		Point{p.row, p.col+1},
	}
	if remaining == 0 {
		mu.Lock()
		(*visit)[*p] = Empty{}
		mu.Unlock()
		return
	}
	for _, newPoint := range dirs {
		if (*grid)[newPoint.col][newPoint.row] {
			wg.Add(1)
			go run(grid, &newPoint, visit, remaining - 1, wg, mu)
		}
	}
}

func PartTwo(path string) int {
	return 0
}

func main() {
	start1 := time.Now()
	r1 := PartOne("data.txt", 64)
	elapsed1 := time.Since(start1).Seconds()
	fmt.Printf("ANSWER ONE: %d; elapsed %fs\n", r1, elapsed1) 
	// start2 := time.Now()
	// r2 := PartTwo("data.txt")
	// elapsed2 := time.Since(start2).Seconds()
	// fmt.Printf("ANSWER TWO: %d; elapsed %fs\n", r2, elapsed2)
}
