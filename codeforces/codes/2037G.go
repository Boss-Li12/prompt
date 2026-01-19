package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 1.基础配置和全局变量
const (
	MAX_V = 1000005
	MOD   = 998244353
)

var (
	spf [MAX_V]int // 最小质因子, 用来分解质因数
	mu  [MAX_V]int // 使用容斥原理
	// DP全局状态数组
	cnt [MAX_V]int // cnt[x]: 目前为止，所有能被x整除的数的DP之和
)

// 2.线性筛 -> 得到所有mu和spf
func initMath() {
	var primes []int
	mu[1] = 1
	for i := 2; i < MAX_V; i++ {
		// 假设是质数
		if spf[i] == 0 {
			spf[i] = i
			primes = append(primes, i)
			mu[i] = -1 // 质数的mu是-1
		}
		// 筛去合数
		for _, p := range primes {
			if p > spf[i] || i*p >= MAX_V {
				break
			}
			spf[i*p] = p // 记录最小因子
			if i%p == 0 {
				mu[i*p] = 0 // 含平方因子，不用进行容斥
			} else {
				mu[i*p] = -mu[i] // 每增加一个质因子符号变一次
			}
		}
	}
}

// 3.得到所有约数
func getDivsors(x int) []int {
	// 1.分解质因数
	var pFactors []int
	temp := x
	for temp > 1 {
		p := spf[temp]
		pFactors = append(pFactors, p)
		for temp%p == 0 {
			temp /= p
		}
	}
	divs := []int{1}
	for _, p := range pFactors {
		k := len(divs)
		for i := 0; i < k; i++ {
			divs = append(divs, p*divs[i])
		}
	}
	return divs
}

// 记录之前所有数累积到因子中的和，需要使用莫比乌斯反演进行去重
// 直觉公式 dp[i] = sigma{j}(dp[j]), 平方复杂度超时，改用因子法
// 核心公式: ans = sigma{d} (-mu[d] * cnt[d])
func CF2037G(in io.Reader, out io.Writer) {
	initMath()
	t := 1
	for tt := 0; tt < t; tt++ {
		var n int
		Fscan(in, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
		}
		// 从因子得到现在的DP，把现在的DP加回到因子中
		for i, aa := range a {
			divs := getDivsors(aa)
			currentDP := 0
			if i == 0 {
				currentDP = 1 // 第一个就是1
			} else {
				for _, div := range divs {
					if div == 1 {
						continue
					}
					tempCnt := cnt[div]
					if mu[div] == -1 { // 容斥原理对应+
						currentDP = (currentDP + tempCnt) % MOD
					} else { // 容斥原理对应-
						currentDP = (currentDP - tempCnt + MOD) % MOD
					}
				}
			}
			// 更新对应的cnt
			for _, div := range divs {
				cnt[div] = (cnt[div] + currentDP) % MOD
			}
			if i == n-1 {
				Fprintln(out, currentDP)
			}
		}

	}
}

// bufio.NewReader坑人一手
func main() { CF2037G(bufio.NewReader(os.Stdin), os.Stdout) }
