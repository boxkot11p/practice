package database

import (
	"context"

	"cloud.google.com/go/spanner"
)

type Mission struct {
	MissionID                         string `spanner:"MissionID"`
	OrderConditionMissionID           string `spanner:"OrderConditionMissionID"`
	GiftItemID                        string `spanner:"GiftItemID"`
	Name                              string `spanner:"Name"`
	Category                          string `spanner:"Category"`
	ResetCycle                        string `spanner:"ResetCycle"`
	ResetWeek                         string `spanner:"ResetWeek"`
	ResetHour                         int64  `spanner:"ResetHour"`
	ResetTime                         int64  `spanner:"ResetTime"`
	ConditionLevelMonsterID           string `spanner:"ConditionLevelMonsterID"`
	ConditionLevel                    int64  `spanner:"ConditionLevel"`
	ConditionLevelHaveMonsterNumber   int64  `spanner:"ConditionLevelHaveMonsterNumber"`
	ConditionItemID                   string `spanner:"ConditionItemID"`
	ConditionTargetMonsterID          string `spanner:"ConditionTargetMonsterID"`
	ConditionSubjugationMonsterNumber int64  `spanner:"ConditionSubjugationMonsterNumber"`
	ConditionHaveCoin                 int64  `spanner:"ConditionHaveCoin"`
	GiftCoin                          int64  `spanner:"GiftCoin"`
}

func getMissions(ctx context.Context, cli *spanner.Client) ([]*Mission, error) {
	iter := cli.Single().Query(ctx, spanner.Statement{
		SQL: "SELECT * FROM Mission",
	})

	results, err := getEntities[Mission](iter)
	if err != nil {
		return nil, err
	}
	return results, nil
}
