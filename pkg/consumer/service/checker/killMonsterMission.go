package checker

import (
	"context"
	"github.com/hideA88/mission-reward/pkg/consumer/model/message"
)

type KillMonsterMissionChecker struct {
	*CommonMissionChecker
}

func NewKillMonsterMission(mc *CommonMissionChecker) *KillMonsterMissionChecker {
	return &KillMonsterMissionChecker{
		mc,
	}
}

func (kmc *KillMonsterMissionChecker) Init(ctx context.Context) error {
	return nil
}

func (kmc *KillMonsterMissionChecker) CheckMission(ctx context.Context, kms *message.KillMonster) error {
	kmc.logger.Info("receive kill monster message:", kms)
	return nil
}
