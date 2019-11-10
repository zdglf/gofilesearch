package base_struct

import (
	"errors"
	"fmt"
	"time"
)

type JsonTime time.Time

func (that JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(that).Format("20060102"))
	return []byte(stamp), nil
}

func (that *JsonTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 2 {
		return errors.New("date must be string")
	}
	layout := "20060102"
	loc, _ := time.LoadLocation("Asia/Chongqing")
	var t time.Time
	t, err = time.ParseInLocation(layout, string(data[1:len(data)-1]), loc)
	*that = JsonTime(t)
	return
}
