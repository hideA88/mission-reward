package event

import "time"

type LevelUp struct {
	Id            int64      `db:"id" json:"Id"`
	UserId        int64      `db:"user_id" json:"userId"`
	UserMonsterId int64      `db:"user_monster_id" json:"userMonsterId"`
	LevelUpSize   int        `db:"level_up_size" json:"levelUpSize"`
	EventAt       *time.Time `db:"event_at" json:"eventAt"`
}
