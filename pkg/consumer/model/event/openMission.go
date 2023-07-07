package event

import "time"

type OpenMission struct {
	Id        int64      `db:"id" json:"id"`
	UserId    int64      `db:"user_id" json:"userId"`
	MissionId int64      `db:"mission_id" json:"missionId"`
	EventAt   *time.Time `db:"event_at" json:"eventAt"`
}
