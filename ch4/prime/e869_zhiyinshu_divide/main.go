// 质因数分解
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n, x int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)

	for ; n != 0; n-- {
		fmt.Fscan(in, &x)
		divide(x)
	}
}

func divide(n int) {
	for i := 2; i <= n/i; i++ {
		if n%i == 0 { // i是质数,  因为：2的话，已经把2的所有倍数都除尽了，然后3的倍数，然后5的倍数,7的倍数...

			s := 0 // 指数
			for n%i == 0 {
				n /= i
				s++
			}
			fmt.Printf("%d %d\n", i, s)
		}
	}

	if n > 1 {
		fmt.Printf("%d %d\n", n, 1)
	}

	fmt.Println()
}
