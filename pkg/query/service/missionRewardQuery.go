package service

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/hideA88/mission-reward/pkg/grpc"
	"github.com/hideA88/mission-reward/pkg/query/repository"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MissionRewardQueryServer struct {
	userRepository *repository.UserRepository
	logger         *zap.SugaredLogger
	pb.UnimplementedMissionRewardQueryServiceServer
}

func NewMissionRewardQuery(ur *repository.UserRepository, logger *zap.SugaredLogger) *MissionRewardQueryServer {
	return &MissionRewardQueryServer{userRepository: ur, logger: logger}
}

func (qs *MissionRewardQueryServer) UserStatus(ctx context.Context, req *pb.UserStatusRequest) (*pb.UserStatusResponse, error) {
	lastReqTime := req.LastRequested.AsTime()
	ufd, err := qs.userRepository.GetFullData(req.UserId, &lastReqTime)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("not found user data. userId: %d", req.UserId))
		}
		qs.logger.Error(err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("internal error."))
	}

	items := make([]*pb.Item, len(ufd.Items))
	for i, it := range ufd.Items {
		items[i] = &pb.Item{ItemId: it.ItemId, Name: it.Name, Size: it.Size}
	}

	ms := make([]*pb.Monster, len(ufd.Monsters))
	for i, it := range ufd.Monsters {
		ms[i] = &pb.Monster{MonsterId: it.MonsterId, Name: it.Name, Level: it.Level}
	}

	as := make([]*pb.Achieve, len(ufd.Achieves))
	for i, it := range ufd.Achieves {
		as[i] = &pb.Achieve{AchieveId: it.AchieveId, Name: it.Name, AchievedAt: timestamppb.New(*it.AchievedAt)}
	}

	var lla *timestamp.Timestamp = nil
	if ufd.LastLoginAt != nil {
		lla = timestamppb.New(*ufd.LastLoginAt)
	}

	return &pb.UserStatusResponse{
		UserId:      ufd.Id,
		Name:        ufd.Name,
		Coin:        ufd.Coin,
		Items:       items,
		Monsters:    ms,
		Achieves:    as,
		LastLoginAt: lla,
	}, nil
}
