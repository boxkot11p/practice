package database

import (
	"context"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
)

type UserMonster struct {
	UserMonsterID string `spanner:"UserMonsterID"`
	UserID        string `spanner:"UserID"`
	MonsterID     string `spanner:"MonsterID"`
	Level         int64  `spanner:"Level"`
}

func replaceUserMonsterMutations(ctx context.Context, ums []*UserMonster) ([]*spanner.Mutation, error) {
	ms := []*spanner.Mutation{}
	for _, v := range ums {
		if v.UserMonsterID == "" {
			v.UserMonsterID = uuid.New().String()
		}
		s, err := spanner.InsertOrUpdateStruct("UserMissionAchievement", v)
		if err != nil {
			return nil, err
		}

		ms = append(ms, s)
	}
	return ms, nil
}

func findUserMonsterByUserID(ctx context.Context, cli *spanner.Client, userID string) ([]*UserMonster, error) {
	iter := cli.Single().Query(ctx, spanner.Statement{
		SQL: "SELECT * FROM UserMonster WHERE UserID = @userID",
		Params: map[string]interface{}{
			"userID": userID,
		},
	})

	results, err := getEntities[UserMonster](iter)
	if err != nil {
		return nil, err
	}
	return results, nil
}
