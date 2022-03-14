package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 1010

var f [N][N]int

func main() {
	defer ot.Flush()

	// input
	n, a := rnI(), " "+rnS()
	m, b := rnI(), " "+rnS()

	// init
	for i := 0; i <= m; i++ {
		//A的前0个匹配b的前i个，有多少个，a要增加多少个
		f[0][i] = i
	}
	for i := 0; i <= n; i++ {
		//A的前i个匹配b的前0个，a有多少个，要删除多少个
		f[i][0] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = min(f[i-1][j]+1, f[i][j-1]+1)

			tmp := f[i-1][j-1]
			if a[i] != b[j] {
				tmp++
			}
			f[i][j] = min(f[i][j], tmp)
		}
	}

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
