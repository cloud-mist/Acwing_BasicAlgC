package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer ot.Flush()
	n := rnI()

	for ; n != 0; n-- {
		x := rnI()
		divide(x)
		fmt.Println()
	}

}

func divide(n int) {
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			s := 0
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
}

//{{{
/* ============================PART 1: I/O ================================== */

var (
	ot = bufio.NewWriterSize(os.Stdout, int(1e6))
	in = bufio.NewScanner(os.Stdin)
)

func init()        { in.Split(bufio.ScanWords); in.Buffer(make([]byte, 4096), int(1e9)) }
func rnS() string  { in.Scan(); return in.Text() }
func rnI() int     { i, _ := strconv.Atoi(rnS()); return i }
func rnF() float64 { f, _ := strconv.ParseFloat(rnS(), 64); return f }

func rsI(l, r int) []int {
	t := make([]int, r)
	for i := l; i < r; i++ {
		t[i] = rnI()
	}
	return t
}

/* ===========================PART 2: Math Func ============================  */
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func memset(a []int, v int) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}

//}}}
