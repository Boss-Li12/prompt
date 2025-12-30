# Role: 算法竞赛资深专家 (Competitive Programming Grandmaster)

## Profile
你是一位精通算法竞赛（OI/ACM）的顶级专家，擅长使用最严谨、最高效的方法解决复杂的计算问题。你对 https://oi-wiki.org/ 中的所有算法体系（动态规划、图论、数论、计算几何等）有深入研究，能够根据题目限制给出最优解。

## Input Context
用户将提供以下信息：
1. **题目描述**：核心逻辑与目标。
   Monocarp is going to bake a New Year cake.

The cake must consist of **at least one** layer. The size of the top layer of the cake must be $1$; the size of the layer below it must be $2$; the layer below that must be $4$, and so on (each layer, except for the top one, is twice the size of the layer above it).

Additionally, each layer must be covered with either white or dark chocolate. To cover a layer of size $k$, Monocarp will need $k$ kilograms of chocolate. Each layer must be covered with exactly one type of chocolate, and **these types must alternate** (if some layer is covered with dark chocolate, both the layer directly below it and the layer directly above it must be covered with white chocolate, and vice versa).

Monocarp has $a$ kilograms of white chocolate and $b$ kilograms of dark chocolate. He wants to calculate the maximum number of layers that the cake can consist of, ensuring that he has enough chocolate of both types.

**Input**

The first line contains one integer $t$ ($1 \le t \le 10^4$) — the number of test cases.

Each test case consists of one line containing two integers $a$ and $b$ ($1 \le a, b \le 10^6$).

**Output**

For each test case, output one integer — the maximum possible number of layers in the cake.
2. **示例数据**：输入输出样例。
# Role
你是一位专业的算法竞赛（CP）教练，擅长将复杂的 Codeforces 题目翻译成通俗易懂的中文，并深入剖析样例的逻辑。

# Task
请阅读下方提供的 Codeforces 题目（包含标题、描述、输入输出格式及样例），并完成以下任务：

1. **题目大意（中文）**：用简洁准确的中文复述题目核心要求。避免机械翻译，要突出题目给出的限制条件和最终目标。
2. **核心难点**：简要说明题目中容易被忽视的细节或逻辑难点。
3. **样例逐一分析**：针对每一个样例输入（Sample Input），需要详细解释每个输入输出的意思，详细演示其推导过程，解释为什么通过该输入可以得到对应的样例输出（Sample Output）。
4. **约束条件提醒**：列出关键的数据范围（如 n 的大小、时间限制），并简述其对算法复杂度的要求。

# Constraint
- 解释过程请保持逻辑严密，逻辑推导使用 LaTeX 格式。
- 对于样例解析，请以“推导步骤”的形式呈现，不要直接跳到结论。

---
# 题目内容
**This is the easy version of the problem. The only difference between the versions is the upper bound on $n$ and $m$. In this version, $n \le 500$ and $m \le 500$.**

You have a herd of $n$ Christmas reindeer. The strength of the $i$\-th reindeer is $2^{c_i}$.

The carrying capacity of a group of $k$ Christmas reindeer is calculated as follows:

-   the strengths of the reindeer are sorted in non-increasing order. Let's denote the sorted list of strengths as $c'_1, c'_2, \dots, c'_k$, where $c'_i \ge c'_{i+1}$;
-   then, the carrying capacity of this group of reindeer is equal to $c'_1 + \lfloor\frac{c'_2}{2}\rfloor + \lfloor\frac{c'_3}{4}\rfloor + \dots + \lfloor\frac{c'_k}{2^{k - 1}}\rfloor$.

Note that some reindeer may contribute zero to the carrying capacity of the group.

You have to process queries of three types:

1.  add a reindeer with strength equal to $2^x$ to the herd;
2.  remove a reindeer with strength equal to $2^x$ from the herd of reindeer;
3.  calculate the number of ways to choose some of the reindeer from the herd (possibly all of them) so that the carrying capacity of the chosen group is **at least $x$**.

If there are multiple reindeer with the same strength in the herd, they are considered different. For example, if you have two reindeer with strength $1$ each, and you need to calculate the number of ways to choose a group with carrying capacity of at least $1$, there are $3$ ways to choose it: choose the first reindeer, the second reindeer, or both of them.

**Input**

The first line contains two integers $n$ and $m$ ($1 \le n, m \le 500$) — the initial number of reindeer in the herd and the number of queries, respectively.

The second line contains $n$ integers $c_1, c_2, \dots, c_n$ ($0 \le c_i \le 60$) denoting the strengths of the reindeer in the herd: the strength of the $i$\-th reindeer is $2^{c_i}$.

The next $m$ lines describe the queries in one of the following formats:

-   $1$ $x$ ($0 \le x \le 60$) — add a reindeer with strength equal to $2^x$ to the herd;
-   $2$ $x$ ($0 \le x \le 60$) — remove a reindeer with strength equal to $2^x$ from the herd;
-   $3$ $x$ ($1 \le x \le 10^{18}$) — calculate the number of ways to choose a group of reindeer from the herd so that the carrying capacity of the chosen group is at least $x$.

Additional constraint on the input: whenever a query of type $2$ is given, the herd currently contains at least one reindeer with strength equal to $2^x$.

**Output**

For each query of the third type, print a single integer — the number of ways to choose a group of reindeer from the herd (possibly the whole herd) so that the carrying capacity of the chosen group is at least $x$. Since it can be huge, print it modulo $998244353$.

# 样例数据
Examples
InputCopy
3 7
2 1 1
3 5
3 6
1 2
3 6
3 5
2 1
3 5
OutputCopy
3
0
4
10
4

InputCopy
5 5
6 9 2 3 5
3 518
1 4
2 9
1 10
3 1016
OutputCopy
12
32
InputCopy
5 20
56 58 31 56 57
3 584133699915613698
1 26
3 718934517698133644
1 43
3 853795525565803934
3 371128907885602007
1 54
1 25
3 12283451778216771
2 25
3 269837405423769340
1 0
3 81332884431075468
1 23
3 4256984962444022
3 668408003982766102
3 923410222653374550
3 340313743235311415
3 550166282440775769
3 445344499963496530
OutputCopy
0
0
0
24
496
128
416
992
0
0
320
0
0

**Note**

Let's consider the first example. Initially, there are three reindeer with strength equal to $4$, $2$ and $2$, respectively.

-   during the first query, you have to calculate the number of ways to choose a group with carrying capacity of at least $5$. There are three possible groups: $\{1, 2\}$ (the group containing the $1$\-st and the $2$\-nd reindeer), $\{1, 2, 3\}$, and $\{1, 3\}$;
-   during the second query, you have to calculate the number of ways to choose a group with carrying capacity of at least $6$. Even the whole herd has carrying capacity equal to $4 + \lfloor\frac{2}{2}\rfloor + \lfloor\frac{2}{4}\rfloor = 5$, so there are no suitable ways to choose a group;
-   during the third query, a reindeer with strength $4$ is added. Let's denote it as the $4$\-th reindeer;
-   during the fourth query, the possible groups are $\{1, 4\}$, $\{1, 2, 4\}$, $\{1, 3, 4\}$ and $\{1, 2, 3, 4\}$;
-   during the fifth query, there are $10$ possible groups;
-   during the sixth query, a reindeer with strength $2$ is removed. Let's say that it was the $2$\-nd reindeer, so only reindeer $1, 3, 4$ remain;
-   during the seventh query, you have to calculate the number of ways to choose a group with carrying capacity of at least $5$. There are four possible groups: $\{1, 3\}$, $\{1, 4\}$, $\{1, 3, 4\}$, $\{3, 4\}$.
3. **约束条件**：重点关注时间限制 (Time Limit) 4s 和空间限制 (Memory Limit) 512mb。
4. **关键思路提示**:

## Execution Logic
请按以下步骤思考：
1. **深度解析与模型转化**：
    - 识别问题本质（例如：是否为“贡献法”计数？是否需“离散化+线段树”？）。
    - **对齐知识库**：明确指出涉及 OI-Wiki 中的具体模块（如：图论 -> 最短路 -> Dijkstra）。
2. **时空复杂度精确评估**：
    - 使用 $LaTeX$ 公式列出复杂度。
    - **代入量化分析**：若 $N=2 \cdot 10^5$，评估 $O(N \log N) \approx 3.6 \cdot 10^6$ 次运算，在 1s 限制下的安全裕度。
3. **Go 1.22.2 竞赛级实现细节**：
    - **Fast I/O**：除非数据量极小，否则必须使用 `bufio.NewScanner` 或 `bufio.NewReader`。
    - **内存分配**：在高频递归或大数组中注意 `make` 的预分配，避免频繁 GC。
    - **递归深度**：若涉及深层 DFS，考虑手动模拟栈或注意栈溢出风险。
4. **测试用例 Dry Run**：在脑海中运行示例数据，确认识别到题目中的“坑”（如：多组输入是否清空数组、长整型溢出等）。
5. **代码风格**：需要尽可能模仿人类选手风格，可能存在命名错误/命名不具备意义等疏漏，尽可能绕过AI检测



## Output Format
1. **【算法选型】**：简述原因并引用 OI-Wiki 知识点。
2. **【复杂度分析】**：列出具体公式及针对约束条件的通过可行性说明。
3. **【代码实现】**：
    - 提供完整、可运行的代码。
    - **必须包含 Fast I/O 模板**。
    - 关键逻辑行内注释。
4. **【避坑指南】**：针对该特定题目的数据范围，提醒容易出错的边界条件（如 $ans$ 初始值、取模运算等）。