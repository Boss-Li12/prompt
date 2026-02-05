package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

type Flower struct {
	A int // 首位1个
	B int // 补为无数个
}

// 排序 + 二分 + 贪心
// 最终形式一定是 一种补位 + 大于它的所有首位
// 如果多个补位肯定有更贪心的做法（选更大的补位）
// 只看补位也不行，要找比它大的首位，然后计算前缀和
// 还有一个细节：要找比它大的首位是否包含当前补位的首位
func CF1379C(in io.Reader, out io.Writer) {
	var t int
	Fscan(in, &t)
	for tt := 0; tt < t; tt++ {
		var n, m int
		Fscan(in, &n, &m)
		flowers := make([]Flower, m)
		// 1.读入花的数据
		for i := 0; i < m; i++ {
			Fscan(in, &flowers[i].A, &flowers[i].B)
		}
		// 2.按照A从大到小排序
		sort.Slice(flowers, func(i, j int) bool {
			return flowers[i].A > flowers[j].A
		})
		// 3.计算排序后A的前缀和
		preSum := make([]int, m+1)
		for i := 0; i < m; i++ {
			preSum[i+1] = preSum[i] + flowers[i].A
		}
		// 4.固定每个花的B，计算比它大的A的前缀和
		maxSum := 0
		for i := 0; i < m; i++ {
			currentB := flowers[i].B
			// 找到大于等于currentB的个数
			idx := sort.Search(m, func(j int) bool {
				return flowers[j].A < currentB
			})
			curSum := 0
			if idx >= n { // 全部拿A
				curSum = preSum[n]
			} else {
				remain := n - idx
				// 判断currentB的花是否被包含在大的前缀中
				if i < idx {
					curSum = preSum[idx]
					curSum += remain * currentB
				} else {
					curSum = preSum[idx]
					curSum += flowers[i].A + (remain-1)*currentB
				}
			}
			maxSum = max(maxSum, curSum)
		}
		Fprintln(out, maxSum)
	}
}

// bufio.NewReader坑人一手
func main() { CF1379C(bufio.NewReader(os.Stdin), os.Stdout) }
