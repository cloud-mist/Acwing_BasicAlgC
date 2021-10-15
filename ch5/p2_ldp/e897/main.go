package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	n, m int
	a, b string
	f    [N][N]int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m, &a, &b)
	// 解决边界问题
	a = " " + a
	b = " " + b

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = max(f[i-1][j], f[i][j-1])
			if a[i] == b[j] {
				f[i][j] = max(f[i][j], f[i-1][j-1]+1)
			}
		}
	}

	fmt.Print(f[n][m])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
