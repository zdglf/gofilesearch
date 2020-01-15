package base_struct

import (
	"errors"
	"fmt"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

type JsonTime time.Time

func (that JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(that).Format(timeLayout))
	return []byte(stamp), nil
}

func (that *JsonTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 2 {
		return errors.New("date must be string")
	}
	loc, _ := time.LoadLocation("Asia/Chongqing")
	var t time.Time
	t, err = time.ParseInLocation(timeLayout, string(data[1:len(data)-1]), loc)
	*that = JsonTime(t)
	return
}
