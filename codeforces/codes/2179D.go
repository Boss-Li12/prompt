package main

import (
	"bufio"
	"fmt"
	"os"
)

// 快速 I/O 模板
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func solve() {
	var n int
	_, err := fmt.Fscan(reader, &n)
	if err != nil {
		return
	}

	total := 1 << n
	used := make([]bool, total)
	p := make([]int, 0, total)

	// 第一个数必须是全 1，以保证 S(p) 起始最大
	curMask := total - 1
	p = append(p, curMask)
	used[curMask] = true

	// 构造逻辑：从最高位开始，逐位消除 1
	// 每次消除一个 bit 后，立即把能维持当前 mask 的最小可用数填入
	for i := n - 1; i >= 0; i-- {
		// 消除第 i 位
		curMask ^= (1 << i)

		// 字典序最小化：
		// 1. 放入当前能维持新 mask 的最小数（即新 mask 本身）
		if !used[curMask] {
			p = append(p, curMask)
			used[curMask] = true
		}

		// 2. 回补：所有包含当前 curMask 且之前未被使用的数
		// 这些数与当前前缀与结果做 AND 后，结果依然是 curMask，不会让 S(p) 变小
		// 为了字典序，我们要从小到大遍历。
		// 注意：这里的回补范围其实是 [curMask, total-1] 之间满足 (x & curMask == curMask) 的数
		for v := 0; v < total; v++ {
			if !used[v] && (v&curMask) == curMask {
				p = append(p, v)
				used[v] = true
			}
		}
	}

	for i := 0; i < len(p); i++ {
		if i == len(p)-1 {
			fmt.Fprintf(writer, "%d\n", p[i])
		} else {
			fmt.Fprintf(writer, "%d ", p[i])
		}
	}
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
