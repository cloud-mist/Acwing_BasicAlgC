package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	defer ot.Flush()
	n := rnI()
	for ; n != 0; n-- {
		op, str := rnS(), rnS()
		if op == "I" {
			insert(str)
		} else {
			fmt.Println(query(str))
		}
	}
}

const N = 20010

var (
	son [N][26]int
	cnt [N]int
	idx int
)

func insert(s string) {
	p := 0
	for i := 0; i < len(s); i++ {
		u := s[i] - 'a'
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
	cnt[p]++
}

func query(s string) int {
	p := 0
	for i := 0; i < len(s); i++ {
		u := s[i] - 'a'
		if son[p][u] == 0 {
			return 0
		}
		p = son[p][u]
	}
	return cnt[p]
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
