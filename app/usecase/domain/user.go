package domain

import (
	"practice.dev/usecase/entity"
)

type UserInterface interface {
	GetUser() *entity.User
	AchievementLogin(
		ms []*entity.Mission,
		umas []*entity.UserMissionAchievement,
		at int64,
	) ([]*entity.UserMissionAchievement, error)
	AchievementBattle(
		ms []*entity.Mission,
		umas []*entity.UserMissionAchievement,
		ubhs []*entity.UserBattleHistory,
		at int64,
	) ([]*entity.UserMissionAchievement, error)
	AchievementLevel(
		ms []*entity.Mission,
		umas []*entity.UserMissionAchievement,
		at int64,
	) ([]*entity.UserMissionAchievement, error)
}

type user struct {
	u *entity.User
}

func NewUser(u *entity.User) UserInterface {
	return &user{
		u: u,
	}
}

func (u user) GetUser() *entity.User {
	return u.u
}

func (u *user) AchievementLogin(
	ms []*entity.Mission,
	umas []*entity.UserMissionAchievement,
	at int64,
) ([]*entity.UserMissionAchievement, error) {
	fms := []*entity.Mission{}
	for _, v := range ms {
		if v.Category.IsLoginRelation() {
			fms = append(fms, v)
		}
	}
	return u.achievement(fms, umas, nil, at)
}

func (u user) AchievementBattle(
	ms []*entity.Mission,
	umas []*entity.UserMissionAchievement,
	ubhs []*entity.UserBattleHistory,
	at int64,
) ([]*entity.UserMissionAchievement, error) {
	fms := []*entity.Mission{}
	for _, v := range ms {
		if v.Category.IsBattleRelation() {
			fms = append(fms, v)
		}
	}
	return u.achievement(fms, umas, ubhs, at)
}

func (u user) AchievementLevel(
	ms []*entity.Mission,
	umas []*entity.UserMissionAchievement,
	at int64,
) ([]*entity.UserMissionAchievement, error) {
	fms := []*entity.Mission{}
	for _, v := range ms {
		if v.Category.IsLevelRelation() {
			fms = append(fms, v)
		}
	}
	return u.achievement(fms, umas, nil, at)
}

func (u *user) achievement(
	ms []*entity.Mission,
	umas []*entity.UserMissionAchievement,
	ubhs []*entity.UserBattleHistory,
	at int64,
) ([]*entity.UserMissionAchievement, error) {
	results := []*entity.UserMissionAchievement{}
	for {
		count := len(results)
		for _, v := range ms {
			d, err := NewMission(v, umas, ubhs)
			if err != nil {
				return nil, err
			}
			ok, err := d.IsAchievement(u.u, at)
			if err != nil {
				return nil, err
			}
			if ok {
				uma := &entity.UserMissionAchievement{
					UserID:    u.u.UserID,
					MissionID: v.MissionID,
					CreatedAt: at,
				}
				results = append(results, uma)
				umas = append(umas, uma)
				u.u.AddCoin(v.GiftCoin)
			}
		}
		// NOTE: ミッション達成で他ミッションに影響を与える可能性があるため達成時は再度ミッションをチェック
		if len(results) == count {
			break
		}
	}
	return results, nil
}
