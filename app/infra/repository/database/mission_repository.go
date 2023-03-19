package database

import (
	"context"

	"cloud.google.com/go/spanner"
	"practice.dev/usecase/entity"
	"practice.dev/usecase/repository"
)

type MissionRepository struct {
	cli *spanner.Client
}

func NewMissionRepository(cli *spanner.Client) repository.MissionRepositoryInterface {
	return &MissionRepository{
		cli: cli,
	}
}

func (r MissionRepository) FindAll(ctx context.Context) ([]*entity.Mission, error) {
	ms, err := getMissions(ctx, r.cli)
	if err != nil {
		return nil, err
	}
	results := []*entity.Mission{}
	for _, v := range ms {
		results = append(results, &entity.Mission{
			MissionID:                         v.MissionID,
			OrderConditionMissionID:           v.OrderConditionMissionID,
			GiftItemID:                        v.GiftItemID,
			Name:                              v.Name,
			ResetCycle:                        v.ResetCycle,
			ResetWeek:                         v.ResetWeek,
			ResetHour:                         v.ResetHour,
			ResetTime:                         v.ResetTime,
			ConditionLevelMonsterID:           v.ConditionLevelMonsterID,
			ConditionLevel:                    v.ConditionLevel,
			ConditionLevelHaveMonsterNumber:   v.ConditionLevelHaveMonsterNumber,
			ConditionItemID:                   v.ConditionItemID,
			ConditionTargetMonsterID:          v.ConditionTargetMonsterID,
			ConditionSubjugationMonsterNumber: v.ConditionSubjugationMonsterNumber,
			ConditionHaveCoin:                 v.ConditionHaveCoin,
			GiftCoin:                          v.GiftCoin,
		})
	}
	return results, nil
}
