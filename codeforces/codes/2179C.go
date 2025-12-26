package main

import (
	"bufio"
	"fmt"
	"os"
)

// 竞赛级 Fast I/O
var reader *bufio.Reader
var writer *bufio.Writer

func solve2179C() {
	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	a := make([]int, n)
	minVal := int(1e9 + 7)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		if a[i] < minVal {
			minVal = a[i]
		}
	}

	// 候选解 1: 将所有数取模变为 0。
	// 取 x_i = a[i], 则 k = min(a)。
	ans := minVal

	// 候选解 2: 将所有数取模变为 minVal。
	// 此时要求每个 a[i] - minVal 的模数 x 必须大于 minVal。
	// 因此对所有 a[i] > minVal，必须满足 a[i] - minVal > minVal => a[i] > 2*minVal。
	canSetToMin := true
	currentMinDiff := int(2e9 + 7)

	countGreater := 0
	for _, v := range a {
		if v > minVal {
			countGreater++
			diff := v - minVal
			if diff <= minVal {
				canSetToMin = false
				break
			}
			if diff < currentMinDiff {
				currentMinDiff = diff
			}
		}
	}

	// 如果数组中除了 minVal 以外还有别的数，且都满足 diff > minVal
	if countGreater > 0 && canSetToMin {
		if currentMinDiff > ans {
			ans = currentMinDiff
		}
	}

	fmt.Fprintln(writer, ans)
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	if _, err := fmt.Fscan(reader, &t); err != nil {
		return
	}
	for t > 0 {
		solve2179C()
		t--
	}
}
