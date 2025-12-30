package main

import (
	"bufio"
	"fmt"
	"os"
)

// 竞赛级 Fast I/O 模板
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func solve() {
	var a, b int64
	if _, err := fmt.Fscan(reader, &a, &b); err != nil {
		return
	}

	ans := 0
	var sum1, sum2 int64
	sum1 = 0
	sum2 = 0

	for n := 1; n <= 60; n++ {
		// 第 n 层的大小是 2^(n-1)
		sz := int64(1) << (n - 1)

		// 根据层数的奇偶性分配到对应的累加器
		if n%2 == 1 {
			sum1 += sz
		} else {
			sum2 += sz
		}

		// 检查两种起始方案是否可行
		// 方案一：白巧克力涂 sum1，黑巧克力涂 sum2
		// 方案二：白巧克力涂 sum2，黑巧克力涂 sum1
		can1 := (a >= sum1 && b >= sum2)
		can2 := (a >= sum2 && b >= sum1)

		if can1 || can2 {
			ans = n
		} else {
			// 如果当前层数都无法满足，更高层数更不可能
			break
		}
	}

	fmt.Fprintln(writer, ans)
}

func main() {
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	for t > 0 {
		solve()
		t--
	}
}
