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

// 判断列表是否包含数字
func (r *arry) Is_exis_arry_int64(name int64, atty []int64) bool {
	for _, uname := range atty {
		if name == uname {
			return true
		}
	}
	return false
}

// 判断列表是否包含数字
func (r *arry) Is_exis_arry_int(name int, atty []int) bool {
	for _, uname := range atty {
		if name == uname {
			return true
		}
	}
	return false
}

// 获取列表最大
func (r *arry) GetMaxInt(arr []int) int {
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

func (r *arry) GetMaxInt64(arr []int64) int64 {
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
