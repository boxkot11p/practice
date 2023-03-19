package repository

import (
	"context"

	"practice.dev/usecase/entity"
)

type UserRepositoryInterface interface {
	FindByUserID(ctx context.Context, userID string) (*entity.User, error)
	FindMissionAchievementByUserID(ctx context.Context, userID string) ([]*entity.UserMissionAchievement, error)
	FindBattleHistoryByUserID(ctx context.Context, userID string) ([]*entity.UserBattleHistory, error)
	Update(ctx context.Context, u *entity.User, umas []*entity.UserMissionAchievement, ubh *entity.UserBattleHistory) error
}
