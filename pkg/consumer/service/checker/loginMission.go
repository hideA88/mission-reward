package checker

import (
	"context"
	"github.com/hideA88/mission-reward/pkg/consumer/model/message"
	"github.com/hideA88/mission-reward/pkg/consumer/model/mission"
	"github.com/hideA88/mission-reward/pkg/consumer/model/reward"
)

type LoginMissionChecker struct {
	*CommonMissionChecker
}

func NewLoginMission(mc *CommonMissionChecker) *LoginMissionChecker {
	return &LoginMissionChecker{
		mc,
	}
}

func (lmc *LoginMissionChecker) Serve(ctx context.Context, lgCh <-chan *message.Login) {
	for lgm := range lgCh {
		fn, err := lmc._checkMission(lgm.UserId)
		if err != nil {
			//TODO implement handle error
			lmc.logger.Error("error occuerd:", err)
			continue
		}
		lmc.checkMission(ctx, lgm.UserId, lgm.EventAt, mission.LOGIN, fn)
	}
}

func (lmc *LoginMissionChecker) _checkMission(userId int64) (func(*reward.MissionWithAchieveHistory) (bool, error), error) {
	lastLogin, err := lmc.ur.GetLastLogin(userId)
	if err != nil {
		return nil, err
	}

	return func(uh *reward.MissionWithAchieveHistory) (bool, error) {
		r, err := lmc.receivedAchieve(lastLogin.EventAt, uh)
		return !r, err
	}, nil
}
