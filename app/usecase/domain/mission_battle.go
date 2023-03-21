package domain

import (
	"practice.dev/usecase/entity"
)

type missionBattle struct {
	mission
}

func newMissionBattle(m *entity.Mission, umas []*entity.UserMissionAchievement, ubhs []*entity.UserBattleHistory) MissionInterface {
	return &missionBattle{
		mission: mission{
			Mission: m,
			umas:    umas,
			ubhs:    ubhs,
		},
	}
}

func (m missionBattle) IsAchievement(_ *entity.User, at int64) (bool, error) {
	umas, ubhs, err := m.getData(at)
	if err != nil {
		return false, err
	}
	if m.isSkip(umas) {
		return false, nil
	}

	if m.Mission.ConditionTargetMonsterID != "" {
		return m.isAchievementSpecifiedBattle(ubhs), nil
	}
	if m.Mission.ConditionSubjugationMonsterNumber > 0 {
		return m.isAchievementRandomBattle(ubhs), nil
	}
	return false, nil
}

func (m missionBattle) isAchievementSpecifiedBattle(ubhs []*entity.UserBattleHistory) bool {
	monsterID := m.Mission.ConditionTargetMonsterID
	for _, v := range ubhs {
		if monsterID == v.OpponentMonsterID {
			return true
		}
	}
	return false
}

func (m missionBattle) isAchievementRandomBattle(ubhs []*entity.UserBattleHistory) bool {
	return int64(len(ubhs)) >= m.Mission.ConditionSubjugationMonsterNumber
}
