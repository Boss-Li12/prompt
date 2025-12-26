package main

import (
	"bufio"
	"fmt"
	"os"
)

// 竞赛级 Fast I/O 模板
func solve2179A() {
	var k, x int
	if _, err := fmt.Fscan(reader, &k, &x); err != nil {
		return
	}
	ans := k*x + 1

	fmt.Fprintln(writer, ans)
}

var reader *bufio.Reader
var writer *bufio.Writer

func main() {
	// 初始化 Fast I/O
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	for t > 0 {
		solve2179A()
		t--
	}
}
