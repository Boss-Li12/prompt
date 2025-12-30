package main

import (
	"bufio"
	"fmt"
	"os"
)

// Fast I/O 模板
var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func solve() {
	var mode string
	fmt.Fscan(reader, &mode)

	if mode == "first" {
		runAgent()
	} else {
		runBlackslex()
	}
}

func runAgent() {
	var t int
	fmt.Fscan(reader, &t)
	colors := []byte{'r', 'g', 'b'}

	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		adj := make([][]int, n+1)
		for i := 0; i < m; i++ {
			var u, v int
			fmt.Fscan(reader, &u, &v)
			adj[u] = append(adj[u], v)
			adj[v] = append(adj[v], u)
		}

		// BFS 计算到 1 的距离
		dist := make([]int, n+1)
		for i := range dist {
			dist[i] = -1
		}
		dist[1] = 0
		queue := []int{1}

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}

		// 核心策略：按 dist % 3 染色
		res := make([]byte, n)
		for i := 1; i <= n; i++ {
			res[i-1] = colors[dist[i]%3]
		}
		fmt.Fprintln(writer, string(res))
	}
}

func runBlackslex() {
	var t int
	fmt.Fscan(reader, &t)
	colorMap := map[byte]int{'r': 0, 'g': 1, 'b': 2}
	idxToChar := "rgb"

	for ; t > 0; t-- {
		var q int
		fmt.Fscan(reader, &q)
		for i := 0; i < q; i++ {
			var d int
			var s string
			fmt.Fscan(reader, &d, &s)

			// 1. 统计邻居中出现了哪些颜色
			has := [3]bool{}
			for j := 0; j < len(s); j++ {
				has[colorMap[s[j]]] = true
			}

			// 2. 找到“消失”的颜色 k，推断出当前点颜色就是 k
			k := -1
			for j := 0; j < 3; j++ {
				if !has[j] {
					k = j
					break
				}
			}

			// 3. 计算父节点的颜色：(k-1)%3
			// 如果没找到 k (理论上不可能)，默认找 'r' 以防崩溃
			targetIdx := 0
			if k != -1 {
				targetIdx = (k - 1 + 3) % 3
			}
			targetChar := idxToChar[targetIdx]

			// 4. 在邻居字符串 s 中寻找目标颜色的位置
			choice := 1
			for j := 0; j < len(s); j++ {
				if s[j] == targetChar {
					choice = j + 1
					break
				}
			}
			fmt.Fprintln(writer, choice)
		}
	}
}

func main() {
	defer writer.Flush()
	solve()
}
