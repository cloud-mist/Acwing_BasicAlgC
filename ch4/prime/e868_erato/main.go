// erato筛法 NloglogN
package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1000100

var (
	n    int
	sign [N]bool
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	primes(n)
}

func primes(n int) {
	k := 0 // 记录个数
	for i := 2; i <= n; i++ {
		if sign[i] { // 如果已经标记
			continue
		}

		k++ // 没被标记就是素数

		// 从i的平方开始标记,前面已经被更小的数标记过了
		for j := i; j <= n/i; j++ {
			sign[i*j] = true
		}
	}
	fmt.Println(k)
}
