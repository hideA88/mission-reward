package reward

import (
	"github.com/hideA88/mission-reward/pkg/consumer/model/event"
	"time"
)

type UserAchieveReq struct {
	UserId         int64      `db:"user_id" json:"userId"`
	MissionId      int64      `db:"mission_id" json:"missionId"`
	AchievedAt     *time.Time `db:"achieved_at" json:"achievedAt"`
	RewardType     string     `db:"reward_type" json:"rewardType"`
	RewardCoinSize int        `db:"reward_coin_size" json:"rewardCoinSize"`
	RewardItemId   int64      `db:"reward_item_id" json:"rewardItemId"`
	OpenMissionId  int64      `db:"open_mission_id" json:"openMissionId"`
}

type UserAchieve struct {
	Id             int64              `db:"id" json:"id"`
	UserId         int64              `db:"user_id" json:"userId"`
	MissionId      int64              `db:"mission_id" json:"missionId"`
	AchievedAt     *time.Time         `db:"achieved_at" json:"achievedAt"`
	RewardType     string             `db:"reward_type" json:"rewardType"`
	RewardId       int64              `db:"reward_id" json:"rewardId"`
	RewardCoinSize int                `db:"reward_coin_size" json:"rewardCoinSize"`
	RewardItemId   int64              `db:"reward_item_id" json:"rewardItemId"`
	OpenMission    *event.OpenMission `db:"open_mission_id" json:"openMissionId"`
}
