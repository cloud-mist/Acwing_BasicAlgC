// 堆排序
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const N = 100010

var (
	h    [N]int
	size int
)

func main() {
	defer ot.Flush()

	n, m := rnI(), rnI()
	for i := 1; i <= n; i++ {
		h[i] = rnI()
	}
	size = n

	// 建堆
	for i := n / 2; i != 0; i-- {
		down(i)
	}

	// 输出最小值
	for ; m != 0; m-- {
		fmt.Fprintf(ot, "%d ", h[1])
		h[1] = h[size]
		size--
		down(1)
	}
}

func down(x int) {
	t := x             // t代表最小值
	l, r := x*2, x*2+1 // 左右儿子

	// 三者进行比较，找到最小
	if l <= size && h[l] < h[t] {
		t = l
	}
	if r <= size && h[r] < h[t] {
		t = r
	}

	// t!=x,说明需要交换，交换完要递归down
	if t != x {
		h[t], h[x] = h[x], h[t]
		down(t)
	}
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

/* ==== G0 ===== */
const BUFSIZE = int(1e6)

var (
	ot = bufio.NewWriterSize(os.Stdout, BUFSIZE)
	in = newScanner(os.Stdin)
)

type scanner struct{ sc *bufio.Scanner }

func newScanner(input io.Reader) *scanner {
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), int(1e9))
	return &scanner{sc}
}

/* ==== G1 ===== */
func rnS() string  { in.sc.Scan(); return in.sc.Text() }
func rnI() int     { i, _ := strconv.Atoi(rnS()); return i }
func rnF() float64 { f, _ := strconv.ParseFloat(rnS(), 64); return f }

/* ==== G2 ===== */
func rsI(n int) []int {
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = rnI()
	}
	return t
}
func rsF(n int) []float64 {
	t := make([]float64, n)
	for i := 0; i < n; i++ {
		t[i] = rnF()
	}
	return t
}
func rsS(n int) []string {
	t := make([]string, n)
	for i := 0; i < n; i++ {
		t[i] = rnS()
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
