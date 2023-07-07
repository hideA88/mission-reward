package repository

import (
	"github.com/hideA88/mission-reward/pkg/consumer/model/mission"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type MissionRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewMissionRepository(db *sqlx.DB, logger *zap.SugaredLogger) *MissionRepository {
	return &MissionRepository{db: db, logger: logger}
}

func (mr *MissionRepository) GetMission(missionId int64) (*mission.Mission, error) {
	query := `SELECT id, name, mission_type, IFNULL(reset_time, ''), IFNULL(open_mission_id, 0), default_mission_status FROM mission where id = ?;`

	var m = mission.Mission{}
	err := mr.db.Select(&m, query, missionId)
	if err != nil {
		mr.logger.Error(err)
		return nil, err
	}

	return &m, nil
}

func (mr *MissionRepository) SelectTotalCoinMissions() ([]*mission.TotalCoinMission, error) {
	query := `SELECT id, mission_id, size FROM total_coin_mission;`

	missions := make([]*mission.TotalCoinMission, 0)
	err := mr.db.Select(&missions, query)
	if err != nil {
		mr.logger.Error(err)
		return nil, err
	}

	return missions, nil
}
