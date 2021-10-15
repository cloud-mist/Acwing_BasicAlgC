// 优化为1维版
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
	// 读入
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &v[i], &w[i])
	}

	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- { // 从小往大，会误用到同一层刚更新的值,造成错误
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
