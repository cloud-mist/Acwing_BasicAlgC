// 石子合并
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	N   = 310
	INF = int(1e8)
)

var (
	f [N][N]int
	a [N]int
)

func main() {
	defer ot.Flush()

	// input & 初始化f & 前缀和
	n := rnI()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			f[i][j] = INF
		}
	}
	for i := 1; i <= n; i++ {
		f[i][i] = 0
		a[i] = rnI()
		a[i] += a[i-1]
	}

	for len := 2; len <= n; len++ { // 阶段
		for l := 1; l <= n-len+1; l++ { //状态：左端点
			r := l + len - 1
			for k := l; k < r; k++ { // 决策
				f[l][r] = min(f[l][r], f[l][k]+f[k+1][r]+a[r]-a[l-1])
			}
		}
	}

	fmt.Print(f[1][n])
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
