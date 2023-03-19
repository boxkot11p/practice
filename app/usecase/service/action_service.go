package service

import (
	"context"

	"practice.dev/usecase/domain"
	"practice.dev/usecase/entity"
	"practice.dev/usecase/repository"
)

type ActionService struct {
	msRepo repository.MissionRepositoryInterface
	uRepo  repository.UserRepositoryInterface
}

func NewActionService(msRepo repository.MissionRepositoryInterface, uRepo repository.UserRepositoryInterface) *ActionService {
	return &ActionService{
		msRepo: msRepo,
		uRepo:  uRepo,
	}
}

func (s ActionService) Login(ctx context.Context, userID string, at int64) error {
	u, err := s.uRepo.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}
	ms, err := s.msRepo.FindAll(ctx)
	if err != nil {
		return err
	}
	umas, err := s.uRepo.FindMissionAchievementByUserID(ctx, userID)
	if err != nil {
		return err
	}
	d := domain.NewUser(u)
	results, err := d.AchievementLogin(ms, umas, at)
	if err != nil {
		return err
	}
	return s.uRepo.Update(ctx, d.GetUser(), results, nil)
}

func (s ActionService) Battle(ctx context.Context, ubh *entity.UserBattleHistory) error {
	userID := ubh.UserID
	at := ubh.CreatedAt
	u, err := s.uRepo.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}
	ms, err := s.msRepo.FindAll(ctx)
	if err != nil {
		return err
	}
	umas, err := s.uRepo.FindMissionAchievementByUserID(ctx, userID)
	if err != nil {
		return err
	}
	ubhs, err := s.uRepo.FindBattleHistoryByUserID(ctx, userID)
	if err != nil {
		return err
	}
	ubhs = append(ubhs, ubh)
	d := domain.NewUser(u)
	results, err := d.AchievementBattle(ms, umas, ubhs, at)
	if err != nil {
		return err
	}
	return s.uRepo.Update(ctx, d.GetUser(), results, ubh)
}

func (s ActionService) LevelUp(ctx context.Context, userID, monsterID string, level, at int64) error {
	u, err := s.uRepo.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}
	ms, err := s.msRepo.FindAll(ctx)
	if err != nil {
		return err
	}
	umas, err := s.uRepo.FindMissionAchievementByUserID(ctx, userID)
	if err != nil {
		return err
	}
	u.UpdateMonsterLevel(monsterID, level)
	d := domain.NewUser(u)
	results, err := d.AchievementLevel(ms, umas, at)
	if err != nil {
		return err
	}
	return s.uRepo.Update(ctx, d.GetUser(), results, nil)
}
