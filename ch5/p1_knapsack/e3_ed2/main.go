// 优化1
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var (
	v, w [N]int
	n, m int
	f    [N][N]int
)

func main() {
	// 输入
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &v[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= v[i] {
				f[i][j] = max(f[i][j], f[i][j-v[i]]+w[i])
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
