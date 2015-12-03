package main

import (
	"fmt"
	"io/ioutil"
	// m "math"
	// "strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("input1.txt")
	check(err)
	count := 0
	for i, c := range data {
		cha := string(c)
		if b := count; b == -1 {
			// part 2
			fmt.Println(i)
			break
		}

		// part 1
		if ch := cha; ch == "(" {
			count += 1
		} else if ch == ")" {
			count -= 1
		} else {
			break
		}
	}
	fmt.Println(count)
}
