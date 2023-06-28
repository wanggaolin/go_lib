package time

import "time"

// date: 2022/02/28
// email: brach@lssin.com

func Current_time1() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Current_time2() string {
	return time.Now().Format("200601021504")
}
