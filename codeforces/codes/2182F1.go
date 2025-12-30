package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MOD      = 998244353
	MAX_N    = 1005
	MAX_BITS = 62 // 驯鹿最大强度为 2^60
)

var (
	C      [MAX_N + 5][MAX_N + 5]int64
	waysEq [MAX_N + 5]int64
	waysGt [MAX_N + 5]int64
	nextEq [MAX_N + 5]int64
	nextGt [MAX_N + 5]int64
)

// 预处理组合数
func init() {
	for i := 0; i < MAX_N+5; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var initialN, m int
	if _, err := fmt.Fscan(reader, &initialN, &m); err != nil {
		return
	}

	counts := make([]int, MAX_BITS+1)
	currentN := initialN
	for i := 0; i < initialN; i++ {
		var val int
		fmt.Fscan(reader, &val)
		if val <= MAX_BITS {
			counts[val]++
		}
	}

	for q := 0; q < m; q++ {
		var op int
		fmt.Fscan(reader, &op)

		if op == 1 {
			var x int
			fmt.Fscan(reader, &x)
			counts[x]++
			currentN++
		} else if op == 2 {
			var x int
			fmt.Fscan(reader, &x)
			counts[x]--
			currentN--
		} else {
			var targetX uint64 // 使用 uint64 处理 10^18 级别的 x
			fmt.Fscan(reader, &targetX)

			// 初始化 DP 状态
			for i := 0; i <= currentN; i++ {
				waysEq[i] = 0
				waysGt[i] = 0
			}
			waysEq[0] = 1

			// 从高位到低位逐层处理贡献
			for i := MAX_BITS; i >= 0; i-- {
				cnt := counts[i]
				// 清空缓冲区
				for idx := 0; idx <= currentN; idx++ {
					nextEq[idx] = 0
					nextGt[idx] = 0
				}

				for j := 0; j <= currentN; j++ {
					if waysEq[j] == 0 && waysGt[j] == 0 {
						continue
					}

					// 1. 检查层间 Gap 位：gapBit = i - j + 1
					// 如果 targetX 在该位为 1，而此位只能填 0，则该路径不再满足 Eq 状态
					gapBit := i - j + 1
					isEqPossible := true
					if waysEq[j] > 0 {
						if gapBit >= 0 && gapBit < 64 {
							if (targetX>>uint(gapBit))&1 == 1 {
								isEqPossible = false
							}
						}
					}

					// 2. 枚举当前强度 2^i 选多少个
					for k := 0; k <= cnt; k++ {
						newTotal := j + k
						if newTotal > currentN {
							break
						}

						combinationWays := C[cnt][k]

						// 转移 Gt 状态：只要之前已经大于，现在怎么选都大于
						if waysGt[j] > 0 {
							nextGt[newTotal] = (nextGt[newTotal] + waysGt[j]*combinationWays) % MOD
						}

						// 转移 Eq 状态
						if waysEq[j] > 0 && isEqPossible {
							high := i - j
							low := i - j - k + 1

							status := 0 // 0: Equal, 1: Greater, -1: Smaller
							for b := high; b >= low; b-- {
								if b < 0 {
									break
								} // 负数位不贡献，且 x 的负数位视为 0

								// 我们的位是 1 (因为是 2^i 贡献的段)
								xBit := 0
								if b < 64 && (targetX>>uint(b))&1 == 1 {
									xBit = 1
								}

								if 1 > xBit {
									status = 1
									break
								} else if 1 < xBit { // 此分支理论上不进入，保留逻辑完整性
									status = -1
									break
								}
							}

							if status == 1 {
								nextGt[newTotal] = (nextGt[newTotal] + waysEq[j]*combinationWays) % MOD
							} else if status == 0 {
								nextEq[newTotal] = (nextEq[newTotal] + waysEq[j]*combinationWays) % MOD
							}
						}
					}
				}
				// 状态滚动
				for idx := 0; idx <= currentN; idx++ {
					waysEq[idx] = nextEq[idx]
					waysGt[idx] = nextGt[idx]
				}
			}

			// 最终统计
			var ans int64 = 0
			for i := 0; i <= currentN; i++ {
				ans = (ans + waysEq[i] + waysGt[i]) % MOD
			}
			fmt.Fprintln(writer, ans)
		}
	}
}
