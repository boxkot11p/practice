package database

import (
	"context"
	"time"

	"cloud.google.com/go/spanner"
)

type UserBattleHistory struct {
	UserActionHistoryID string    `spanner:"UserActionHistoryID"`
	UserID              string    `spanner:"UserID"`
	MonsterID           string    `spanner:"MonsterID"`
	CreatedAt           time.Time `spanner:"CreatedAt"`
}

func createUserBattleHistoryMutation(ctx context.Context, ubh *UserBattleHistory) (*spanner.Mutation, error) {
	return spanner.InsertStruct("UserBattleHistory", ubh)
}

func findUserBattleHistoryByUserID(ctx context.Context, cli *spanner.Client, userID string) ([]*UserBattleHistory, error) {
	iter := cli.Single().Query(ctx, spanner.Statement{
		SQL: "SELECT * FROM UserBattleHistory WHERE UserID = @userID",
		Params: map[string]interface{}{
			"userID": userID,
		},
	})

	results, err := getEntities[UserBattleHistory](iter)
	if err != nil {
		return nil, err
	}
	return results, nil
}
