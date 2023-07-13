package w

import "strings"

var GoString *go_string

func init() {
	GoString = &go_string{}
}

// 按照逗号分割,返回列表 "1,2,3" => ["1","2","3"]
func (s *go_string) CommaStringFormatArry(x string) []string {
	var data []string
	for _, n := range strings.Split(x, ",") {
		n = strings.TrimSpace(strings.ReplaceAll(n, ",", ""))
		if n != "" {
			data = append(data, n)
		}
	}
	return data
}

// 按照逗号分割,重组
// "1,2,3" => "1,2,3"
// "1,2,,3" => "1,2,3"
// "1,2,,  3" => "1,2,3"
func (s *go_string) StringSpaceCommaFormat(x string) string {
	var data []string
	for _, n := range strings.Split(x, ",") {
		n = strings.TrimSpace(strings.ReplaceAll(n, ",", ""))
		if n != "" {
			data = append(data, n)
		}
	}
	return strings.Join(data, ",")
}
