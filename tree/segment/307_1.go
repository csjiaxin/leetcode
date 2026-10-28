package segment

// #segment tree implementation
// 线段树把数组不断分成左右两部分。每个父节点保存左右子节点合并后的结果：
type NumArray struct {
	size int
	tree []int
}

func Constructor(nums []int) NumArray {
	size := 1
	for size < len(nums) {
		size *= 2
	}
	s := NumArray{
		size: size,
		tree: make([]int, 2*size),
	}
	copy(s.tree[size:], nums)
	// 从叶子向上计算父节点。
	for i := size - 1; i > 0; i-- {
		s.tree[i] = s.tree[i*2] + s.tree[i*2+1]
	}
	return s
}

func (this *NumArray) Update(index int, val int) {
	indexInTree := this.size + index
	this.tree[indexInTree] = val
	// 从叶子向上计算父节点。
	for i := indexInTree / 2; i > 0; i /= 2 {
		this.tree[i] = this.tree[i*2] + this.tree[i*2+1]
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	left += this.size
	right += this.size + 1 // 转为左闭右开区间 [left, right)
	var sum int
	for left < right {
		if left%2 == 1 {
			sum += this.tree[left]
			left += 1
		}
		if right%2 == 1 {
			right--
			sum += this.tree[right]
		}
		left /= 2
		right /= 2
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
