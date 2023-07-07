package event

import "time"

type KillMonster struct {
	Id            int64      `db:"id" json:"Id"`
	UserId        int64      `db:"user_id" json:"userId"`
	KillMonsterId int64      `db:"kill_monster_id" json:"killMonsterId"`
	EventAt       *time.Time `db:"event_at" json:"eventAt"`
}
