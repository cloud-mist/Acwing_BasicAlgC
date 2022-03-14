package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 100010

var (
	p, d [N]int
)

func find(x int) int {
	if p[x] != x {
		t := find(p[x]) // 记住根的编号
		d[x] += d[p[x]] // 计算距离
		p[x] = t
	}
	return p[x]
}

func main() {
	defer ot.Flush()
	n, k := rnI(), rnI()
	for i := 0; i < n; i++ {
		p[i] = i
	}

	res := 0
	for ; k != 0; k-- {
		dis, x, y := rnI(), rnI(), rnI()
		if x > n || y > n {
			// 超过了动物标号,直接假话
			res++
		} else {
			px, py := find(x), find(y)
			if dis == 1 {
				if px == py && (d[x]-d[y])%3 != 0 {
					// 如果是在同一集合，但余数不同的话，说明假话
					res++
				} else if px != py {
					// 如果不同集合，就接上去，更新距离
					p[px] = py
					d[px] = d[y] - d[x]
				}
			} else {
				if px == py && (d[x]-d[y]-1)%3 != 0 {
					res++
				} else if px != py {
					p[px] = py
					d[px] = d[y] - d[x] + 1
				}
			}
		}
	}

	fmt.Print(res)
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
