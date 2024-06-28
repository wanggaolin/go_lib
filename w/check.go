package w

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
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
		a1, _ := strconv.Atoi(filter_ip[1])
		b1, _ := strconv.Atoi(filter_ip[2])
		c1, _ := strconv.Atoi(filter_ip[3])
		d1, _ := strconv.Atoi(filter_ip[4])
		if a1 <= 255 && b1 <= 255 && c1 <= 255 && d1 <= 255 && a1 > 0 && d1 > 0 {
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

func (c *check) CheckTimeDay(x string) bool {
	_, err := time.Parse("2006-01-02", x)
	if err != nil {
		return false
	}
	return true
}

func (c *check) CheckTimeBeijing1(x string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", x)
	if err != nil {
		return false
	}
	return true
}

func (c *check) CheckTimeBeijing2(x string) bool {
	_, err := time.Parse("2006-01-02 15:04", x)
	if err != nil {
		return false
	}
	return true
}

func Valid(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		for _, item := range strings.Split(t.Field(i).Tag.Get("validate"), ",") {
			e := strings.Split(item, ":")
			if len(e) > 1 {
				fmt.Println(e)
			}
		}
	}
}
