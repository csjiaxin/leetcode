# 扫描线与差分：统一指南入口

[返回总目录](../README.md)

扫描线内容已与差分指南合并到 **[差分与扫描线：端点变化、前缀恢复与活跃状态](../prefix_sum/difference/README.md)**，避免端点规则、事件解释和题单重复维护。本目录保留为原有链接的入口，相关源码位置不变。

## 按问题跳转

| 想了解什么 | 统一指南入口 |
| --- | --- |
| 都能解决什么问题？差分和扫描线有什么区别？ | [关系与选型](../prefix_sum/difference/README.md#relationship) |
| 什么是事件、状态、答案？ | [事件与状态及逐步例子](../prefix_sum/difference/README.md#events) |
| 为什么只改两个端点，前缀和就能恢复？ | [差分推导](../prefix_sum/difference/README.md#derivation) |
| 原数组不是零，应该怎么处理？ | [完整差分与增量差分](../prefix_sum/difference/README.md#initial-array) |
| 区间更新的 Go 模板 | [泛型差分模板](../prefix_sum/difference/README.md#template) |
| 坐标巨大，不能创建数组怎么办？ | [稀疏事件扫描与坐标压缩](../prefix_sum/difference/README.md#sparse-sweep) |
| 闭区间、半开区间，同坐标谁先处理？ | [边界约定](../prefix_sum/difference/README.md#endpoint-rules) |
| 如何处理天际线、最短区间和在线预订？ | [进阶扫描线](../prefix_sum/difference/README.md#advanced-sweep) |
| 接下来练哪些题？ | [分类题单](../prefix_sum/difference/README.md#extended-practice) · [练习顺序](../prefix_sum/difference/README.md#practice-order) |

差分是记录变化的方法；扫描线是按位置推进、处理事件并维护状态的框架。可加总量的端点扫描与差分关系紧密，但建筑最大高度等问题还需要堆、有序集合或线段树，不能统一简化为一个累计和。
