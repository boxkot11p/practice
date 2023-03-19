package database

import (
	"context"
	"time"

	"cloud.google.com/go/spanner"
	"github.com/google/uuid"
)

type UserMissionAchievement struct {
	UserMissionAchievementID string    `spanner:"UserMissionAchievementID"`
	UserID                   string    `spanner:"UserID"`
	MissionID                string    `spanner:"MissionID"`
	CreatedAt                time.Time `spanner:"CreatedAt"`
}

func createUserMissionAchievementInsertMutations(ctx context.Context, umas []*UserMissionAchievement) ([]*spanner.Mutation, error) {
	ms := []*spanner.Mutation{}
	for _, v := range umas {
		v.UserMissionAchievementID = uuid.New().String()
		s, err := spanner.InsertStruct("UserMissionAchievement", v)
		if err != nil {
			return nil, err
		}

		ms = append(ms, s)
	}
	return ms, nil
}

func findUserMissionAchievementByUserID(ctx context.Context, cli *spanner.Client, userID string) ([]*UserMissionAchievement, error) {
	iter := cli.Single().Query(ctx, spanner.Statement{
		SQL: "SELECT * FROM UserMissionAchievement WHERE UserID = @userID",
		Params: map[string]interface{}{
			"userID": userID,
		},
	})

	results, err := getEntities[UserMissionAchievement](iter)
	if err != nil {
		return nil, err
	}
	return results, nil
}
