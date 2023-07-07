package repository

import (
	"github.com/hideA88/mission-reward/pkg/consumer/model/event"
	"github.com/hideA88/mission-reward/pkg/consumer/model/reward"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewUserRepository(db *sqlx.DB, logger *zap.SugaredLogger) *UserRepository {
	return &UserRepository{db: db, logger: logger}
}

func (ur *UserRepository) GetTotalCoin(userId int64) (*reward.UserCoin, error) {
	query := `SELECT user_id, sum(size) AS total_coin FROM user_coin WHERE user_id = ? group by user_id`

	var userCoin = reward.UserCoin{}
	err := ur.db.Get(&userCoin, query, userId)
	if err != nil {
		ur.logger.Error(err)
		return nil, err
	}
	return &userCoin, nil
}

func (ur *UserRepository) GetLastLogin(userId int64) (*event.Login, error) {
	query := `SELECT user_id, MAX(event_at) as event_at  FROM user_login_event WHERE user_id = ? GROUP BY user_id;`

	var userLogin = event.Login{}
	err := ur.db.Get(&userLogin, query, userId)
	if err != nil {
		ur.logger.Error(err)
		return nil, err
	}
	return &userLogin, nil
}
