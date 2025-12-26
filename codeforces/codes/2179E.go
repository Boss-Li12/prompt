package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve() {
	var n int
	var x, y int64
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		fmt.Fscan(reader, &n, &x, &y)
		var s string
		fmt.Fscan(reader, &s)
		p := make([]int64, n)

		var sumP int64
		var minA, minB int64
		has0, has1 := false, false

		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &p[i])
			sumP += p[i]
			req := p[i]/2 + 1
			if s[i] == '0' {
				minA += req
				has0 = true
			} else {
				minB += req
				has1 = true
			}
		}

		// 条件 1: 总票数必须覆盖 p_i 需求
		if x+y < sumP {
			fmt.Fprintln(writer, "NO")
			continue
		}

		// 条件 2: 双方必须能赢下各自的保底选区
		if x < minA || y < minB {
			fmt.Fprintln(writer, "NO")
			continue
		}

		// 条件 3: 溢出票数处理 (最关键的逻辑)
		possible := true
		if !has0 {
			// 全是 s=1, B 赢。A 的票必须满足 a_i <= b_i - 1
			// 累加得 x <= y - n
			if x > y-int64(n) {
				possible = false
			}
		} else if !has1 {
			// 全是 s=0, A 赢。B 的票必须满足 b_i <= a_i - 1
			// 累加得 y <= x - int64(n)
			if y > x-int64(n) {
				possible = false
			}
		}
		// 如果既有 0 又有 1，多余的 x 可以塞进 s=0 的区，多余的 y 塞进 s=1 的区，永远可行

		if possible {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func main() {
	solve()
}
