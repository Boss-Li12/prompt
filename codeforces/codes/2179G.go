package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF2179G(in io.Reader, out io.Writer) {
	var t int
	Fscan(in, &t)
	for tt := 0; tt < t; tt++ {
		var n int
		Fscan(in, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
		}
		ans := 0
		Fprintln(out, ans)
	}
}

// bufio.NewReader坑人一手
func main() { CF2179G(bufio.NewReader(os.Stdin), os.Stdout) }
