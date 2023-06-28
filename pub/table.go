package pub

import (
	"fmt"
	"strings"
)

// date: 2022/01/28
// email: brach@lssin.com
var Table *table

func init() {
	Table = &table{}
}

func (t *table) AddHeader(args ...interface{}) {
	t.w = new(write)
	t.to_build(args)
}

func (t *table) AddLine(args ...interface{}) {
	t.to_build(args)
}

func (t *table) to_build(args []interface{}) {
	if len(t.w.cel) < len(args) {
		c := len(args) - len(t.w.cel)
		for _c := range Make_range(1, c) {
			_ = _c
			t.w.cel = append(t.w.cel, 0)
		}
	}
	var ar []string
	for n, i := range args {
		idx := t.to_string(i)
		if t.w.cel[n] < len(idx) {
			t.w.cel[n] = len(idx)
		}
		ar = append(ar, idx)
	}
	t.w.line = append(t.w.line, ar)
}

func (t *table) to_string(args interface{}) string {
	return fmt.Sprintf("%v", args)
}

func (t *table) to_interval(x string) string {
	format_text := ""
	for _, n := range t.w.cel {
		interval := x
		for s := range Make_range(1, n) {
			_ = s
			interval += x
		}
		format_text += interval + "\t"
	}
	return format_text
}

func (t *table) _center(x string) string {
	s := strings.Split(x, "")
	return strings.Join(s[:len(x)/2], "")
}

func (t *table) get_list(interval string) []string {
	var data []string
	for n, h := range t.w.line {
		format_text := ""
		for cel_num, cel_line := range h {
			space := ""
			space_size := t.w.cel[cel_num] - len(cel_line)
			for s := range Make_range(1, space_size) {
				_ = s
				space += " "
			}
			if n == 0 && len(space) > 1 {
				r := t._center(space)
				l := t._center(space)
				if (len(r) + len(l)) < len(space) {
					l += " "
				}
				format_text += r + cel_line + l + "\t"
			} else {
				format_text += cel_line + space + "\t"
			}
		}
		if n == 1 {
			data = append(data, strings.Trim(t.to_interval(interval), "\t"))
		}
		data = append(data, strings.Trim(format_text, "\t"))
	}
	return data
}

func (t *table) Get_text() string {
	return strings.Join(t.get_list("-"), "\n")
}

func (t *table) Print() {
	for _, line := range t.get_list("-") {
		fmt.Println(line)
	}
}
