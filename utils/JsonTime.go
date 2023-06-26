package utils

import (
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type JsonTime time.Time

// JsonDate反序列化
func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	newTime, err := time.ParseInLocation("\""+timeFormat+"\"", string(data), time.Local)
	*t = JsonTime(newTime)
	return
}

// JsonDate序列化
func (t JsonTime) MarshalJSON() ([]byte, error) {
	timeStr := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return []byte(timeStr), nil
}

// string方法
func (t JsonTime) String() string {
	return time.Time(t).Format(timeFormat)
}
