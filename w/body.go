package w

import (
	"regexp"
	"sync"
)

type loging struct {
}

type fileing struct {
}

type OS struct {
}

type System struct {
}

type hash struct {
}

type timeing struct {
}

type set struct {
}

type table struct {
	w            *write
	lineInterval string // 行间隔符
}

type write struct {
	cel  []int
	line [][]string
}

type check struct {
	check_ip *regexp.Regexp
}

type go_string struct {
}

type tea struct {
}

type arry struct {
}

type w_json struct {
}

type lock struct {
	lock_map sync.Map
}

type xlsx struct {
}

type webhook struct {
}
