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

func (luc *LevelUpMissionChecker) Init(ctx context.Context) error {
	return nil
}

func (luc *LevelUpMissionChecker) CheckMission(ctx context.Context, lum *message.LevelUp) error {
	luc.logger.Info("receive level up message:", lum)
	return nil
}
