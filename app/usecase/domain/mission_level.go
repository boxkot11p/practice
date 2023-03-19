package domain

import (
	"practice.dev/usecase/entity"
)

type missionLevel struct {
	mission
}

func newMissionLevel(m *entity.Mission, umas []*entity.UserMissionAchievement, ubhs []*entity.UserBattleHistory) MissionInterface {
	return &missionLevel{
		mission: mission{
			Mission: m,
			umas:    umas,
			ubhs:    ubhs,
		},
	}
}

func (m missionLevel) IsAchievement(u *entity.User, at int64) (bool, error) {
	umas, _, err := m.getData(at)
	if err != nil {
		return false, err
	}
	if m.isSkip(umas) {
		return false, nil
	}

	if m.ConditionLevelMonsterID != "" {
		return m.isAchievementSpecifiedMonsterLevel(u), nil
	}
	if m.ConditionLevelHaveMonsterNumber > 0 {
		return m.isAchievementHaveConditionLevelMonster(u), nil
	}
	return false, nil
}

func (m missionLevel) isAchievementSpecifiedMonsterLevel(u *entity.User) bool {
	for _, v := range u.Monsters {
		if v.MonsterID == m.ConditionLevelMonsterID {
			return v.Level >= m.ConditionLevel
		}
	}
	return false
}

func (m mission) isAchievementHaveConditionLevelMonster(u *entity.User) bool {
	cnt := 0
	for _, v := range u.Monsters {
		if v.Level >= m.ConditionLevel {
			cnt++
		}
	}
	return cnt >= int(m.ConditionLevelHaveMonsterNumber)
}
