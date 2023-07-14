package checker

import (
	"context"
	"github.com/hideA88/mission-reward/pkg/consumer/model/message"
	"github.com/hideA88/mission-reward/pkg/consumer/model/mission"
	"github.com/hideA88/mission-reward/pkg/consumer/model/reward"
	"github.com/hideA88/mission-reward/pkg/consumer/repository"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"sync"
	"time"
)

type CommonMissionChecker struct {
	mr     *repository.MissionRepository
	rr     *repository.MissionRewardRepository
	ur     *repository.UserRepository
	gcCh   chan<- *message.GetCoin
	giCh   chan<- *message.GetItem
	omCh   chan<- *message.OpenMission
	p      *cron.Parser
	logger *zap.SugaredLogger
}

func NewCommonMission(mr *repository.MissionRepository,
	rr *repository.MissionRewardRepository,
	ur *repository.UserRepository,
	gcCh chan<- *message.GetCoin,
	giCh chan<- *message.GetItem,
	omCh chan<- *message.OpenMission,
	logger *zap.SugaredLogger) *CommonMissionChecker {
	p := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	return &CommonMissionChecker{
		mr:     mr,
		rr:     rr,
		ur:     ur,
		p:      &p,
		gcCh:   gcCh,
		giCh:   giCh,
		omCh:   omCh,
		logger: logger,
	}
}

func (msc *CommonMissionChecker) checkMission(ctx context.Context, userId int64,
	eventAt *time.Time, missionType mission.Type, checkMissionFn func(*reward.MissionWithAchieveHistory) (bool, error)) {
	rewardHistories, err := msc.rr.SelectAchieveHistory(userId, missionType)
	if err != nil {
		msc.logger.Error(err)
		return
	}

	var reqs = make([]*reward.UserAchieveReq, 0)
	for _, rh := range rewardHistories {
		received, err := msc.receivedAchieve(eventAt, rh)
		if err != nil {
			msc.logger.Error(err)
			return
		}
		if received {
			continue
		}
		r, err := checkMissionFn(rh)
		if err != nil {
			msc.logger.Error(err)
			return
		}

		if r {
			req := reward.UserAchieveReq{
				UserId:         userId,
				MissionId:      rh.MissionId,
				AchievedAt:     eventAt,
				RewardType:     rh.RewardType,
				RewardCoinSize: rh.RewardCoinSize,
				RewardItemId:   rh.RewardItemId,
				OpenMissionId:  rh.OpenMissionId,
			}
			reqs = append(reqs, &req)
		}
	}

	userAchieves, err := msc.rr.InsertAchieveMissions(ctx, reqs)
	if err != nil {
		return
	}

	for _, ua := range userAchieves {
		if ua.OpenMission != nil {
			msc.omCh <- &message.OpenMission{UserId: ua.UserId, MissionId: ua.OpenMission.MissionId, EventAt: ua.AchievedAt}
		}

		msc.logger.Info("new user achieves", ua)
		if ua.RewardType == string(reward.COIN) {
			msc.gcCh <- &message.GetCoin{
				UserId:  ua.UserId,
				EventAt: ua.AchievedAt,
			}
		} else {
			msc.giCh <- &message.GetItem{
				UserId:  ua.UserId,
				EventAt: ua.AchievedAt,
			}
			msc.logger.Info("get Item achieve", ua)
		}
	}
}

// TODO implement resetが来てないから的なのがわかるような返り値にしたほうがよさそう
func (msc *CommonMissionChecker) receivedAchieve(eventAt *time.Time, ah *reward.MissionWithAchieveHistory) (bool, error) {
	msc.logger.Info("check achieve history:", ah)
	if ah.AchievedAt == nil {
		return false, nil
	}
	if ah.ResetTime == "" {
		return true, nil
	}

	s, err := msc.p.Parse(ah.ResetTime)
	if err != nil {
		msc.logger.Error(err)
		return false, err
	}
	rt := s.Next(*ah.AchievedAt)
	msc.logger.Info("next reset time:", rt)
	if !eventAt.Before(rt) {
		return false, nil
	} else {
		return true, nil
	}
}

//TODO implement 本当はMessage型とかの制約をつけたい

type Checker[T any] interface {
	Init(ctx context.Context) error
	CheckMission(ctx context.Context, data *T) error
}

func ServeChecker[T any](ctx context.Context, wCh <-chan *T, checker Checker[T], gStopCh <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	checker.Init(ctx)
	for {
		select {
		case wd := <-wCh:
			checker.CheckMission(ctx, wd)
			break
		case _ = <-gStopCh:
			//TODO implement openMissionをしてさらにループするケースがあるので、永続化して次回起動時にキューにつめるなどをしたほうがよさそう
			l := len(wCh)
			if l > 0 {
				wd := <-wCh
				checker.CheckMission(ctx, wd)
			}
			wg.Done()
			return
		}
	}
}
