package data

import (
	"fmt"
	"time"
)

// 创建时间
// t: 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
// now:  2009-11-10 23:00:00 +0000 UTC m=+0.000000001
// n: 2025-01-27 10:10:10.00000001 +0000 UTC
func now() {
	// 获取当前时间
	t := time.Now()
	fmt.Printf("t: %v\n", t)
	fmt.Println("now: ", t)

	// 创建特定的时间
	n := time.Date(2025, time.January, 27, 10, 10, 10, 10, time.Local)
	fmt.Printf("n: %v\n", n)
}

// 时间打印
// format1: 2025-01-27 10:10:10
// format2: 2025-01-27T10:10:10Z
// format3: 2025年01月27日 10时10分
func formatTime() {
	n := time.Date(2025, time.January, 27, 10, 10, 10, 10, time.Local)
	format1 := n.Format("2006-01-02 15:04:05")
	fmt.Printf("format1: %v\n", format1)

	format2 := n.Format(time.RFC3339) // 使用预定义格式
	fmt.Printf("format2: %v\n", format2)

	// 自定义格式
	format3 := n.Format("2006年01月02日 15时04分")
	fmt.Printf("format3: %v\n", format3)
}

// 时间计算
// format1: 2025-01-28 10:10:10
// format2: 2025-01-28 10:10:10
// add24Hour equal add1Day : true
// future sub now : 24h0m0s
// Duration hours: 24, hoursF: 24
// Duration minutes: 1440, minutesF: 1440
// Duration seconds: 86400, secondsF: 86400
// Duration : 24:00:00
func calTime() {
	n := time.Date(2025, time.January, 27, 10, 10, 10, 10, time.Local)

	add24Hour := n.Add(24 * time.Hour)
	format1 := add24Hour.Format("2006-01-02 15:04:05")
	fmt.Printf("format1: %v\n", format1)

	add1Day := n.AddDate(0, 0, 1)
	format2 := add24Hour.Format("2006-01-02 15:04:05")
	fmt.Printf("format2: %v\n", format2)

	equal := add24Hour.Equal(add1Day)
	fmt.Printf("add24Hour equal add1Day : %v\n", equal)

	// 时间比较
	future := time.Date(2025, time.January, 28, 10, 10, 10, 10, time.Local)
	duration := future.Sub(n)
	fmt.Printf("future sub now : %v\n", duration)

	// 转换为小时
	hours := int(duration.Hours()) // 总小时数
	hoursF := duration.Hours()     // 带小数的小时数
	fmt.Printf("Duration hours: %v, hoursF: %v\n", hours, hoursF)

	// 转换为分钟
	minutes := int(duration.Minutes()) // 总分钟数
	minutesF := duration.Minutes()     // 带小数的分钟数
	fmt.Printf("Duration minutes: %v, minutesF: %v\n", minutes, minutesF)

	// 转换为秒
	seconds := int(duration.Seconds()) // 总秒数
	secondsF := duration.Seconds()     // 带小数的秒数
	fmt.Printf("Duration seconds: %v, secondsF: %v\n", seconds, secondsF)

	// 分别获取时分秒
	timeStr := formatDuration(duration)
	fmt.Printf("Duration : %v\n", timeStr)
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// 时区转换
// 2025-01-26 21:10:10 EST
// 2025-01-27 11:10:10 JST
// 2025-01-27 10:10:10 CST
// 2025-01-27 13:10:10 AEDT
// 2025-01-27 02:10:10 UTC
func timezoneConvert() {
	// 加载时区
	nyc, _ := time.LoadLocation("America/New_York")
	tokyo, _ := time.LoadLocation("Asia/Tokyo")
	beijing, _ := time.LoadLocation("Asia/Shanghai")
	canberra, _ := time.LoadLocation("Australia/Canberra")

	now := time.Date(2025, time.January, 27, 10, 10, 10, 10, beijing)

	// 转换时区
	nycTime := now.In(nyc)
	tokyoTime := now.In(tokyo)
	beijingTime := now.In(beijing)
	canberraTime := now.In(canberra)

	fmt.Println(nycTime.Format("2006-01-02 15:04:05 MST"))
	fmt.Println(tokyoTime.Format("2006-01-02 15:04:05 MST"))
	fmt.Println(beijingTime.Format("2006-01-02 15:04:05 MST"))
	fmt.Println(canberraTime.Format("2006-01-02 15:04:05 MST"))

	// UTC 转换
	utcTime := now.UTC()
	fmt.Println(utcTime.Format("2006-01-02 15:04:05 MST"))
}

// 字符串转Time
// 2025-01-27 10:30:00 UTC
// 2025-01-27 10:30:00 CST
// 2024-01-24 10:30:00 UTC
func strToTime() {
	// 解析标准格式
	t1, _ := time.Parse("2006-01-02 15:04:05", "2025-01-27 10:30:00")
	fmt.Println(t1.Format("2006-01-02 15:04:05 MST"))

	// 解析带时区的时间
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-01-27 10:30:00", loc)
	fmt.Println(t2.Format("2006-01-02 15:04:05 MST"))

	// 解析时间戳
	timestamp := int64(1706092200)
	t3 := time.Unix(timestamp, 0)
	fmt.Println(t3.Format("2006-01-02 15:04:05 MST"))
}
