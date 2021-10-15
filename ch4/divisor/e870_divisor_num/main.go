package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = int(1e9 + 7)

var (
	t, n int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	sign := make(map[int]int)

	// 求各数约数的次数,存入sign
	for ; t != 0; t-- {
		fmt.Fscan(in, &n)
		for i := 2; i <= n/i; i++ {
			for n%i == 0 {
				n /= i
				sign[i]++
			}
		}

		if n > 1 {
			sign[n]++
		}
	}

	// 根据公式算出个数
	res := 1
	for _, v := range sign {
		res = res * (v + 1) % MOD
	}
	fmt.Println(res)
}
