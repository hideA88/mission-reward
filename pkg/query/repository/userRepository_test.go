package repository

import (
	"github.com/hideA88/mission-reward/pkg"
	"github.com/hideA88/mission-reward/pkg/query/model"
	"github.com/hideA88/mission-reward/pkg/test"
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

type args struct {
	userId      int64
	lastReqTime *time.Time
}

func NewTestUserRepository() *UserRepository {
	config, err := test.NewTestConfig()
	if err != nil {
		panic(err)
	}
	logger := pkg.NewLogger(config.General.Verbose)
	db := test.NewTestDbConn(config, logger)
	ur := &UserRepository{
		db:     db,
		logger: logger,
	}
	return ur
}

func TestUserRepository_GetFullData(t *testing.T) {
	ur := NewTestUserRepository()
	lq := time.Date(2023, 7, 1, 12, 0, 0, 0, time.Local)
	a1Date := time.Date(2023, 6, 30, 10, 0, 0, 0, time.Local)
	a2Date := time.Date(2023, 7, 1, 10, 0, 0, 0, time.Local)

	tests := []struct {
		name       string
		args       args
		want       *model.UserFullData
		wantErr    bool
		wantErrMeg string
	}{
		// TODO: Add test cases.
		{name: "invalid user id", args: args{userId: 5, lastReqTime: &lq}, want: nil, wantErr: true, wantErrMeg: "sql: no rows in result set"},
		{name: "user all data", args: args{userId: 1, lastReqTime: &lq},
			want: &model.UserFullData{Id: 1, Name: "taro", Coin: 1800,
				Items:    []*model.UserItem{},
				Monsters: []*model.UserMonster{&model.UserMonster{MonsterId: 1, Name: "pochi", Level: 1}, &model.UserMonster{MonsterId: 2, Name: "monsterA", Level: 1}},
				Achieves: []*model.UserAchieve{&model.UserAchieve{AchieveId: 1, Name: "daily_login_reward", AchievedAt: &a1Date},
					&model.UserAchieve{AchieveId: 2, Name: "init", AchievedAt: &a2Date}},
				LastLoginAt: nil,
			},
			wantErr: false, wantErrMeg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ur.GetFullData(tt.args.userId, tt.args.lastReqTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				assert.Error(t, err, tt.wantErrMeg)
			}
			assert.DeepEqual(t, got, tt.want)
		})
	}
}
