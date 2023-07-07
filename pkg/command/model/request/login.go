package request

import (
	"time"
)

type LoginReq struct {
	UserId  int64      `db:"user_id" json:"userId"`
	EventAt *time.Time `db:"event_at" json:"eventAt"`
}
