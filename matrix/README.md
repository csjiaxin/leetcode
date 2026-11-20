# 矩阵与原地操作：坐标、边界与更新顺序

[返回总目录](../README.md)

这是矩阵模拟与原地变换指南，不重复已有的[网格 DP](../dp/grid/README.md)和[图 BFS](../graph/bfs/README.md)。下方为参考模板，题单不代表全部已有实现。

## 识别信号与分类

| 类型 | 思路 / 不变量 | 代表题 |
| --- | --- | --- |
| 方阵旋转 | 转置后反转每一行 | 48 旋转图像 |
| 按层走边界 | 上下左右四条边，每走完一条缩边 | 54 螺旋矩阵、59 螺旋矩阵 II |
| 原地标记 | 第一行、第一列存状态，原始状态额外保存 | 73 矩阵置零 |
| 行列都排序 | 右上角 / 左下角每次排除整行或整列 | 240 搜索二维矩阵 II |
| 同步更新 | 额外状态位保存旧值和新值，最后统一提交 | 289 生命游戏 |
| 一维旋转 | 三次反转或环状置换 | 189 轮转数组 |

## 模板一：顺时针旋转正方形矩阵

```go
// 输入必须为 n*n 方阵；原地旋转 90 度。
func Rotate(matrix [][]int) {
    n := len(matrix)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    for _, row := range matrix {
        for l, r := 0, n-1; l < r; l, r = l+1, r-1 { row[l], row[r] = row[r], row[l] }
    }
}
```

坐标映射 `(r,c)→(c,n-1-r)`，由转置 `(r,c)→(c,r)` 加行反转完成。只交换对角线一侧，否则会换回原状。时间 O(n²)、额外空间 O(1)。非方阵旋转会改变形状，不能直接照搬原地模板。

## 模板二：矩形螺旋遍历

```go
// 输入为规则矩形，不支持各行长度不同的切片。
func SpiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 { return []int{} }
    top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
    result := make([]int, 0, len(matrix)*len(matrix[0]))
    for top <= bottom && left <= right {
        for c := left; c <= right; c++ { result = append(result, matrix[top][c]) }
        top++
        for r := top; r <= bottom; r++ { result = append(result, matrix[r][right]) }
        right--
        if top <= bottom {
            for c := right; c >= left; c-- { result = append(result, matrix[bottom][c]) }
            bottom--
        }
        if left <= right {
            for r := bottom; r >= top; r-- { result = append(result, matrix[r][left]) }
            left++
        }
    }
    return result
}
```

每轮只访问尚未消费的边界；第三、第四条边必须检查边界，避免单行、单列重复输出。r 行 c 列，时间 O(rc)，除输出外额外空间 O(1)。59 将“读取”换成“按计数写入”。

## 其他高频模型的解题步骤

- **73**：先保存原第一行 / 列是否有零；扫描其余区域，把需要清零的行列标记在首列 / 首行；再清内部，最后处理第一行 / 列。太早改零会污染判断。时间 O(rc)，额外空间 O(1)。
- **240**：从右上角开始，若当前值大于 target 就左移（本列下方更大，可排除本列），否则下移（本行左侧更小，可排除本行）。时间 O(r+c)、空间 O(1)。仅各行排序不够；也不要将它当成全局一维有序矩阵。
- **289**：读取旧状态位，写入新状态位，所有邻居计算完后统一迁移。避免更新后的细胞影响同一轮其他位置。
- **189**：长度 n 先处理空数组，再令 k%=n；反转全数组、反转前 k 个、反转后 n-k 个。时间 O(n)、空间 O(1)。

## 检查与路线

先画坐标映射，再决定循环；明确闭区间 / 半开区间；检查空矩阵、单行、单列、奇数中心、更新污染，以及“原地”是否只是复用了行切片。

练习：189 → 48 → 54 → 59 → 73 → 240 → 289。
