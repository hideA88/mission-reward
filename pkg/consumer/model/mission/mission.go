package mission

type Type string

const (
	LOGIN        Type = "login"
	GET_ITEM     Type = "getItem"
	TOTAL_COIN   Type = "totalCoin"
	LEVEL_UP     Type = "levelUp"
	KILL_MONSTER Type = "killMonster"
)

type Mission struct {
	Id            int64  `db:"id" json:"id"`
	Name          string `db:"name" json:"name"`
	MissionType   string `db:"mission_type" json:"missionType"`
	RewardType    string `db:"reward_type" json:"rewardType"`
	ResetTime     string `db:"reset_time" json:"resetTime"`
	OpenMissionId int64  `db:"open_mission_id" json:"openMissionId"`
	MissionStatus string `db:"mission_status" json:"missionStatus"`
}
