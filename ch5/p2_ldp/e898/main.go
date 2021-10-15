package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N   = 510
	INF = int(1e7)
)

var (
	a, f [N][N]int
	n    int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)

	// 输入
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fscan(r, &a[i][j])
		}
	}

	// 初始化
	for i := 0; i <= n; i++ {
		for j := 0; j <= i+1; j++ {
			f[i][j] = -INF
		}
	}

	// 计算
	f[1][1] = a[1][1]
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			f[i][j] = max(f[i-1][j-1]+a[i][j], f[i-1][j]+a[i][j])
		}
	}

	// 处理结果,是最后每个数字结果的一个最大值
	res := -INF
	for i := 1; i <= n; i++ {
		res = max(res, f[n][i])
	}
	fmt.Print(res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
