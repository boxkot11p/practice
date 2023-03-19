package database

import (
	"context"

	"cloud.google.com/go/spanner"
)

type UserItem struct {
	UserID string `spanner:"UserID"`
	ItemID string `spanner:"ItemID"`
}

func replaceUserItemMutations(ctx context.Context, ums []*UserItem) ([]*spanner.Mutation, error) {
	ms := []*spanner.Mutation{}
	for _, v := range ums {
		s, err := spanner.InsertOrUpdateStruct("UserItem", v)
		if err != nil {
			return nil, err
		}

		ms = append(ms, s)
	}
	return ms, nil
}

func findUserItemByUserID(ctx context.Context, cli *spanner.Client, userID string) ([]*UserItem, error) {
	iter := cli.Single().Query(ctx, spanner.Statement{
		SQL: "SELECT * FROM UserItem WHERE UserID = @userID",
		Params: map[string]interface{}{
			"userID": userID,
		},
	})

	results, err := getEntities[UserItem](iter)
	if err != nil {
		return nil, err
	}
	return results, nil
}
