// 朴素版，和完全背包的朴素版一致
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 2010

var (
	n, m    int
	v, w, s [N]int
	f       [N][N]int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &v[i], &w[i], &s[i])
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k <= s[i] && k*v[i] <= j; k++ {
				f[i][j] = max(f[i][j], f[i-1][j-v[i]*k]+w[i]*k)
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
