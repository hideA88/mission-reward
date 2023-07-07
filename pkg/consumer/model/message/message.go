package message

import "time"

type Login struct {
	UserId  int64      `db:"user_id" json:"userId"`
	EventAt *time.Time `db:"event_at" json:"eventAt"`
}

type GetCoin struct {
	UserId  int64      `db:"user_id" json:"userId"`
	EventAt *time.Time `db:"event_at" json:"eventAt"`
}

type GetItem struct {
	UserId  int64      `db:"user_id" json:"userId"`
	EventAt *time.Time `db:"event_at" json:"eventAt"`
}

type KillMonster struct {
	UserId  int64      `db:"user_id" json:"userId"`
	EventAt *time.Time `db:"event_at" json:"eventAt"`
}

type LevelUp struct {
	UserId  int64      `db:"user_id" json:"userId"`
	EventAt *time.Time `db:"event_at" json:"eventAt"`
}

type OpenMission struct {
	UserId    int64      `db:"user_id" json:"userId"`
	MissionId int64      `db:"mission_id" json:"missionId"`
	EventAt   *time.Time `db:"event_at" json:"eventAt"`
}
