package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var max = readIntArray(n)
	if max == 0 {
		fmt.Println("none")
	} else {
		fmt.Println(max)
	}
}

func readIntArray(n int) int {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	fields := strings.Fields(s)
	var duplicates = make([]int, 7)
	var single = make([]int, 7)
	for i, item := range fields {
		y, _ := strconv.ParseInt(item, 10, 32)
		z := int(y)
		if duplicates[z] == 0 {
			if single[z] == 0 {
				single[z] = i + 1
			} else {
				single[z] = 0
				duplicates[z] = 1
			}
		}
	}

	var index int
	for _, j := range single {
		if j > 0 {
			index = j
		}
	}

	return index
}
