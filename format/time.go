package format

import "time"

// TimeToString 时间转字符串 默认 Format "2006-01-02 15:04:05"
func TimeToString(time time.Time, format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return time.Format(format)
}
