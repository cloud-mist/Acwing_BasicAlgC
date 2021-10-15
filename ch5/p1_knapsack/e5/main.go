// 多重背包的 优化版

package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N = 25000
	M = 2010
)

var (
	m, n int
	v, w [N]int
	f    [N]int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	cnt := 0 // 重新排的编号
	for i := 1; i <= n; i++ {
		var a, b, s int // 体积，价值，数量
		fmt.Fscan(r, &a, &b, &s)

		k := 1
		for k <= s {
			cnt++
			v[cnt] = a * k
			w[cnt] = b * k
			s -= k
			k *= 2
		}
		if s > 0 { //剩的那个常数
			cnt++
			v[cnt] = a * s
			w[cnt] = b * s
		}
	}

	// 01背包
	for i := 1; i <= cnt; i++ {
		for j := m; j >= v[i]; j-- {
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
