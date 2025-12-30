package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 使用 bufio 加速 I/O
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func solve() {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)

	// 条件 A: 包含 2026
	has2026 := strings.Contains(s, "2026")
	// 条件 B: 不包含 2025
	no2025 := !strings.Contains(s, "2025")

	// 只要满足其中一个条件，就是 New Year string
	if has2026 || no2025 {
		fmt.Fprintln(writer, 0)
		return
	}

	// 如果两个都不满足，说明现在有 2025 且没有 2026
	fmt.Fprintln(writer, 1)
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
