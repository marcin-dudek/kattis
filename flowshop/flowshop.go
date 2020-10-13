package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, i, j int32
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &n, &m)
	res := make([]int32, m)
	for i = 0; i < n; i++ {
		var x int32
		for j = 0; j < m; j++ {
			fmt.Fscan(in, &x)
			if j > 0 {
				if res[j-1] > res[j] {
					res[j] = res[j-1] + x
				} else {
					res[j] = res[j] + x
				}
			} else {
				res[j] += x
			}
		}

		if i != n-1 {
			fmt.Printf("%d ", res[m-1])
		} else {
			fmt.Printf("%d", res[m-1])
		}
	}
}
