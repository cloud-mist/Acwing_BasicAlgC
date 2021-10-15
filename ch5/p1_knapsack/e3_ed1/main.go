// 基础版
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
			for k := 0; k*v[i] <= j; k++ {
				f[i][j] = max(f[i-1][j-k*v[i]]+k*w[i], f[i][j])
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
