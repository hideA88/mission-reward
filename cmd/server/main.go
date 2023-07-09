package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/hideA88/mission-reward/cmd"
	"github.com/hideA88/mission-reward/pkg"
	crepo "github.com/hideA88/mission-reward/pkg/command/repository"
	cserv "github.com/hideA88/mission-reward/pkg/command/service"
	"github.com/hideA88/mission-reward/pkg/consumer/model/message"
	"github.com/hideA88/mission-reward/pkg/consumer/repository"
	"github.com/hideA88/mission-reward/pkg/consumer/service/checker"
	pb "github.com/hideA88/mission-reward/pkg/grpc"
	qrepo "github.com/hideA88/mission-reward/pkg/query/repository"
	qserv "github.com/hideA88/mission-reward/pkg/query/service"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := pkg.ParseConfig()
	if err != nil {
		fmt.Println("config file parse error.")
		fmt.Println(err)
		return
	}

	logger := pkg.NewLogger(config.Verbose)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	port := config.Server.Port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error("service error: %v", err)
		panic(err)
	}

	istLogger, opts := cmd.NewInterceptorLogger(logger.Desugar())
	gsrv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(istLogger, opts...),
		// TODO Add auth interceptor
	))

	//connect DB
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FTokyo",
		config.DB.User,
		config.DB.Password,
		config.DB.Address,
		config.DB.Port,
		config.DB.Database)
	db, err := sqlx.Open(config.DB.Driver, dataSource)
	if err != nil {
		logger.Fatal("db connection error: ", err)
		defer db.Close()
	}

	lgCh := make(chan *message.Login)
	kmCh := make(chan *message.KillMonster)
	luCh := make(chan *message.LevelUp)
	defer close(lgCh)
	defer close(kmCh)
	defer close(luCh)

	gcCh := make(chan *message.GetCoin)
	giCh := make(chan *message.GetItem)
	omCh := make(chan *message.OpenMission)
	defer close(gcCh)
	defer close(giCh)
	defer close(omCh)

	// wire command
	er := crepo.NewEventRepository(db, logger)
	pb.RegisterMissionRewardCommandServiceServer(gsrv, cserv.NewMissionRewardCommand(er, lgCh, kmCh, luCh, logger))

	// wire query
	qur := qrepo.NewUserRepository(db, logger)
	pb.RegisterMissionRewardQueryServiceServer(gsrv, qserv.NewMissionRewardQuery(qur, logger))

	// wire consumer
	mr := repository.NewMissionRepository(db, logger)
	rr := repository.NewMissionRewardRepository(db, logger)
	ur := repository.NewUserRepository(db, logger)

	mc := checker.NewCommonMission(mr, rr, ur, gcCh, giCh, omCh, logger)
	tc := checker.NewTotalCoin(mc)
	gc := checker.NewGetItem(mc)
	oc := checker.NewOpenMission(mc, lgCh, kmCh, luCh)

	lm := checker.NewLoginMission(mc)
	lu := checker.NewLevelUpMission(mc)
	km := checker.NewKillMonsterMission(mc)

	go lm.Serve(ctx, lgCh)
	go lu.Serve(ctx, luCh)
	go km.Serve(ctx, kmCh)
	go tc.Serve(ctx, gcCh)
	go gc.Serve(ctx, giCh)
	go oc.Serve(ctx, omCh)

	go func() {
		logger.Infof("start gRPC service port: %v", port)
		err := gsrv.Serve(listener)
		if err != nil {
			logger.Error("gRPC service Error:", err)
			return
		}
	}()

	<-ctx.Done()
	logger.Infof("stopping service...")

	_, cancel := context.WithCancel(ctx)
	defer cancel()
	//TODO implement service以下で作成しているゴルーチンを停止

	gsrv.GracefulStop()

	logger.Infof("stopping service... done")
}
