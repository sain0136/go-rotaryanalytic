package utils

import (
	"fmt"
	"time"
)

const UtcDateFormat = "2006-01-02T15:04:05.000Z"

func ParseToTimeWithUtc(date string) (time.Time, error) {
	d, err := time.Parse(UtcDateFormat, date)
	if err != nil {
		fmt.Println(err)
		return time.UnixMilli(0), err
	}
	return d, nil
}
