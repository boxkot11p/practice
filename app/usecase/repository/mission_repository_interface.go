package repository

import (
	"context"

	"practice.dev/usecase/entity"
)

type MissionRepositoryInterface interface {
	FindAll(ctx context.Context) ([]*entity.Mission, error)
}
