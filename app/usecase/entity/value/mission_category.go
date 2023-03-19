package value

type MissionCategory string

var (
	LoginCategory  MissionCategory = "LOGIN"
	BattleCategory MissionCategory = "BATTLE"
	CoinCategory   MissionCategory = "COIN"
	LevelCategory  MissionCategory = "LEVEL"
	ItemCategory   MissionCategory = "ITEM"
)

func (m MissionCategory) IsLoginRelation() bool {
	return m == LoginCategory || m.isAllRelation()
}

func (m MissionCategory) IsBattleRelation() bool {
	return m == BattleCategory || m.isAllRelation()
}

func (m MissionCategory) IsLevelRelation() bool {
	return m == LevelCategory || m.isAllRelation()
}

func (m MissionCategory) isAllRelation() bool {
	return m == CoinCategory || m == ItemCategory
}
