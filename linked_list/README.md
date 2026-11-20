# 链表：指针变换、区间连接与节点身份

[返回总目录](../README.md) · [已有快慢指针](../two_pointers/readme.md)

本目录补齐链表操作学习指南；以下为参考模板，不批量创建空题解。已有环、回文与中点题仍保留在双指针目录。

## 识别信号与分类

| 类型 | 通用分解 | 代表题 |
| --- | --- | --- |
| 反转 | 保存后继 → 改指向 → 推进 | 206、92 |
| 合并 | dummy＋尾指针，连接较小头节点 | 21、23 |
| 分组变换 | 找齐 k 个节点，明确组前、组尾和组后连接 | 25 |
| 重排 | 中点拆分＋反转后半＋交替合并 | 143 |
| 相交 | 比较节点指针身份，切换两条路径抵消长度差 | 160 |
| 深拷贝 | 旧指针→新节点映射，再补全边 | 138 |
| 排序 | 拆半、递归排序、合并；进阶自底向上 | 148 |

## 模板一：整体反转

```go
type ListNode struct { Val int; Next *ListNode }

func reverseList(head *ListNode) *ListNode {
    var previous *ListNode
    for head != nil {
        next := head.Next
        head.Next = previous
        previous = head
        head = next
    }
    return previous
}
```

不变量：previous 是已反转的前缀，head 是尚未处理后缀的头；修改 Next 前必须保存后继。时间 O(n)，额外空间 O(1)。会修改输入连接，返回的是新头。

## 模板二：合并两个有序链表

```go
// 两条输入链表升序、无环且节点互不共享；原地复用节点。
func mergeTwoLists(a, b *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy
    for a != nil && b != nil {
        if a.Val <= b.Val {
            tail.Next = a; a = a.Next
        } else {
            tail.Next = b; b = b.Next
        }
        tail = tail.Next
    }
    if a != nil { tail.Next = a } else { tail.Next = b }
    return dummy.Next
}
```

tail 指向已选有序前缀的尾部。时间 O(n+m)，额外空间 O(1)。输入共享尾节点时不能直接照搬，可能连成自环。23 还可用[最小堆](../heap_priority_queue/l23/23.go)。

## 分组、复制与排序的通用步骤

- **25**：用 dummy 保存组前节点；先探测第 k 个节点，不足 k 个保持原样；保存组后节点，反转当前组，再连接新组头与新组尾。时间 O(n)，迭代辅助空间 O(1)。
- **92**：只反转 `[left,right]`，重点是接回区间前驱与原区间头（新尾），不是丢弃区间外节点。
- **138**：第一遍为每个旧节点分配新节点，第二遍根据映射设置 Next、Random；平均时间 O(n)、空间 O(n)。穿插法可降低辅助空间，但会暂时修改原链，必须完整恢复。
- **148**：递归归并排序时间 O(n log n)、调用栈 O(log n)；要求常数辅助空间时考虑自底向上归并。

## 固定检查清单

1. 节点值相等不代表同一个节点：160、环检测和映射必须比较指针身份。
2. 先保存后继，再断开或改写指针；画清“组前、组头、组尾、组后”。
3. 修改头节点的操作优先考虑 dummy，返回 dummy.Next 而非 dummy。
4. 每次循环检查是否推进；最终尾节点是否正确接回、是否意外成环。
5. 空链、单节点、重复值、恰好 k 个、不足 k 个均要测试。
6. 递归反转不是 O(1) 空间；链表排序与复制是否修改输入要写清。

练习：206 → 21 → 92 → 160 → 143 → 25 → 138 → 148。
