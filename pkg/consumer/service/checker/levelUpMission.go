package checker

import (
	"context"
	"github.com/hideA88/mission-reward/pkg/consumer/model/message"
)

type LevelUpMissionChecker struct {
	*CommonMissionChecker
}

func NewLevelUpMission(mc *CommonMissionChecker) *LevelUpMissionChecker {
	return &LevelUpMissionChecker{
		mc,
	}
}

func (luc *LevelUpMissionChecker) Serve(ctx context.Context, luCh <-chan *message.LevelUp) {
	for lum := range luCh {
		//TODO implement
		luc.logger.Info("receive level up message:", lum)
	}
}
