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
func (r *arry) GetMaxInt(arr []int) (maxInt int) {
	if len(arr) == 0 {
		return 0
	}
	maxInt = arr[0]
	for _, num := range arr {
		if num > maxInt {
			maxInt = num
		}
	}
	return maxInt
}

func (r *arry) GetMaxInt64(arr []int64) (maxInt int64) {
	if len(arr) == 0 {
		return 0
	}
	maxInt = arr[0]
	for _, num := range arr {
		if num > maxInt {
			maxInt = num
		}
	}
	return maxInt
}

func (r *arry) GetMaxFloat64(arr []float64) (maxFloat float64) {
	if len(arr) == 0 {
		return 0
	}
	maxFloat = arr[0]
	for _, num := range arr {
		if num > maxFloat {
			maxFloat = num
		}
	}
	return maxFloat
}

func (r *arry) SumFloat64(x []float64) (sum float64) {
	for _, num := range x {
		sum += num
	}
	sum = float64(int(sum*100)) / 100
	return
}

func (r *arry) SumInt64(x []int64) (sum int64) {
	for _, num := range x {
		sum += num
	}
	return
}
