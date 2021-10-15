package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	a, f [N]int
	n    int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &a[i])
	}

	for i := 1; i <= n; i++ {
		f[i] = 1 // 只有i一个数
		for j := 1; j <= i; j++ {
			if a[j] < a[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		res = max(res, f[i])
	}
	fmt.Print(res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
