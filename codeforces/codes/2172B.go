package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// SegTree 线段树用于维护区间最大值
type SegTree struct {
	n    int
	tree []float64
}

func NewSegTree(n int) *SegTree {
	return &SegTree{n: n, tree: make([]float64, 2*n)}
}

func (st *SegTree) Update(i int, val float64) {
	for i += st.n; i > 0; i >>= 1 {
		if val > st.tree[i] {
			st.tree[i] = val
		} else {
			break // 剪枝：如果当前值不比树中大，则无需向上更新
		}
	}
}

func (st *SegTree) Query(l, r int) float64 {
	res := 0.0
	l += st.n
	r += st.n
	for l < r {
		if l&1 == 1 {
			res = math.Max(res, st.tree[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = math.Max(res, st.tree[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

type Bus struct {
	s, t int
	val  float64
}

type Person struct {
	id  int
	p   int
	ans float64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, l, x, y int
	fmt.Fscan(in, &n, &m, &l, &x, &y)

	invX := 1.0 / float64(x)
	invY := 1.0 / float64(y)
	constantC := invY - invX

	buses := make([]Bus, n)
	coords := make([]int, 0, n+m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &buses[i].s, &buses[i].t)
		buses[i].val = float64(buses[i].t)*constantC + float64(buses[i].s)*invX
		coords = append(coords, buses[i].t)
	}

	people := make([]Person, m)
	for i := 0; i < m; i++ {
		people[i].id = i
		fmt.Fscan(in, &people[i].p)
		coords = append(coords, people[i].p)
	}

	// 坐标离散化 (处理 t_j 和 p_i)
	sort.Ints(coords)
	uniqueCount := 0
	if len(coords) > 0 {
		uniqueCount = 1
		for i := 1; i < len(coords); i++ {
			if coords[i] != coords[i-1] {
				coords[uniqueCount] = coords[i]
				uniqueCount++
			}
		}
		coords = coords[:uniqueCount]
	}

	getIdx := func(val int) int {
		return sort.SearchInts(coords, val)
	}

	// 排序以便双指针/扫描线
	sort.Slice(buses, func(i, j int) bool { return buses[i].s < buses[j].s })
	sortedPeople := make([]*Person, m)
	for i := 0; i < m; i++ {
		sortedPeople[i] = &people[i]
	}
	sort.Slice(sortedPeople, func(i, j int) bool { return sortedPeople[i].p < sortedPeople[j].p })

	st := NewSegTree(uniqueCount)
	busIdx := 0
	ly := float64(l) * invY

	for _, p := range sortedPeople {
		// 将所有起点 s_j <= p.p 的公交车加入线段树
		for busIdx < n && buses[busIdx].s <= p.p {
			st.Update(getIdx(buses[busIdx].t), buses[busIdx].val)
			busIdx++
		}
		// 查询所有 t_j >= p.p 的最大 Val_j
		maxBusVal := st.Query(getIdx(p.p), uniqueCount)
		// 比较步行和坐车
		walkVal := float64(p.p) * invY
		bestZ := math.Max(walkVal, maxBusVal)
		p.ans = ly - bestZ
	}

	for i := 0; i < m; i++ {
		fmt.Fprintf(out, "%.10f\n", people[i].ans)
	}
}
