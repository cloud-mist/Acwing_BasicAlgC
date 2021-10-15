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
	// 读入
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &v[i], &w[i])
	}

	// 因为 当背包容量为0,无法装任何，且被初始化了0,所以不再写
	// 当背包 物品数量为0,也无法选任何，且被初始化为0,所以不写
	// 所以都从1开始
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= v[i] { // 因为只有可以装了，才能选右边,所以再取一次max
				f[i][j] = max(f[i][j], f[i-1][j-v[i]]+w[i])
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
