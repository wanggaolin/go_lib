package w

var Set *set

func init() {
	Set = &set{}
}

// 求并集
func (s set) union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// 求交集
func (s set) IntersectInt64(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int64)
	nn := make([]int64, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// 求差集
func (s set) DifferenceInt64(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int64)
	nn := make([]int64, 0)
	inter := s.IntersectInt64(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// 列表去重
func (s set) ArryStringSet(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	if arr == nil {
		return newArr
	}
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// 列表去重
func (s set) ArryInt64Set(arr []int64) (newArr []int64) {
	newArr = make([]int64, 0)
	if arr == nil {
		return newArr
	}
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}
