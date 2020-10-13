package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, x, y, col, row int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//fmt.Sscanf(scanner.Text(), "%d", &n)
	n = toInt1(scanner.Bytes())
	//in := bufio.NewReader(os.Stdin)
	//fmt.Fscan(in, &n)
	sum := make(map[int]int, n)
	dif := make(map[int]int, n)

	S := (n - 1) * n / 2

	for i := 0; i < n; i++ {
		//fmt.Fscan(in, &x, &y)
		scanner.Scan()
		x, y = toInt2(scanner.Bytes())
		//fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		col += x
		row += y
		sum[x+y] = 1
		dif[x-y] = 1
	}

	if col != S || row != S {
		fmt.Println("INCORRECT")
		return
	}

	if len(dif) != n || len(sum) != n {
		fmt.Println("INCORRECT")
		return
	}

	fmt.Println("CORRECT")
}

func toInt1(buf []byte) (x int) {
	for _, v := range buf {
		x = x*10 + int(v-'0')
	}
	return x
}

func toInt2(buf []byte) (x, y int) {
	first := true
	for _, v := range buf {
		if v == 32 {
			first = false
			continue
		}
		if first {
			x = x*10 + int(v-'0')
		} else {
			y = y*10 + int(v-'0')
		}
	}

	return x, y
}
