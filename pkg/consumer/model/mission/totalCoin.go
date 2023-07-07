package mission

type TotalCoinMission struct {
	Id        int64 `db:"id" json:"id"`
	MissionId int64 `db:"mission_id" json:"missionId"`
	Size      int   `db:"size" json:"size"`
}
