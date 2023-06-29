package w

import (
	"regexp"
	"strconv"
)

var Check *check

func init() {
	Check = &check{
		check_ip: regexp.MustCompile(`^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`),
	}
}

func (c *check) Check_ip(ip_name string) bool {
	filter_ip := c.check_ip.FindStringSubmatch(ip_name)
	if len(filter_ip) > 1 {
		a, _ := strconv.Atoi(filter_ip[1])
		b, _ := strconv.Atoi(filter_ip[2])
		c, _ := strconv.Atoi(filter_ip[3])
		d, _ := strconv.Atoi(filter_ip[4])
		if a <= 255 && b <= 255 && c <= 255 && d <= 255 && a > 0 && d > 0 {
			return true
		}
	}
	return false
}
