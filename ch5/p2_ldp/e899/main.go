package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 1010

func main() {
	defer ot.Flush()

	n, m := rnI(), rnI()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = " " + rnS()
	}

	for ; m != 0; m-- {
		a, limit := " "+rnS(), rnI()
		res := 0

		for i := 0; i < n; i++ {
			if calc(a, s[i]) <= limit {
				res++
			}
		}

		fmt.Fprintln(ot, res)
	}

}
func calc(a, b string) int {
	n, m := len(a)-1, len(b)-1
	var f [N][N]int

	for i := 0; i <= n; i++ {
		f[i][0] = i
	}
	for i := 0; i <= m; i++ {
		f[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = min(f[i-1][j]+1, f[i][j-1]+1)

			t := f[i-1][j-1]
			if a[i] != b[j] {
				t++
			}
			f[i][j] = min(f[i][j], t)
		}
	}

	return f[n][m]
}

/* ======================================================================== */
//
//
//
//							 _____   _   _   ____
//							| ____| | \ | | |  _ \
//							|  _|   |  \| | | | | |
//							| |___  | |\  | | |_| |
//							|_____| |_| \_| |____/
//
//
//
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
