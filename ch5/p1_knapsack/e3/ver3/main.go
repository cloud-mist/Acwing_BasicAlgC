// 优化2
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
	f    [N]int
)

func main() {
	// 输入
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &v[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := v[i]; j <= m; j++ {
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}
	fmt.Print(f[m])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
