package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

// Fast I/O
var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve() {
	var n, m int
	var l, x, y float64
	fmt.Fscan(in, &n, &m, &l, &x, &y)

	type bus struct {
		s, t, r, val float64
	}

	buses := make([]bus, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &buses[i].s, &buses[i].t)
		// R_i = t_i - (y/x)*(t_i - s_i)
		buses[i].r = buses[i].t - (y/x)*(buses[i].t-buses[i].s)
		// B_i = (t_i - s_i)/x + (L - t_i)/y
		buses[i].val = (buses[i].t-buses[i].s)/x + (l-buses[i].t)/y
	}

	peopleP := make([]float64, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &peopleP[i])
	}

	// 坐标离散化：收集所有 s_i, r_i, p_j
	coords := make([]float64, 0, n*2+m)
	for _, b := range buses {
		coords = append(coords, b.s, b.r)
	}
	coords = append(coords, peopleP...)
	sort.Float64s(coords)

	// 去重
	uniqueCoords := coords[:0]
	for i := 0; i < len(coords); i++ {
		if i == 0 || coords[i] != coords[i-1] {
			uniqueCoords = append(uniqueCoords, coords[i])
		}
	}

	getIdx := func(val float64) int {
		return sort.SearchFloat64s(uniqueCoords, val)
	}

	// 线段树：区间覆盖取 min，单点查询
	// 节点数 4 * len(uniqueCoords)
	size := len(uniqueCoords)
	tree := make([]float64, 4*size)
	for i := range tree {
		tree[i] = math.MaxFloat64
	}

	var update func(v, tl, tr, l, r int, val float64)
	update = func(v, tl, tr, l, r int, val float64) {
		if l > r {
			return
		}
		if l == tl && r == tr {
			if val < tree[v] {
				tree[v] = val
			}
		} else {
			tm := (tl + tr) / 2
			update(2*v, tl, tm, l, min(r, tm), val)
			update(2*v+1, tm+1, tr, max(l, tm+1), r, val)
		}
	}

	var query func(v, tl, tr, pos int) float64
	query = func(v, tl, tr, pos int) float64 {
		if tl == tr {
			return tree[v]
		}
		tm := (tl + tr) / 2
		res := tree[v]
		if pos <= tm {
			sub := query(2*v, tl, tm, pos)
			if sub < res {
				res = sub
			}
		} else {
			sub := query(2*v+1, tm+1, tr, pos)
			if sub < res {
				res = sub
			}
		}
		return res
	}

	// 插入公交区间
	for _, b := range buses {
		update(1, 0, size-1, getIdx(b.s), getIdx(b.r), b.val)
	}

	// 查询每个人
	for i := 0; i < m; i++ {
		p := peopleP[i]
		walkTime := (l - p) / y
		busTime := query(1, 0, size-1, getIdx(p))
		ans := walkTime
		if busTime < ans {
			ans = busTime
		}
		fmt.Fprintf(out, "%.10f\n", ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer out.Flush()
	solve()
}
