package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

const MOD = 1_000_000_007

type Node struct {
	// originalVal = base * pow(2, exp)
	base   int // 奇数部分
	exp    int
	valMod int // base * pow(2, exp) % MOD 的值
}

// 把原来数字的因子2不断往右移动，得到和最大
func CF2035D(in io.Reader, out io.Writer) {
	var t int
	Fscan(in, &t)
	for tt := 0; tt < t; tt++ {
		var n int
		Fscan(in, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
		}
		stackTotal := 0 // "幸存者"的总和
		deadTotal := 0  // "被吸干"的总和
		stack := make([]Node, 0)
		for _, aa := range a {
			// 1. 分解aa成base和exp, 得到curNode
			cnt := 0
			temp := aa
			for temp > 0 && temp%2 == 0 {
				temp /= 2
				cnt += 1
			}
			curNode := Node{
				base:   temp,
				exp:    cnt,
				valMod: aa,
			}
			// 2. 看栈顶元素的base跟当前新的值相比，是否进行吸取
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				shouldAbsord := false
				if curNode.exp > 30 { // 必定大于aa < 10 ** 9
					shouldAbsord = true
				} else {
					// 这里需要用curNode进行判断，因为curNode会进行吸取更新
					if curNode.base*quickPow(2, curNode.exp) > top.base {
						shouldAbsord = true
					}
				}
				if shouldAbsord {
					deadTotal = (deadTotal + top.base) % MOD
					stackTotal = (stackTotal - top.valMod + MOD) % MOD
					curNode.exp += top.exp
					curNode.valMod = (curNode.base * quickPow(2, curNode.exp)) % MOD
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			// 3.新节点入栈
			stack = append(stack, curNode)
			stackTotal = (stackTotal + curNode.valMod) % MOD
			// 4.计算当前前缀最大值并输出
			final := (deadTotal + stackTotal) % MOD
			Fprintf(out, "%d ", final)
		}
		Fprintln(out, "")
	}
}

func quickPow(base, exp int) int {
	res := 1
	for exp > 0 {
		if exp%2 == 1 {
			res = (res * base) % MOD
		}
		base = (base * base) % MOD
		exp /= 2
	}
	return res
}

// bufio.NewReader坑人一手
func main() { CF2035D(bufio.NewReader(os.Stdin), os.Stdout) }
