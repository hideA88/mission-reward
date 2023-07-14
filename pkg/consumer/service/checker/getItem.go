package checker

import (
	"context"
	"github.com/hideA88/mission-reward/pkg/consumer/model/message"
	"github.com/hideA88/mission-reward/pkg/consumer/model/mission"
	"github.com/hideA88/mission-reward/pkg/consumer/model/reward"
)

type GetItemChecker struct {
	*CommonMissionChecker
}

func NewGetItem(mc *CommonMissionChecker) *GetItemChecker {
	return &GetItemChecker{
		mc,
	}
}

func (gic *GetItemChecker) Init(ctx context.Context) error {
	return nil
}

func (gic *GetItemChecker) CheckMission(ctx context.Context, gcm *message.GetItem) error {
	gic.logger.Info("receive message: ", gcm)
	gic.checkMission(ctx, gcm.UserId, gcm.EventAt, mission.GET_ITEM, gic._checkMission)
	return nil
}

func (gic *GetItemChecker) _checkMission(m *reward.MissionWithAchieveHistory) (bool, error) {
	//TODO implement
	return false, nil
}
