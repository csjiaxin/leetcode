package l1901

// 1901：复用本文件的边界模板；特殊结构按对应排除不变量处理。
func findPeakGrid(mat [][]int) []int {
	left, right := 0, len(mat[0])-1
	for left <= right {
		column := left + (right-left)/2
		row := 0
		for r := 1; r < len(mat); r++ {
			if mat[r][column] > mat[row][column] {
				row = r
			}
		}
		if column > 0 && mat[row][column-1] > mat[row][column] {
			right = column - 1
		} else if column+1 < len(mat[0]) && mat[row][column+1] > mat[row][column] {
			left = column + 1
		} else {
			return []int{row, column}
		}
	}
	return nil
}

func FirstTrue(n int, predicate func(int) bool) int {
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if predicate(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
