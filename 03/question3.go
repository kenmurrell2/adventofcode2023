package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type key struct {
	row int
	col int
}

func PartOne(path string) int {
	file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	var m [][]string
	row := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		m = append(m, make([]string, len(line)))
		copy(m[row], line)
		row += 1
    }

	num_map := make(map[key]int)
	buffer := 0
	idctr := 0
	var key_list []key
	for row := len(m[0]) -1 ; row >=0 ; row-- {
		place := 1
		for col := len(m) - 1; col >= 0; col-- {
			if integer, err := strconv.Atoi(m[row][col]); err == nil {
				buffer += place*integer
				place = place * 10
				key_list = append(key_list, key{row, col})
			} else {
				if buffer > 0 {
					for _, k := range(key_list) {
						num_map[k] = buffer
					}
					buffer =0
					place = 1
					key_list = key_list[:0]
				}
			}
		}
		if buffer > 0 {
			for _, k := range(key_list) {
				num_map[k] = buffer
			}
			buffer =0
			place = 1
			key_list = key_list[:0]
		}
	}

	total :=0
	for row := 0; row < len(m[0]); row++ {
		for col := 0; col < len(m); col++ {
			sym := m[row][col]
			if sym == "*" || sym == "#" || sym == "$" || sym == "+" {
				//check above
				if row > 0 {
					k := key{row-1, col}
					if val, ok := num_map[k]; ok {
						total += val
						remove(&num_map, val)
					}
				}
				//check below
				if row < len(m) - 1 {
					k := key{row+1, col}
					if val, ok := num_map[k]; ok {
						total += val
						remove(&num_map, val)
					}
				}
				//check left
				if col > 0 {
					k := key{row, col-1}
					if val, ok := num_map[k]; ok {
						total += val
						remove(&num_map, val)
					}
				}
				//check right
				if col < len(m[0]) -1 {
					k := key{row, col+1}
					if val, ok := num_map[k]; ok {
						total += val
						remove(&num_map, val)
					}
				}

				//check above left
				if row > 0 && col > 0 {
					k := key{row-1, col-1}
					if val, ok := num_map[k]; ok {
						total += val
						remove(&num_map, val)
					}
				}
				//check above right
				if row > 0 && col < len(m[0]) -1 {
					k := key{row-1, col+1}
					if val, ok := num_map[k]; ok {
						total += val
						remove(&num_map, val)
					}
				}
				//check below left
				if row < len(m) - 1 && col > 0 {
					k := key{row+1, col-1}
					if val, ok := num_map[k]; ok {
						total += val
						remove(&num_map, val)
					}
				}
				//check below right
				if row < len(m) - 1 && col < len(m[0]) -1 {
					k := key{row+1, col+1}
					if val, ok := num_map[k]; ok {
						total += val
						remove(&num_map, val)
					}
				}
			}
		}
	}

	return total
}

func remove(num_map *map[key]int, val int) {
	var test []key
	for k, v := range(*num_map) {
		if v == val {
			test = append(test, k)
		}
	}
	for _, k := range(test) {
		delete(*num_map, k)
	}
}



func PartTwo(path string) int {
	return 0
}