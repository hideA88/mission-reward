package checker

import (
	"context"
	"github.com/hideA88/mission-reward/pkg/consumer/model/message"
	"github.com/hideA88/mission-reward/pkg/consumer/model/mission"
	"github.com/hideA88/mission-reward/pkg/consumer/model/reward"
)

type TotalCoinChecker struct {
	*CommonMissionChecker
}

func NewTotalCoin(mc *CommonMissionChecker) *TotalCoinChecker {
	return &TotalCoinChecker{
		mc,
	}
}

func (tcc *TotalCoinChecker) Serve(ctx context.Context, gcCh <-chan *message.GetCoin) {
	missions, err := tcc.mr.SelectTotalCoinMissions()
	if err != nil {
		tcc.logger.Error("select total coin missions error:", err)
		return
	}

	for gcm := range gcCh {
		fn, err := tcc._checkMission(gcm.UserId, missions)
		if err != nil {
			//TODO implement handle error
			tcc.logger.Error("error occuerd:", err)
			continue
		}
		tcc.checkMission(ctx, gcm.UserId, gcm.EventAt, mission.TOTAL_COIN, fn)
	}
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
