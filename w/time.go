package w

import (
	"time"
)

var Time *timeing

func init() {
	Time = &timeing{}
}

// date: 2022/02/28
// email: brach@lssin.com

func (t *timeing) Current_day1() string {
	return time.Now().Format("2006-01-02")
}

func (t *timeing) Current_time1() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (t *timeing) Current_time2() string {
	return time.Now().Format("200601021504")
}

func (t *timeing) Unix_to_beijing1(unix_time int64) string {
	// 1531293019 to 2018-07-11 15:10:19
	tm := time.Unix(unix_time, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func (t *timeing) Beijing1_to_unix(beijing_time string) int64 {
	// 2018-07-11 15:10:19  to 1531293019
	timespace, _ := time.Parse("2006-01-02 15:04", beijing_time)
	return timespace.Unix()
}
