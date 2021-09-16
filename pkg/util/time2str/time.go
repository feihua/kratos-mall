package time2str

import "time"

func String(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
