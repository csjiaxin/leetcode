package stack

type frame struct{
	times int 
	prefix string
}

func isASCIIDigit(r char) bool {
    return r >= '0' && r <= '9'
}

func decodeString(s string) string {
	var (
		stack []frame
		times int
		current strings.Builder
	)
    stack := make([]E,0)
	times :=0
	current := ""

	for _,c in range(s){
		if isASCIIDigit(c){
			times = times*10+(c-'0')
		} else if c=="["{
			stack = append(stack, E{
				times: times,
				current: current,
			})
		} else if c=="]"{
			e := stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			for i:=1;i<times;i++{
				current +=current
			}
			current = e.current + current
		}
	}
}

// 3[a3[b]c4[d]]
// 3, 
