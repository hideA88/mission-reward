package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/hideA88/mission-reward/cmd"
	pb "github.com/hideA88/mission-reward/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	cClient pb.MissionRewardCommandServiceClient
	qClient pb.MissionRewardQueryServiceClient
)

func main() {
	config, err := cmd.ParseConfig()
	if err != nil {
		fmt.Println("config file parse error.")
		fmt.Println(err)
		return
	}

	fmt.Println("start gRPC Client.")
	address := fmt.Sprintf("%s:%d", config.Server.Address, config.Server.Port)
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	qClient = pb.NewMissionRewardQueryServiceClient(conn)
	cClient = pb.NewMissionRewardCommandServiceClient(conn)

	scanner = bufio.NewScanner(os.Stdin)
exitLoop:
	for {
		fmt.Println("1: send login Request")
		fmt.Println("2: send kill monster Request")
		fmt.Println("3: send level up Request")
		fmt.Println("4: send get user Request")
		fmt.Println("5: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Login()
		case "2":
			Status()
		case "3":
			Status()
		case "4":
			Status()
		case "5":
			fmt.Println("bye.")
			break exitLoop
		}
	}
}

func Login() {
	fmt.Println("Please enter userId")
	scanner.Scan()
	userId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	req := &pb.PostLoginEventRequest{
		UserId:  int64(userId),
		EventAt: timestamppb.Now(),
	}
	res, err := cClient.PostLoginEvent(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func Status() {
	fmt.Println("Please enter userId")
	scanner.Scan()
	userId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	req := &pb.UserStatusRequest{
		UserId:        int64(userId),
		LastRequested: timestamppb.Now(),
	}
	res, err := qClient.UserStatus(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.UserId)
	}
}
