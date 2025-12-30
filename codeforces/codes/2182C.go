package main

import (
	"bufio"
	"fmt"
	"os"
)

// 使用 bufio 实现快速读写，应对 n=5000 的大数据量
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func solve() {
	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	// 选拔赛 1: 找出 A 与 B 之间合法的偏移量数量
	countAB := 0
	for shift := 0; shift < n; shift++ {
		ok := true
		for i := 0; i < n; i++ {
			if a[i] >= b[(i+shift)%n] {
				ok = false
				break
			}
		}
		if ok {
			countAB++
		}
	}

	// 选拔赛 2: 找出 B 与 C 之间合法的偏移量数量
	countBC := 0
	for shift := 0; shift < n; shift++ {
		ok := true
		for i := 0; i < n; i++ {
			if b[i] >= c[(i+shift)%n] {
				ok = false
				break
			}
		}
		if ok {
			countBC++
		}
	}

	result := int64(n) * int64(countAB) * int64(countBC)
	fmt.Fprintln(writer, result)
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
