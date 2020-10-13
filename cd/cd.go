package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, n, i, j, x int32
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
	for a != 0 && b != 0 {
		n = 0
		jill := make(map[int32]bool)
		for i = 0; i < a; i++ {
			scanner.Scan()
			//fmt.Fscanln(in, &x)
			x = toInt1(scanner.Bytes())
			jill[x] = true
		}

		for j = 0; j < b; j++ {
			scanner.Scan()
			//fmt.Fscanln(in, &x)
			x = toInt1(scanner.Bytes())
			if jill[x] {
				n++
			}
		}
		fmt.Println(n)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
	}
}

func toInt1(buf []byte) (x int32) {
	for _, v := range buf {
		x = x*10 + int32(v-'0')
	}
	return x
}
