package util

import "time"

const layout = "2006-01-02"

func Today() string {
	return time.Now().Format(layout)
}

func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse(layout, dateStr)
}

func DaysBetween(a, b time.Time) int {
	diff := b.Sub(a)
	return int(diff.Hours() / 24)
}
