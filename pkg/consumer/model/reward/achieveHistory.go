package reward

import "time"

type MissionWithAchieveHistory struct {
	MissionId      int64      `db:"mission_id" json:"missionId"`
	Name           string     `db:"mission_name" json:"name"`
	MissionType    string     `db:"mission_type" json:"missionType"`
	RewardType     string     `db:"reward_type" json:"rewardType"`
	RewardCoinSize int        `db:"reward_coin_size" json:"rewardCoinSize"`
	RewardItemId   int64      `db:"reward_item_id" json:"rewardItemId"`
	ResetTime      string     `db:"reset_time" json:"resetTime"`
	OpenMissionId  int64      `db:"open_mission_id" json:"openMissionId"`
	MissionStatus  string     `db:"mission_status" json:"missionStatus"`
	UserId         int64      `db:"user_id" json:"userId"`
	AchievedAt     *time.Time `db:"achieved_at" json:"achievedAt"`
}
