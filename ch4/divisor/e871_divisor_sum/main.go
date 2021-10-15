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

	// 根据公式算出之和
	res := 1
	for i, v := range sign {
		tmp := 1
		for v != 0 {
			v--
			tmp = (tmp*i + 1) % MOD
		}
		res = res * tmp % MOD
	}

	fmt.Println(res)
}
