package database

import (
	"errors"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

var (
	ErrDataNotExists = errors.New("data not exists")
)

func getEntities[T comparable](iter *spanner.RowIterator) ([]*T, error) {
	results := []*T{}
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var entity T
		if err := row.ToStruct(&entity); err != nil {
			return nil, err
		}
		results = append(results, &entity)
	}
	return results, nil
}
