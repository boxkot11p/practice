package entity

import (
	"time"

	"practice.dev/usecase/entity/value"
)

type Mission struct {
	MissionID                         string
	OrderConditionMissionID           string
	GiftItemID                        string
	Category                          value.MissionCategory
	Name                              string
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
	rd := value.NewResetData(m.ResetCycle, m.ResetWeek, m.ResetHour, m.ResetTime)
	return rd.GetResetStartTime(at)
}
