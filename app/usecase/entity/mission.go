package entity

import (
	"fmt"
	"math"
	"time"

	"practice.dev/usecase/entity/value"
)

type Mission struct {
	MissionID                         string
	OrderConditionMissionID           string
	GiftItemID                        int64
	Category                          value.MissionCategory
	Name                              int64
	ResetCycle                        string
	ResetWeek                         string
	ResetHour                         int
	ResetTime                         int
	ConditionLevelMonsterID           string
	ConditionLevel                    int64
	ConditionLevelHaveMonsterNumber   int64
	ConditionItemID                   string
	ConditionTargetMonsterID          string
	ConditionSubjugationMonsterNumber int64
	ConditionHaveCoin                 int64
	GiftCoin                          int64
}

func (m Mission) IsNeedReset() bool {
	return m.ResetCycle != "NONE"
}

func (m Mission) ResetStartTime(at time.Time) (time.Time, error) {
	switch m.ResetCycle {
	case "NONE":
		return time.Time{}, nil
	case "DAILY":
		return time.Date(at.Year(), at.Month(), at.Day(), m.ResetHour, m.ResetTime, 0, 0, at.Location()), nil
	case "WEEKLY":
		wdays := map[string]int{
			"SUNDAY":    int(time.Sunday),
			"MONDAY":    int(time.Monday),
			"TUESDAY":   int(time.Tuesday),
			"WEDNESDAY": int(time.Wednesday),
			"THURSDAY":  int(time.Thursday),
			"FRIDAY":    int(time.Friday),
			"SATURDAY":  int(time.Saturday),
		}
		w, ok := wdays[m.ResetWeek]
		if !ok {
			return time.Time{}, fmt.Errorf("not match reset week: %s", m.ResetWeek)
		}
		d := int(math.Abs(float64(w - int(at.Weekday()))))
		return time.Date(at.Year(), at.Month(), at.Day()-d, m.ResetHour, m.ResetTime, 0, 0, at.Location()), nil
	default:
		return time.Time{}, fmt.Errorf("not match reset cycle: %s", m.ResetCycle)
	}
}
