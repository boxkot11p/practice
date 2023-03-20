package database

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/spanner"
	"practice.dev/usecase/entity"
	"practice.dev/usecase/repository"
)

type UserRepository struct {
	cli *spanner.Client
}

func NewUserRepository(cli *spanner.Client) repository.UserRepositoryInterface {
	return &UserRepository{
		cli: cli,
	}
}

func (r UserRepository) FindByUserID(ctx context.Context, userID string) (*entity.User, error) {
	us, err := findOneUserByUserID(ctx, r.cli, userID)
	if err != nil {
		return nil, err
	}
	ums, err := findUserMonsterByUserID(ctx, r.cli, userID)
	if err != nil {
		return nil, err
	}
	uis, err := findUserItemByUserID(ctx, r.cli, userID)
	if err != nil {
		return nil, err
	}
	result := &entity.User{
		UserID: us.UserID,
		Name:   us.Name,
		Coin:   us.Coin,
	}
	for _, v := range ums {
		result.Monsters = append(result.Monsters, &entity.Monster{
			UserMonsterID: v.UserMonsterID,
			MonsterID:     v.MonsterID,
			Level:         v.Level,
		})
	}
	for _, v := range uis {
		result.Items = append(result.Items, &entity.Item{
			ItemID: v.ItemID,
		})
	}
	return result, nil
}

func (r UserRepository) FindMissionAchievementByUserID(ctx context.Context, userID string) ([]*entity.UserMissionAchievement, error) {
	umas, err := findUserMissionAchievementByUserID(ctx, r.cli, userID)
	if err != nil {
		return nil, err
	}
	results := []*entity.UserMissionAchievement{}
	for _, v := range umas {
		results = append(results, &entity.UserMissionAchievement{
			UserID:    v.UserID,
			MissionID: v.MissionID,
			CreatedAt: v.CreatedAt.Unix(),
		})
	}
	return results, nil
}

func (r UserRepository) FindBattleHistoryByUserID(ctx context.Context, userID string) ([]*entity.UserBattleHistory, error) {
	ubhs, err := findUserBattleHistoryByUserID(ctx, r.cli, userID)
	if err != nil {
		return nil, err
	}
	results := []*entity.UserBattleHistory{}
	for _, v := range ubhs {
		results = append(results, &entity.UserBattleHistory{
			UserBattleHistoryID: v.UserBattleHistoryID,
			UserID:              v.UserID,
			MonsterID:           v.MonsterID,
			CreatedAt:           v.CreatedAt.Unix(),
		})
	}
	return results, nil
}

func (r UserRepository) Update(
	ctx context.Context,
	u *entity.User,
	umas []*entity.UserMissionAchievement,
	ubhs *entity.UserBattleHistory) error {

	user := &User{
		UserID: u.UserID,
		Name:   u.Name,
		Coin:   u.Coin,
	}
	userMs := []*UserMonster{}
	for _, v := range u.Monsters {
		userMs = append(userMs, &UserMonster{
			UserMonsterID: v.UserMonsterID,
			UserID:        u.UserID,
			MonsterID:     v.MonsterID,
			Level:         v.Level,
		})
	}
	userIm := []*UserItem{}
	for _, v := range u.Items {
		userIm = append(userIm, &UserItem{
			UserID: u.UserID,
			ItemID: v.ItemID,
		})
	}

	umads := []*UserMissionAchievement{}
	for _, v := range umas {
		umads = append(umads, &UserMissionAchievement{
			UserID:    v.UserID,
			MissionID: v.MissionID,
			CreatedAt: time.Unix(v.CreatedAt, 0),
		})
	}

	var ubhds *UserBattleHistory
	if ubhs != nil {
		ubhds = &UserBattleHistory{
			UserID:    ubhs.UserID,
			MonsterID: ubhs.MonsterID,
			CreatedAt: time.Unix(ubhs.CreatedAt, 0),
		}
	}

	cr, err := r.cli.ReadWriteTransactionWithOptions(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		ms := []*spanner.Mutation{}
		uu, err := updateUser(ctx, user)
		if err != nil {
			return err
		}
		ms = append(ms, uu)
		ums, err := replaceUserMonsterMutations(ctx, userMs)
		if err != nil {
			return err
		}
		ms = append(ms, ums...)
		uis, err := replaceUserItemMutations(ctx, userIm)
		if err != nil {
			return err
		}
		ms = append(ms, uis...)
		cumaim, err := createUserMissionAchievementInsertMutations(ctx, umads)
		if err != nil {
			return err
		}
		ms = append(ms, cumaim...)
		if ubhds != nil {
			cubhm, err := createUserBattleHistoryMutation(ctx, ubhds)
			if err != nil {
				return err
			}
			ms = append(ms, cubhm)
		}
		return txn.BufferWrite(ms)
	}, spanner.TransactionOptions{CommitOptions: spanner.CommitOptions{ReturnCommitStats: true}})
	log.Default().Printf("%v", cr)
	return err
}
