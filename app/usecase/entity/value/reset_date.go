package value

import (
	"fmt"
	"math"
	"time"
)

type ResetDate struct {
	resetCycle string
	resetWeek  string
	resetHour  int
	resetTime  int
}

func NewResetData(resetCycle, resetWeek string, resetHour, resetTime int) *ResetDate {
	return &ResetDate{
		resetCycle: resetCycle,
		resetWeek:  resetWeek,
		resetHour:  resetHour,
		resetTime:  resetTime,
	}
}

func (r ResetDate) GetResetStartTime(at time.Time) (time.Time, error) {
	switch r.resetCycle {
	case "NONE":
		return time.Time{}, nil
	case "DAILY":
		return r.getResetStartTimeDaily(at.UTC()), nil
	case "WEEKLY":
		return r.getResetStartTimeWeekly(at.UTC())
	default:
		return time.Time{}, fmt.Errorf("not match reset cycle: %s", r.resetCycle)
	}
}

func (r ResetDate) getResetStartTimeDaily(at time.Time) time.Time {
	day := at.Day()
	if r.resetHour >= at.Hour() && r.resetTime >= at.Minute() {
		day = at.Day() - 1
	}
	return time.Date(at.Year(), at.Month(), day, r.resetHour, r.resetTime, 0, 0, time.UTC)
}

func (r ResetDate) getResetStartTimeWeekly(at time.Time) (time.Time, error) {
	wdays := map[string]int{
		"SUNDAY":    int(time.Sunday),
		"MONDAY":    int(time.Monday),
		"TUESDAY":   int(time.Tuesday),
		"WEDNESDAY": int(time.Wednesday),
		"THURSDAY":  int(time.Thursday),
		"FRIDAY":    int(time.Friday),
		"SATURDAY":  int(time.Saturday),
	}
	w, ok := wdays[r.resetWeek]
	if !ok {
		return time.Time{}, fmt.Errorf("not match reset week: %s", r.resetWeek)
	}
	// NOTE: 前のリセット日付までの日数を計算
	day := w - int(at.Weekday())
	if day > 0 {
		day = 7 - day
	} else if day < 0 {
		day = int(math.Abs(float64(day)))
	} else if r.resetHour >= at.Hour() && r.resetTime >= at.Minute() {
		day = 7
	}
	return time.Date(at.Year(), at.Month(), at.Day()-day, r.resetHour, r.resetTime, 0, 0, time.UTC), nil
}
