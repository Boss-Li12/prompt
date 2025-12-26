package main

import (
	"bufio"
	"fmt"
	"os"
)

// 快速 I/O 模板
var reader *bufio.Reader
var writer *bufio.Writer

func solve2179B() {
	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	if n <= 1 {
		fmt.Fprintln(writer, 0)
		return
	}

	// 1. 计算原始总和 S
	totalSum := 0
	for i := 0; i < n-1; i++ {
		totalSum += abs(a[i] - a[i+1])
	}

	// 2. 遍历每一个位置，计算删除该位置能带来的最大收益 (减少量)
	maxDelta := 0

	for i := 0; i < n; i++ {
		delta := 0
		if i == 0 {
			// 删除首位
			delta = abs(a[0] - a[1])
		} else if i == n-1 {
			// 删除末位
			delta = abs(a[n-2] - a[n-1])
		} else {
			// 删除中间位：减去旧的，加上新的
			oldDist := abs(a[i-1]-a[i]) + abs(a[i]-a[i+1])
			newDist := abs(a[i-1] - a[i+1])
			delta = oldDist - newDist
		}
		if delta > maxDelta {
			maxDelta = delta
		}
	}
	fmt.Fprintln(writer, totalSum-maxDelta)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	for t > 0 {
		solve2179B()
		t--
	}
}
