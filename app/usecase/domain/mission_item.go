package domain

import (
	"practice.dev/usecase/entity"
)

type missionItem struct {
	mission
}

func newMissionItem(m *entity.Mission, umas []*entity.UserMissionAchievement, ubhs []*entity.UserBattleHistory) MissionInterface {
	return &missionItem{
		mission: mission{
			Mission: m,
			umas:    umas,
			ubhs:    ubhs,
		},
	}
}

func (m missionItem) IsAchievement(u *entity.User, at int64) (bool, error) {
	umas, _, err := m.getData(at)
	if err != nil {
		return false, err
	}
	if m.isSkip(umas) {
		return false, nil
	}

	for _, v := range u.Items {
		if v.ItemID == m.ConditionItemID {
			return true, nil
		}
	}
	return false, nil
}
