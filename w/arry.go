package w

var (
	Arry *arry
)

func init() {
	Arry = &arry{}
}

// 判断列表是否包含字符串
func (r *arry) Is_exis_arry_string(name string, atty []string) bool {
	for _, uname := range atty {
		if name == uname {
			return true
		}
	}
	return false
}

func (r *arry) GetMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}
