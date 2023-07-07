package service

import (
	"context"
	"github.com/hideA88/mission-reward/pkg/command/model/request"
	"github.com/hideA88/mission-reward/pkg/command/repository"
	"github.com/hideA88/mission-reward/pkg/consumer/model/message"
	pb "github.com/hideA88/mission-reward/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MissionRewardCommandServer struct {
	er     *repository.EventRepository
	lgCh   chan<- *message.Login
	mkCh   chan<- *message.KillMonster
	luCh   chan<- *message.LevelUp
	logger *zap.SugaredLogger
	pb.UnimplementedMissionRewardCommandServiceServer
}

func NewMissionRewardCommand(er *repository.EventRepository,
	lgCh chan<- *message.Login,
	mkCh chan<- *message.KillMonster,
	luCh chan<- *message.LevelUp,
	logger *zap.SugaredLogger) *MissionRewardCommandServer {
	return &MissionRewardCommandServer{
		er:     er,
		lgCh:   lgCh,
		mkCh:   mkCh,
		luCh:   luCh,
		logger: logger}
}

func (mc *MissionRewardCommandServer) PostLoginEvent(ctx context.Context, req *pb.PostLoginEventRequest) (*pb.PostLoginEventResponse, error) {
	time := req.EventAt.AsTime()
	event, err := mc.er.SaveLoginEvent(ctx, &request.LoginReq{UserId: req.UserId, EventAt: &time})
	if err != nil {
		mc.logger.Fatalf(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	mc.lgCh <- &message.Login{UserId: event.UserId, EventAt: event.EventAt}
	return &pb.PostLoginEventResponse{
		LoginEventId: event.Id,
		EventAt:      timestamppb.New(*event.EventAt),
	}, nil
}

func (mc *MissionRewardCommandServer) PostKillMonsterEvent(ctx context.Context, req *pb.PostKillMonsterEventRequest) (*pb.PostKillMonsterEventResponse, error) {
	time := req.EventAt.AsTime()
	event, err := mc.er.SaveKillMonsterEvent(ctx, &request.KillMonsterReq{UserId: req.UserId, KillMonsterId: req.TargetMonsterId, EventAt: &time})
	if err != nil {
		mc.logger.Fatalf(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	mc.mkCh <- &message.KillMonster{UserId: event.UserId, EventAt: event.EventAt}
	return &pb.PostKillMonsterEventResponse{
		KillMonsterEventId: event.Id,
		//EventAt:            timestamppb.New(*request.EventAt),
	}, nil
}

func (mc *MissionRewardCommandServer) PostLevelUpEvent(ctx context.Context, req *pb.PostLevelUpEventRequest) (*pb.PostLevelUpEventResponse, error) {
	time := req.EventAt.AsTime()
	event, err := mc.er.SaveLevelUpEvent(ctx, &request.LevelUpReq{UserId: req.UserId, UserMonsterId: req.UserMonsterId, LevelUpSize: int(req.LevelUpSize), EventAt: &time})
	if err != nil {
		mc.logger.Fatalf(err.Error())
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	mc.luCh <- &message.LevelUp{UserId: event.UserId, EventAt: event.EventAt}
	return &pb.PostLevelUpEventResponse{
		LevelUpEventId: event.Id,
		//EventAt:            timestamppb.New(*request.EventAt),
	}, nil
}
