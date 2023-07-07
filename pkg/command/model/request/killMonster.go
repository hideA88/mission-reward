package request

import "time"

type KillMonsterReq struct {
	UserId        int64      `db:"user_id" json:"userId"`
	KillMonsterId int64      `db:"kill_monster_id" json:"killMonsterId"`
	EventAt       *time.Time `db:"event_at" json:"eventAt"`
}
