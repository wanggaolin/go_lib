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

func (t *timeing) Current_time3() string {
	return time.Now().Format("2006-01-02 15:04")
}

// 1531293019 to 2018-07-11 15:10:19
func (t *timeing) Unix_to_beijing1(unix_time int64) string {
	tm := time.Unix(unix_time, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// 1704211140000 to 2024-01-02 23:59:00
func (t *timeing) Unix_to_beijing2(n int64) (beijingTimeString string) {
	timestamp := n / 1000 //
	tm := time.Unix(timestamp, 0)
	beijingTime := tm.In(time.FixedZone("CST", 8*60*60))
	beijingTimeString = beijingTime.Format("2006-01-02 15:04:05")
	return beijingTimeString
}

// 1735109090777 to 2024-12-25 14:44:50
func (t *timeing) UnixMilli_to_beijing1(unix_time int64) string {
	tm := time.UnixMilli(unix_time)
	return tm.Format("2006-01-02 15:04:05")
}

// 1735109090777 to 2024-12-25 14:44
func (t *timeing) UnixMilli_to_beijing2(unix_time int64) string {
	tm := time.UnixMilli(unix_time)
	return tm.Format("2006-01-02 15:04")
}

// 2018-07-11 15:10  to 1531293019
func (t *timeing) Beijing1_to_unix(beijing_time string) (ux int64, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return ux, err
	}
	tt, err := time.ParseInLocation("2006-01-02 15:04", beijing_time, loc)
	if err != nil {
		return ux, err
	}
	ux = tt.Unix()
	return ux, err
}

// 2018-07-11 15:10:19  to 1531293019
func (t *timeing) Beijing1_to_unix1(beijing_time string) (ux int64, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return ux, err
	}
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", beijing_time, loc)
	if err != nil {
		return ux, err
	}
	ux = tt.Unix()
	return ux, err
}

// 2018-07-11 15:10 to 1531293019000
func (t *timeing) Beijing1_to_unixMilli(beijing_time string) (ux int64, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return ux, err
	}
	tt, err := time.ParseInLocation("2006-01-02 15:04", beijing_time, loc)
	if err != nil {
		return ux, err
	}
	ux = tt.UnixMilli()
	return ux, err
}

// 2018-07-11 15:10:19 to 1531293019000
func (t *timeing) Beijing1_to_unixMilli1(beijing_time string) (ux int64, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return ux, err
	}
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", beijing_time, loc)
	if err != nil {
		return ux, err
	}
	ux = tt.UnixMilli()
	return ux, err
}
