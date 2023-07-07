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

func (kmc *KillMonsterMissionChecker) Serve(ctx context.Context, kmCh <-chan *message.KillMonster) {
	for kms := range kmCh {
		//TODO implement
		kmc.logger.Info("receive kill monster message:", kms)
	}
}
