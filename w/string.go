package w

import (
	"strconv"
	"strings"
)

var GoString *go_string

func init() {
	GoString = &go_string{}
}

type Args_GoString_hide_keyword struct {
	Text  string
	Start int //显示文本开始的n个字符
	End   int //显示文本结束的的n个字符
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

// 18611114444 -> 186****4444
func (s *go_string) Hide_keyword(arg Args_GoString_hide_keyword) (text string) {
	size := len(arg.Text)
	if size > arg.End && size > arg.Start {
		if size > (arg.End + arg.Start) {
			xin := size - arg.End - arg.Start
			text = arg.Text[:arg.Start] + s.MushString("*", xin) + arg.Text[size-arg.End:]
			return text
		}
	}
	return s.MushString("*", size)
}

// 根据字符串生成多个
func (s *go_string) MushString(x string, n int) (text string) {
	for i := 0; i < n; i++ {
		text = text + x
	}
	return text
}

// 转in64,无论是否成功都返回int64
func (s *go_string) ToInt64V1(x string) (strInt64 int64) {
	strInt64, _ = strconv.ParseInt(x, 10, 64)
	return
}

// 转in64
func (s *go_string) ToInt64(x string) (strInt64 int64, err error) {
	strInt64, err = strconv.ParseInt(x, 10, 64)
	return
}

// 转float64
func (s *go_string) ToFloat64(x string) (strFloat64 float64, err error) {
	strFloat64, err = strconv.ParseFloat(x, 64)
	return
}
