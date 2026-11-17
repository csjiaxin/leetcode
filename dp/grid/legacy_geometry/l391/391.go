package main

// 完整覆盖要求面积相等，且奇数次角点只有外接矩形的四角。
func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) == 0 {
		return false
	}
	left, bottom, right, top := rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	area := int64(0)
	corners := make(map[[2]int]bool)
	for _, rectangle := range rectangles {
		x1, y1, x2, y2 := rectangle[0], rectangle[1], rectangle[2], rectangle[3]
		if x1 >= x2 || y1 >= y2 {
			return false
		}
		left, bottom = min(left, x1), min(bottom, y1)
		right, top = max(right, x2), max(top, y2)
		area += int64(x2-x1) * int64(y2-y1)
		for _, corner := range [][2]int{{x1, y1}, {x1, y2}, {x2, y1}, {x2, y2}} {
			if corners[corner] {
				delete(corners, corner)
			} else {
				corners[corner] = true
			}
		}
	}
	return area == int64(right-left)*int64(top-bottom) && len(corners) == 4 &&
		corners[[2]int{left, bottom}] && corners[[2]int{left, top}] &&
		corners[[2]int{right, bottom}] && corners[[2]int{right, top}]
}
