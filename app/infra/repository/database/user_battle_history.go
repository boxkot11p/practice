package database

import (
	"context"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
)

type UserBattleHistory struct {
	UserBattleHistoryID string    `spanner:"UserBattleHistoryID"`
	UserID              string    `spanner:"UserID"`
	MyMonsterID         string    `spanner:"MyMonsterID"`
	OpponentMonsterID   string    `spanner:"OpponentMonsterID"`
	CreatedAt           time.Time `spanner:"CreatedAt"`
}

func createUserBattleHistoryMutation(ctx context.Context, ubh *UserBattleHistory) (*spanner.Mutation, error) {
	ubh.UserBattleHistoryID = uuid.New().String()
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
