package utils

import "time"

var local *time.Location

func init() {
	local, _ = time.LoadLocation("Asia/Shanghai")
}

func ParseDate(oldTime string) (time.Time, error) {
	retTime, err := time.ParseInLocation("2006-01-02", oldTime, local)
	return retTime, err
}

func ParseDateAndTime(oldTime string) (time.Time, error) {
	retTime, err := time.ParseInLocation("2006-01-02 15:04:05", oldTime, local)
	return retTime, err
}

func RParseDate(oldTime time.Time) string {
	return oldTime.Format("2006-01-02")
}

func RParseDateAndTime(oldTime time.Time) string {
	return oldTime.Format("2006-01-02 15:04:05")
}
