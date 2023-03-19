package domain

import (
	"practice.dev/usecase/entity"
)

type missionCoin struct {
	mission
}

func newMissionCoin(m *entity.Mission, umas []*entity.UserMissionAchievement, ubhs []*entity.UserBattleHistory) MissionInterface {
	return &missionCoin{
		mission: mission{
			Mission: m,
			umas:    umas,
			ubhs:    ubhs,
		},
	}
}

func (m missionCoin) IsAchievement(u *entity.User, at int64) (bool, error) {
	umas, _, err := m.getData(at)
	if err != nil {
		return false, err
	}
	if m.isSkip(umas) {
		return false, nil
	}

	if u.Coin >= m.ConditionHaveCoin {
		return true, nil
	}
	return false, nil
}
