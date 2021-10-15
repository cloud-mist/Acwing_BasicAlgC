package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	n, x int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)

	for ; n != 0; n-- {
		fmt.Fscan(in, &x)
		res := divisor(x)
		for _, v := range res {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

func divisor(n int) []int {
	res := make([]int, 0)

	for i := 1; i <= n/i; i++ {
		if n%i == 0 {
			res = append(res, i)
			if i*i != n {
				/* 看是否这个数是个完全平方数，如果是的话，
				根号n就是一个约数,不需要再加了，否则是一对约数*/
				res = append(res, n/i)
			}
		}
	}
	sort.Ints(res)
	return res
}
