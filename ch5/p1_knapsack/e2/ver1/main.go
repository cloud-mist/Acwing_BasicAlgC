// 基础版
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 1010

var (
	v, w [N]int
	f    [N][N]int
)

func main() {
	defer ot.Flush()
	// input
	n, m := rnI(), rnI()
	for i := 1; i <= n; i++ {
		v[i], w[i] = rnI(), rnI()
	}

	// calc
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = f[i-1][j]
			// 下面的情况不一定存在,所以需要判断一下，以保证存在
			if j >= v[i] {
				f[i][j] = max(f[i][j], f[i][j-v[i]]+w[i])
			}
		}
	}

	// output
	fmt.Print(f[n][m])
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
/* ============================PART1: I/O ================================== */

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
