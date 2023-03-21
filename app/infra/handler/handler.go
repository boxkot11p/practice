package handler

import (
	"context"

	"practice.dev/api"
	"practice.dev/usecase/entity"
)

func (s Server) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	if err := s.actionService.Login(ctx, req.UserID, req.CreatedAt); err != nil {
		return nil, err
	}
	return &api.LoginResponse{}, nil
}

func (s Server) Battle(ctx context.Context, req *api.BattleRequest) (*api.BattleResponse, error) {
	ubh := &entity.UserBattleHistory{
		UserID:            req.UserID,
		MyMonsterID:       req.MyMonsterID,
		OpponentMonsterID: req.OpponentMonsterID,
		CreatedAt:         req.CreatedAt,
	}
	if err := s.actionService.Battle(ctx, ubh); err != nil {
		return nil, err
	}
	return &api.BattleResponse{}, nil
}

func (s Server) LevelUp(ctx context.Context, req *api.LevelUpRequest) (*api.LevelUpResponse, error) {
	if err := s.actionService.LevelUp(ctx, req.UserID, req.MonsterID, req.Level, req.CreatedAt); err != nil {
		return nil, err
	}
	return &api.LevelUpResponse{}, nil
}
