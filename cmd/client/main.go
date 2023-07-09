package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hideA88/mission-reward/pkg"
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
	config, err := pkg.ParseConfig("./configs/config.toml")
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
		fmt.Println("1: send get user data")
		fmt.Println("2: send login Request")
		fmt.Println("3: send kill monster Request")
		fmt.Println("4: send level up Request")
		fmt.Println("5: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Status()
		case "2":
			Login()
		case "3":
			KillMonster()
		case "4":
			LevelUp()
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
		j, _ := json.Marshal(res)
		fmt.Println(fmt.Sprintf("%s", j))
	}
}

func KillMonster() {
	fmt.Println("Please enter userId")
	scanner.Scan()
	userId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Please enter userMonsterId")
	scanner.Scan()
	umId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Please enter targetMonsterId")
	scanner.Scan()
	tmId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	req := &pb.PostKillMonsterEventRequest{
		UserId:          int64(userId),
		UserMonsterId:   int64(umId),
		TargetMonsterId: int64(tmId),
		EventAt:         timestamppb.Now(),
	}
	res, err := cClient.PostKillMonsterEvent(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		j, _ := json.Marshal(res)
		fmt.Println(fmt.Sprintf("%s", j))
	}
}

func LevelUp() {
	fmt.Println("Please enter userId")
	scanner.Scan()
	userId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Please enter userMonsterId")
	scanner.Scan()
	userMonsterId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Please enter level up size")
	scanner.Scan()
	levelUpSize, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	req := &pb.PostLevelUpEventRequest{
		UserId:        int64(userId),
		UserMonsterId: int64(userMonsterId),
		LevelUpSize:   int32(levelUpSize),
		EventAt:       timestamppb.Now(),
	}
	res, err := cClient.PostLevelUpEvent(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		j, _ := json.Marshal(res)
		fmt.Println(fmt.Sprintf("%s", j))
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
		items, _ := json.Marshal(res.Items)
		monsters, _ := json.Marshal(res.Monsters)
		achieves, _ := json.Marshal(res.Achieves)
		fmt.Println(fmt.Sprintf("userId:%#v name:%#v coin:%#v", res.UserId, res.Name, res.Coin))
		fmt.Println(fmt.Sprintf("items: %s", items))
		fmt.Println(fmt.Sprintf("monsters: %s", monsters))
		fmt.Println(fmt.Sprintf("achieves: %s", achieves))
	}
}
