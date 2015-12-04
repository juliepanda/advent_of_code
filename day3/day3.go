package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("input3.txt")
	check(err)
	visited := make(map[string]bool)
	count := 1
	x := 0
	y := 0
	rx := 0
	ry := 0
	visited["0-0"] = true
	key := ""
	for i, c := range data {
		cha := string(c)
		if i%2 == 0 {
			if cha == "^" {
				y += 1
			} else if cha == ">" {
				x += 1
			} else if cha == "v" {
				y -= 1
			} else if cha == "<" {
				x -= 1
			} else {
				break
			}
			key = strconv.Itoa(x) + "-" + strconv.Itoa(y)
		} else {
			if cha == "^" {
				ry += 1
			} else if cha == ">" {
				rx += 1
			} else if cha == "v" {
				ry -= 1
			} else if cha == "<" {
				rx -= 1
			} else {
				break
			}
			key = strconv.Itoa(rx) + "-" + strconv.Itoa(ry)
		}
		if visited[key] {
			// do nothing
		} else {
			count += 1
			visited[key] = true
		}
	}
	fmt.Println(count)
}
