package main

import (
	"bufio"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// Constant & Global Variables
// ---------------------------------------------------------

const MOD = 998244353

// ---------------------------------------------------------
// Main Function
// ---------------------------------------------------------

func main() {
	// Fast I/O Setup
	// 针对大量测试用例 (T=5000)，必须使用带缓冲的 I/O
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		solve(reader, writer)
	}
}

// ---------------------------------------------------------
// Logic
// ---------------------------------------------------------

func solve(r *bufio.Reader, w *bufio.Writer) {
	var n int
	fmt.Fscan(r, &n)

	// a_0 到 a_n，共 n+1 个数
	// a[0] 是公共盒子，a[1]...a[n] 是私有盒子
	a := make([]int, n+1)
	sumS := 0
	for i := 0; i <= n; i++ {
		fmt.Fscan(r, &a[i])
		sumS += a[i]
	}

	// 核心逻辑推导：
	// 总饰品数 S = sum(a)。总轮数 = S。
	// 每个人轮流挂饰品，意味着前 r 个人会多做一轮。
	// 基础轮数 q = S / n。
	// 余数 r = S % n。
	// 也就是有 r 个位置的任务量是 q+1 (High Load)，n-r 个位置的任务量是 q (Low Load)。

	// 合法性条件：对于排列中的第 i 个位置，其对应的人 p 的私有饰品 a_p 必须 <= 该位置的任务量 L_i。
	// 证明简述：若 a_p > L_i，则该人在完成 L_i 次任务后，私有盒子仍有剩余。
	// 但题目要求“直到所有饰品都在树上”，由于总轮数固定为 S，若有人私有盒子没空，
	// 必有其他人的需求没被满足或公共盒子透支。因此必须满足 a_p <= L_i。

	q := sumS / n
	rem := sumS % n // High Load 位置的数量

	// 我们只需关注私有盒子 a[1]...a[n]
	people := a[1:]

	// 统计两类人：
	// 1. Forced: a_i > q。这些必须去 High Load (q+1) 的位置。
	//    注意：如果 a_i > q+1，则直接无解（连 High Load 都不够）。
	// 2. Free: a_i <= q。这些人既可以去 High Load 也可以去 Low Load。

	cntForced := 0
	possible := true

	for _, val := range people {
		if val > q+1 {
			possible = false
			break
		}
		if val > q {
			cntForced++
		}
	}

	if !possible {
		fmt.Fprintln(w, 0)
		return
	}

	// 如果强制去 High Load 的人数超过了 High Load 的坑位数，无解
	if cntForced > rem {
		fmt.Fprintln(w, 0)
		return
	}

	// 计算方案数：
	// 1. 从 rem 个 High Load 位置中，选出 cntForced 个位置安排那些“必须去的人”。
	//    由于人和位置都是独特的（排列不同），这是排列数 P(rem, cntForced)。
	// 2. 剩下的人 (n - cntForced) 可以任意安排在剩下的 (n - cntForced) 个位置。
	//    这是全排列 (n - cntForced)!。

	ans := int64(1)

	// P(rem, cntForced) = rem * (rem-1) * ... * (rem - cntForced + 1)
	for i := 0; i < cntForced; i++ {
		ans = (ans * int64(rem-i)) % MOD
	}

	remaining := n - cntForced
	for i := 1; i <= remaining; i++ {
		ans = (ans * int64(i)) % MOD
	}

	fmt.Fprintln(w, ans)
}
