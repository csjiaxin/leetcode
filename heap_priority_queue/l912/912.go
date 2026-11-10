package l912

// sortArray performs in-place heapsort. O(n log n) time, O(1) extra space.
func sortArray(nums []int) []int {
	for i:=len(nums)/2-1; i>=0; i-- { siftDown(nums,i,len(nums)) }
	for end:=len(nums)-1; end>0; end-- {
		nums[0],nums[end]=nums[end],nums[0]
		siftDown(nums,0,end)
	}
	return nums
}
func siftDown(a []int, root,end int) {
	for root<end/2 {
		child:=2*root+1
		if child+1<end && a[child+1]>a[child] { child++ }
		if a[root]>=a[child] { return }
		a[root],a[child]=a[child],a[root]
		root=child
	}
}
