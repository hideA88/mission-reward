package repository

import (
	"context"
	"github.com/hideA88/mission-reward/pkg/consumer/model/event"
	"github.com/hideA88/mission-reward/pkg/consumer/model/mission"
	"github.com/hideA88/mission-reward/pkg/consumer/model/reward"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type MissionRewardRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewMissionRewardRepository(db *sqlx.DB, logger *zap.SugaredLogger) *MissionRewardRepository {
	return &MissionRewardRepository{db: db, logger: logger}
}

func (ur *MissionRewardRepository) SelectAchieveHistory(userId int64, missionType mission.Type) ([]*reward.MissionWithAchieveHistory, error) {
	query := `
SELECT * FROM (SELECT m.mission_id           AS  mission_id,
                      m.mission_name         AS  mission_name,
                      m.reward_type          AS  reward_type,
                      m.reward_coin_size     AS  reward_coin_size,
                      m.reward_item_id       AS  reward_item_id,
                      m.reset_time           AS  reset_time,
                      m.open_mission_id      AS  open_mission_id,
                      IFNULL(uam.user_id, 0) AS  user_id,
                      uam.achieved_at        AS  achieved_at,
                      IF(m.default_mission_status = 'blocked' AND uom.id IS NOT NULL, 'open', m.default_mission_status) AS mission_status
              FROM mission_reward AS m
                   LEFT JOIN
                      (SELECT user_id, mission_id, MAX(achieved_at) AS achieved_at
                       FROM user_achieve_mission
                       WHERE user_id = ?
                       GROUP BY user_id, mission_id) AS uam
                   ON m.mission_id = uam.mission_id
              LEFT JOIN (select * from user_open_mission where user_id = ? )AS uom ON m.mission_id = uom.mission_id
              WHERE m.mission_type = ?) AS a
WHERE a.mission_status = 'open'`

	missions := make([]*reward.MissionWithAchieveHistory, 0)
	err := ur.db.Select(&missions, query, userId, userId, missionType)
	if err != nil {
		ur.logger.Error(err)
		return nil, err
	}

	return missions, nil
}

func (ur *MissionRewardRepository) InsertAchieveMissions(ctx context.Context, reqs []*reward.UserAchieveReq) ([]*reward.UserAchieve, error) {
	ur.logger.Info("user achieve req size", len(reqs))

	tx, err := ur.db.Beginx()
	if err != nil {
		return nil, err
	}

	var uas = make([]*reward.UserAchieve, len(reqs))
	for i, req := range reqs {
		ur.logger.Info("insert data user achieve :", req)
		r, err := tx.NamedExecContext(ctx, `INSERT INTO user_achieve_mission (user_id, mission_id, achieved_at) VALUES (:user_id, :mission_id, :achieved_at)`, req)
		if err != nil {
			ur.logger.Error("error insert user_achieve_mission: ", err)
			ur.logger.Error(err)
			err := tx.Rollback()
			if err != nil {
				ur.logger.Error("rollback error: ", err)
				return nil, err
			}
			return nil, err
		}
		aid, _ := r.LastInsertId()

		// open mission
		var om *event.OpenMission = nil
		if req.OpenMissionId > 0 {
			om2, err2 := ur.saveOpenMission(ctx, tx, req, aid)
			if err2 != nil {
				return nil, err
			}
			om = om2
		}

		var qq string
		if req.RewardType == "coin" {
			qq = `INSERT INTO user_coin (user_id, size, user_achieve_mission_id) VALUES (:user_id, :coin_size, :user_achieve_mission_id)`
		} else {
			qq = `INSERT INTO user_item (user_id, item_id, user_achieve_mission_id) VALUES (:user_id, :item_id, :user_achieve_mission_id)`
		}
		rr, err := tx.NamedExecContext(ctx, qq, map[string]interface{}{
			"user_id":                 req.UserId,
			"user_achieve_mission_id": aid,
			"coin_size":               req.RewardCoinSize,
			"item_id":                 req.RewardItemId,
		})
		if err != nil {
			ur.logger.Error("error insert query:", qq)
			ur.logger.Error(err)
			err := tx.Rollback()
			if err != nil {
				ur.logger.Error("rollback error: ", err)
				return nil, err
			}
			return nil, err
		}
		rid, _ := rr.LastInsertId()
		uas[i] = &reward.UserAchieve{
			Id:             aid,
			UserId:         req.UserId,
			MissionId:      req.MissionId,
			AchievedAt:     req.AchievedAt,
			RewardType:     req.RewardType,
			RewardId:       rid,
			RewardCoinSize: req.RewardCoinSize,
			RewardItemId:   req.RewardItemId,
			OpenMission:    om,
		}
	}

	err = tx.Commit()
	if err != nil {
		ur.logger.Error("commit error", err)
		err := tx.Rollback()
		if err != nil {
			ur.logger.Error("rollback error: ", err)
			return nil, err
		}
		return nil, err
	}
	ur.logger.Info("uas size:", len(uas))
	return uas, nil
}

func (ur *MissionRewardRepository) saveOpenMission(ctx context.Context, tx *sqlx.Tx, req *reward.UserAchieveReq, aid int64) (*event.OpenMission, error) {
	query := `INSERT INTO user_open_mission (user_id, mission_id, user_achieve_mission_id) VALUES (:user_id, :mission_id, :achieved_at)`
	r, err := tx.NamedExecContext(ctx, query,
		map[string]interface{}{
			"user_id":                 req.UserId,
			"mission_id":              req.MissionId,
			"user_achieve_mission_id": aid,
		})
	if err != nil {
		ur.logger.Error("error insert user_achieve_mission: ", err)
		ur.logger.Error(err)
		err := tx.Rollback()
		if err != nil {
			ur.logger.Error("rollback error: ", err)
			return nil, err
		}
		return nil, err
	}
	id, _ := r.LastInsertId()
	return &event.OpenMission{
		Id:        id,
		UserId:    req.UserId,
		MissionId: req.OpenMissionId,
		EventAt:   req.AchievedAt,
	}, nil
}
