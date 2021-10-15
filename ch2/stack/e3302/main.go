package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var (
	s   string
	num []int
	op  []rune
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &s)
	sign := map[rune]int{'+': 1, '-': 1, '*': 2, '/': 2}

	long := len(s)
	for i := 0; i < long; i++ {
		c := rune(s[i])

		// 如果是数字,就读入并且同时“计算”这个数字，直到不是数字的地方为止
		if unicode.IsDigit(c) {
			x, j := 0, i
			for j < long && unicode.IsDigit(rune(s[j])) {
				x = x*10 + int(s[j]-'0')
				j++
			}
			num = append(num, x) // 将整个数字压入栈
			i = j - 1            // i要跳到数字末尾字符的下一个位置

			// 左括号直接入栈
		} else if c == '(' {
			op = append(op, c)

			// 右括号循环计算栈上面两个数，直到遇到左括号,最后左括号出栈
		} else if c == ')' {
			for op[len(op)-1] != '(' {
				eval()
			}
			op = op[:len(op)-1]

			// 如果遇到运算符的优先级<= 栈顶优先级且栈不空,就循环计算
		} else {
			for len(op) > 0 && sign[op[topO(op)]] >= sign[c] {
				eval()
			}
			op = append(op, c) // 之后，新符号入栈
		}
	}

	// 取出剩余符号,计算
	for len(op) != 0 {
		eval()
	}
	fmt.Print(num[0])

}

// 从栈取出两个数计算
func eval() {
	b := num[topN(num)]
	num = num[:topN(num)]
	a := num[topN(num)]
	num = num[:topN(num)]
	c := op[topO(op)]
	op = op[:topO(op)]

	x := 0
	if c == '+' {
		x = a + b
	} else if c == '-' {
		x = a - b
	} else if c == '*' {
		x = a * b
	} else {
		x = a / b
	}
	num = append(num, x)
}

func topN(a []int) int {
	return len(a) - 1
}

func topO(a []rune) int {
	return len(a) - 1
}
