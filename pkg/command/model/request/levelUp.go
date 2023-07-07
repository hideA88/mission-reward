package request

import "time"

type LevelUpReq struct {
	UserId        int64      `db:"user_id" json:"userId"`
	UserMonsterId int64      `db:"user_monster_id" json:"userMonsterId"`
	LevelUpSize   int        `db:"level_up_size" json:"levelUpSize"`
	EventAt       *time.Time `db:"event_at" json:"eventAt"`
}
