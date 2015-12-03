package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func bubbleSort(sl []int) []int {
	tmp := 0
	for i := 0; i < len(sl); i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i] > sl[j] {
				tmp = sl[i]
				sl[i] = sl[j]
				sl[j] = tmp
			}
		}
	}
	return sl
}

func main() {
	file, err := os.Open("input2.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	words := [][]int{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "x")
		w, err := strconv.Atoi(line[0])
		check(err)
		h, err := strconv.Atoi(line[1])
		check(err)
		l, err := strconv.Atoi(line[2])
		check(err)
		sortedDimensions := bubbleSort([]int{w, h, l})
		words = append(words, sortedDimensions)
	}
	totalWrapNeeded := 0
	totalRibbonNeeded := 0
	s := 0
	m := 0
	l := 0
	for _, word := range words {
		if len(word) < 3 {
			break
		}
		s = word[0]
		m = word[1]
		l = word[2]
		totalWrapNeeded += (2*s*m + 2*m*l + 2*l*s) + s
		totalRibbonNeeded += 2*s + 2*m + s*m*l
	}
	fmt.Println(totalRibbonNeeded, totalRibbonNeeded)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
