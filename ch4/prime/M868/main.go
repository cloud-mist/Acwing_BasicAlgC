// 试除法---质数判定
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	n, a := 0, 0
	fmt.Fscan(r, &n)

	for ; n != 0; n-- {
		fmt.Fscan(r, &a)
		if isPrime(a) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
