package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1000100

var (
	n, cnt int
	sign   [N]bool
	num    [N]int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	primesPro(n)
	fmt.Println(cnt)
}

func primesPro(n int) {
	for i := 2; i <= n; i++ {
		if !sign[i] { // 如果没有被标记，说明是素数
			num[cnt] = i
			cnt++
		}

		for j := 0; num[j] <= n/i; j++ { // 枚举质数
			sign[num[j]*i] = true // 把质数和i的乘积筛掉

			if i%num[j] == 0 { //
				break
			}
		}
	}

}

// erato筛法 NloglogN
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
