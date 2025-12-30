package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

// IntHeap 实现标准库的 heap 接口，用于存储 diff (z - y)
type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int64)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Friend struct {
	x, y, z, diff int64
}

func solve() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n, m int
		var k int64
		fmt.Fscan(reader, &n, &m, &k)

		boxes := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(reader, &boxes[i])
		}
		// 盒子按美丽值降序排列
		sort.Slice(boxes, func(i, j int) bool { return boxes[i] > boxes[j] })

		friends := make([]Friend, n)
		var totalY int64
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &friends[i].x, &friends[i].y, &friends[i].z)
			friends[i].diff = friends[i].z - friends[i].y
			totalY += friends[i].y
		}

		// 剩余可支配预算
		remK := k - totalY

		// 朋友按要求美丽值 x 降序排列
		sort.Slice(friends, func(i, j int) bool {
			return friends[i].x > friends[j].x
		})

		h := &IntHeap{}
		heap.Init(h)
		boxIdx := 0
		var paidList []int64

		for i := 0; i < n; i++ {
			// 如果还有盒子，且当前盒子能满足该朋友
			if boxIdx < m && int64(boxes[boxIdx]) >= friends[i].x {
				heap.Push(h, friends[i].diff)
				boxIdx++
			} else {
				// 没盒子了，或者盒子不够美。
				// 贪心：如果这个人的 diff 比堆里最小的还大，
				// 且堆里的那个人的 x 限制更松（一定能被当前盒子满足），
				// 则把盒子抢过来给当前 diff 大的人，让 diff 小的人去“花钱”
				if h.Len() > 0 && (*h)[0] < friends[i].diff {
					// 弹出 diff 最小的人，他必须花钱了
					paidList = append(paidList, heap.Pop(h).(int64))
					// 当前朋友占用这个腾出来的盒子
					heap.Push(h, friends[i].diff)
				} else {
					// 否则，当前朋友直接进入花钱名单
					paidList = append(paidList, friends[i].diff)
				}
			}
		}

		// 现在堆里的人都通过盒子变开心了。
		// 剩下 paidList 里的朋友需要靠 remK 变开心。
		sort.Slice(paidList, func(i, j int) bool { return paidList[i] < paidList[j] })

		ans := h.Len()
		for _, cost := range paidList {
			if remK >= cost {
				remK -= cost
				ans++
			} else {
				break
			}
		}
		fmt.Fprintln(writer, ans)
	}
}

func main() {
	solve()
}
