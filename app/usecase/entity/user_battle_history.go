package entity

type UserBattleHistory struct {
	UserBattleHistoryID string
	UserID              string
	MyMonsterID         string
	OpponentMonsterID   string
	CreatedAt           int64
}
