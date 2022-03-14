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
	a := rsI(0, n)

	mergeSort(a, 0, n-1)

	for i := range a {
		fmt.Fprintf(ot, "%d ", a[i])
	}

}

func mergeSort(a []int, l, r int) {
	if l >= r {
		return
	}

	// 递归
	mid := (l + r) >> 1
	mergeSort(a, l, mid)
	mergeSort(a, mid+1, r)

	// 合并
	i, j := l, mid+1
	tmp := make([]int, 0)
	for i <= mid && j <= r {
		if a[i] <= a[j] {
			tmp = append(tmp, a[i])
			i++
		} else {
			tmp = append(tmp, a[j])
			j++
		}
	}

	// 如果有剩余,吧剩余合并
	for ; i <= mid; i++ {
		tmp = append(tmp, a[i])
	}
	for ; j <= r; j++ {
		tmp = append(tmp, a[j])
	}

	// 将东西赋值给源数组
	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		a[i] = tmp[j]
	}

}

/* ======================================================================== */
//
//
//
//				 _____   _   _   ____
//				| ____| | \ | | |  _ \
//				|  _|   |  \| | | | | |
//				| |___  | |\  | | |_| |
//				|_____| |_| \_| |____/
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
