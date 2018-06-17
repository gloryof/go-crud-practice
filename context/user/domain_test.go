package user

import "time"

func createDateTime(value string) time.Time {

	loc, _ := time.LoadLocation("Asia/Tokyo")
	b, _ := time.ParseInLocation("2006-01-02", "1986-12-16", loc)

	return b
}
