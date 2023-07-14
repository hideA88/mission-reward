package checker

import (
	"context"
	"github.com/hideA88/mission-reward/pkg/consumer/model/message"
	"github.com/hideA88/mission-reward/pkg/consumer/model/mission"
	"github.com/hideA88/mission-reward/pkg/consumer/model/reward"
)

type TotalCoinChecker struct {
	*CommonMissionChecker
	missions []*mission.TotalCoinMission
}

func NewTotalCoin(mc *CommonMissionChecker) *TotalCoinChecker {
	return &TotalCoinChecker{
		mc,
		nil,
	}
}

func (tcc *TotalCoinChecker) Init(ctx context.Context) error {
	missions, err := tcc.mr.SelectTotalCoinMissions()
	if err != nil {
		tcc.logger.Error("select total coin missions error:", err)
		return err
	}
	tcc.missions = missions

	return nil
}

func (tcc *TotalCoinChecker) CheckMission(ctx context.Context, gcm *message.GetCoin) error {
	fn, err := tcc._checkMission(gcm.UserId, tcc.missions)
	if err != nil {
		//TODO implement handle error
		tcc.logger.Error("error occuerd:", err)
		return err
	}
	tcc.checkMission(ctx, gcm.UserId, gcm.EventAt, mission.TOTAL_COIN, fn)
	return nil
}

func (tcc *TotalCoinChecker) _checkMission(userId int64, missions []*mission.TotalCoinMission) (func(*reward.MissionWithAchieveHistory) (bool, error), error) {
	userCoin, err := tcc.ur.GetTotalCoin(userId)

	if err != nil {
		return nil, err
	}
	return func(uh *reward.MissionWithAchieveHistory) (bool, error) {
		mid := uh.MissionId
		for _, m := range missions {
			if mid == m.MissionId {
				return userCoin.TotalCoin >= m.Size, nil
			}
		}
		return false, nil
	}, nil
}
