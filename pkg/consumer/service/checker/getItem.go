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

func (gic *GetItemChecker) Serve(ctx context.Context, giCh <-chan *message.GetItem) {
	for gcEvent := range giCh {
		gic.logger.Info("receive message: ", gcEvent)
		gic.checkMission(ctx, gcEvent.UserId, gcEvent.EventAt, mission.GET_ITEM, gic._checkMission)
	}

}

func (gic *GetItemChecker) _checkMission(m *reward.MissionWithAchieveHistory) (bool, error) {
	//TODO implement
	return false, nil
}
