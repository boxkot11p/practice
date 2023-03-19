package database

import (
	"context"

	"cloud.google.com/go/spanner"
)

type User struct {
	UserID string `spanner:"UserID"`
	Name   string `spanner:"Name"`
	Coin   int64  `spanner:"Coin"`
}

func updateUser(ctx context.Context, u *User) (*spanner.Mutation, error) {
	return spanner.InsertOrUpdateStruct("User", u)
}

func findOneUserByUserID(ctx context.Context, cli *spanner.Client, userID string) (*User, error) {
	iter := cli.Single().Query(ctx, spanner.Statement{
		SQL: "SELECT * FROM User Where UserID = @userID",
		Params: map[string]interface{}{
			"userID": userID,
		},
	})

	results, err := getEntities[User](iter)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, ErrDataNotExists
	}
	return results[0], nil
}
