| 日期         | 题号 | 题目名称 | 难度 | 核心知识点 | 用到的模版/技巧 | 备注 |
|:-----------| :--- | :--- | :--- | :--- | :--- | :--- |
| 2026-01-18 | [CF2035D](https://codeforces.com/problemset/problem/2035/D) | Giant's Earring | 1800 | **贪心 + 单调栈** | **快速幂 (QuickPow)**<br>栈模拟 (Stack Simulation) | 核心是将因子2全部转移到最大的奇数基底上；利用栈维护前缀最优状态，避免重复计算。 |
| 2026-01-19 | [CF2037G](https://codeforces.com/problemset/problem/2037/G) | Natlan Exploring | 1900 | **数论 DP + 容斥原理** | **欧拉筛 (Euler Sieve)** (求 SPF & $\mu$)<br>**快速分解质因数** (SPF)<br>莫比乌斯反演思路 | $O(N^2)$ 会超时，必须转化为“对因子做DP”；利用 SPF 实现 $O(\log N)$ 分解；利用 $\mu$ 系数处理重复路径。 |