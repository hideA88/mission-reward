package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/hideA88/mission-reward/cmd"
	crepo "github.com/hideA88/mission-reward/pkg/command/repository"
	cserv "github.com/hideA88/mission-reward/pkg/command/service"
	pb "github.com/hideA88/mission-reward/pkg/grpc"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := cmd.ParseConfig()
	if err != nil {
		fmt.Println("config file parse error.")
		fmt.Println(err)
		return
	}

	logger := cmd.NewLogger(config.Verbose)
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
