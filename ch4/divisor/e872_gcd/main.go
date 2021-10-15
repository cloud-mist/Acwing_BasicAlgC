package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100010

var (
	n, a, b int
)

func main() {
	in := bufio.NewReader(os.Stdin)
	ot := bufio.NewWriter(os.Stdout)
	defer ot.Flush()
	fmt.Fscan(in, &n)

	for ; n != 0; n-- {
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(ot, gcd(a, b))
	}

}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}
