package event

import (
	"time"
)

type Login struct {
	Id      int64      `db:"id" json:"id"`
	UserId  int64      `db:"user_id" json:"userId"`
	EventAt *time.Time `db:"event_at" json:"eventAt"`
}
