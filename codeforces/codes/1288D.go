package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1288D(in io.Reader, out io.Writer) {
	for tt := 0; tt < 1; tt++ {
		var n, m int
		Fscan(in, &n)
		Fscan(in, &m)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			a[i] = make([]int, m)
			for j := 0; j < m; j++ {
				Fscan(in, &a[i][j])
			}

		}
		// 二分答案 + check函数判断是否有两行合成的b的最小值大于等于x，用bitmask操作
		ansRow1, ansRow2 := 1, 1
		l, r := 0, 1000000000
		for l <= r {
			mid := (l + r) / 2
			ok, row1, row2 := check(mid, n, m, a)
			if ok {
				l = mid + 1
				ansRow1, ansRow2 = row1, row2
			} else {
				r = mid - 1
			}
		}
		Fprintf(out, "%d %d\n", ansRow1, ansRow2)
	}
}

// check函数判断是否有两行合成的b的最小值大于等于x, 同时记录对应的行号
func check(x int, n int, m int, a [][]int) (bool, int, int) {
	limit := 1 << m
	visit := make([]int, limit)
	for i := 0; i < limit; i++ {
		visit[i] = -1 // 初始化为-1
	}
	// 记录bitmask到行号的映射
	for i := 0; i < n; i++ {
		tempBitMask := 0
		for j := 0; j < m; j++ {
			if a[i][j] >= x {
				tempBitMask += 1 << j
			}
		}
		visit[tempBitMask] = i
	}
	// 两次循环看是否有两行bitmask的最大值即or操作为全1即limit - 1
	for i := 0; i < limit; i++ {
		if visit[i] != -1 {
			for j := 0; j < limit; j++ {
				if visit[j] != -1 && i|j == limit-1 {
					return true, visit[i] + 1, visit[j] + 1
				}
			}
		}
	}

	return false, -1, -1
}

// bufio.NewReader坑人一手
func main() { CF1288D(bufio.NewReader(os.Stdin), os.Stdout) }
