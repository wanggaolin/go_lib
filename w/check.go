package w

import (
	"regexp"
	"strconv"
	"strings"
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

func (c *check) Check_ip_private(ip_name string) bool {
	if ip_name == "0.0.0.0" {
		return false
	}
	re_ip1 := regexp.MustCompile(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`)
	re_ip2 := regexp.MustCompile(`^10\.|^192\.168\.`)
	ip_name = strings.Split(ip_name, "/")[0]
	if re_ip1.FindAllStringSubmatch(ip_name, -1) != nil {
		if re_ip2.FindAllStringSubmatch(ip_name, -1) != nil {
			return true
		}
		ip_str := strings.Split(ip_name, ".")
		if ip_str[0] == "172" {
			_ip1, _ := strconv.ParseInt(ip_str[1], 10, 64)
			if 16 <= _ip1 && _ip1 <= 32 {
				return true
			}
		}
	}
	return false
}

func (c *check) Check_int(x string) bool {
	_, err := strconv.ParseInt(x, 10, 64)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (c *check) Check_float(x string) bool {
	_, err := strconv.ParseFloat(x, 64)
	if err == nil {
		return true
	} else {
		return false
	}
}
