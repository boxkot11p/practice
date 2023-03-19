package domain

import (
	"practice.dev/usecase/entity"
)

type missionLogin struct {
	mission
}

func newMissionLogin(m *entity.Mission, umas []*entity.UserMissionAchievement, ubhs []*entity.UserBattleHistory) MissionInterface {
	return &missionLogin{
		mission: mission{
			Mission: m,
			umas:    umas,
			ubhs:    ubhs,
		},
	}
}

func (m missionLogin) IsAchievement(_ *entity.User, at int64) (bool, error) {
	umas, _, err := m.getData(at)
	if err != nil {
		return false, err
	}
	if m.isSkip(umas) {
		return false, nil
	}

	return len(umas) == 0, nil
}
