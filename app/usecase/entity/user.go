package entity

type User struct {
	UserID   string
	Name     string
	Coin     int64
	Monsters []*Monster
	Items    []*Item
}

func (u *User) AddCoin(c int64) {
	u.Coin += c
}

func (u *User) AddItem(itemID string) {
	if itemID == "" {
		return
	}
	u.Items = append(u.Items, &Item{
		ItemID: itemID,
	})
}

func (u *User) UpdateMonsterLevel(monsterID string, level int64) {
	for k, v := range u.Monsters {
		if v.MonsterID == monsterID {
			u.Monsters[k].Level += level
			break
		}
	}
}
