package domain

import (
	"fmt"
	"time"

	"practice.dev/usecase/entity"
	"practice.dev/usecase/entity/value"
)

type MissionInterface interface {
	IsAchievement(u *entity.User, at int64) (bool, error)
}

type mission struct {
	*entity.Mission
	umas []*entity.UserMissionAchievement
	ubhs []*entity.UserBattleHistory
}

func NewMission(m *entity.Mission, umas []*entity.UserMissionAchievement, ubhs []*entity.UserBattleHistory) (MissionInterface, error) {
	switch m.Category {
	case value.LoginCategory:
		return newMissionLogin(m, umas, ubhs), nil
	case value.BattleCategory:
		return newMissionBattle(m, umas, ubhs), nil
	case value.LevelCategory:
		return newMissionLevel(m, umas, ubhs), nil
	case value.CoinCategory:
		return newMissionCoin(m, umas, ubhs), nil
	case value.ItemCategory:
		return newMissionItem(m, umas, ubhs), nil
	default:
		return nil, fmt.Errorf("not match category: %s", m.Category)
	}
}

func (m mission) getData(at int64) ([]*entity.UserMissionAchievement, []*entity.UserBattleHistory, error) {
	umas := m.umas
	ubhs := m.ubhs
	if m.Mission.IsNeedReset() {
		var err error
		umas, ubhs, err = m.filterResetData(at)
		if err != nil {
			return nil, nil, err
		}
	}
	return umas, ubhs, nil
}

func (m mission) isSkip(umas []*entity.UserMissionAchievement) bool {
	// NOTE: 前提条件の確認
	if m.OrderConditionMissionID != "" {
		ok := false
		for _, v := range umas {
			if v.MissionID == m.OrderConditionMissionID {
				ok = true
				break
			}
		}
		if !ok {
			return true
		}
	}

	// NOTE: 達成済みの場合はSKIP
	for _, v := range umas {
		if v.MissionID == m.MissionID {
			return true
		}
	}
	return false
}

func (m mission) filterResetData(at int64) ([]*entity.UserMissionAchievement, []*entity.UserBattleHistory, error) {
	umas := []*entity.UserMissionAchievement{}
	ubhs := []*entity.UserBattleHistory{}
	now := time.Unix(at, 0).UTC()
	resetTime, err := m.ResetStartTime(now)
	if err != nil {
		return nil, nil, err
	}
	for _, v := range m.umas {
		if v.CreatedAt > resetTime.Unix() {
			umas = append(umas, v)
		}
	}
	for _, v := range m.ubhs {
		if v.CreatedAt > resetTime.Unix() {
			ubhs = append(ubhs, v)
		}
	}
	return umas, ubhs, nil
}
