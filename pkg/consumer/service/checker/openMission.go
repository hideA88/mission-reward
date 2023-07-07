package checker

import (
	"context"

	"github.com/hideA88/mission-reward/pkg/consumer/model/message"
	"github.com/hideA88/mission-reward/pkg/consumer/model/mission"
)

type OpenMissionChecker struct {
	*CommonMissionChecker
	lgCh chan<- *message.Login
	kmCh chan<- *message.KillMonster
	luCh chan<- *message.LevelUp
}

func NewOpenMission(mc *CommonMissionChecker,
	lgCh chan<- *message.Login,
	kmCh chan<- *message.KillMonster,
	luCh chan<- *message.LevelUp,
) *OpenMissionChecker {
	return &OpenMissionChecker{
		mc,
		lgCh, kmCh, luCh,
	}
}

func (oc *OpenMissionChecker) Serve(ctx context.Context, omCh <-chan *message.OpenMission) {
	for omEvent := range omCh {
		oc.logger.Info("receive meesage: ", omEvent)
		oc.handleEvent(ctx, omEvent)
	}
}

func (oc *OpenMissionChecker) handleEvent(ctx context.Context, omEvent *message.OpenMission) {
	m, err := oc.mr.GetMission(omEvent.MissionId)
	if err != nil {
		oc.logger.Error(err)
		return
	}

	switch mission.Type(m.MissionType) {
	case mission.GET_ITEM:
		oc.giCh <- &message.GetItem{
			UserId:  omEvent.UserId,
			EventAt: omEvent.EventAt,
		}
	case mission.TOTAL_COIN:
		oc.gcCh <- &message.GetCoin{
			UserId:  omEvent.UserId,
			EventAt: omEvent.EventAt,
		}
	case mission.LOGIN:
		oc.lgCh <- &message.Login{
			UserId:  omEvent.UserId,
			EventAt: omEvent.EventAt,
		}
	case mission.KILL_MONSTER:
		oc.kmCh <- &message.KillMonster{
			UserId:  omEvent.UserId,
			EventAt: omEvent.EventAt,
		}
	case mission.LEVEL_UP:
		oc.luCh <- &message.LevelUp{
			UserId:  omEvent.UserId,
			EventAt: omEvent.EventAt,
		}
	}
}
