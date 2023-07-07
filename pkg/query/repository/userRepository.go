package repository

import (
	"github.com/hideA88/mission-reward/pkg/query/model"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
)

type UserRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewUserRepository(db *sqlx.DB, logger *zap.SugaredLogger) *UserRepository {
	return &UserRepository{db: db, logger: logger}
}

func (ur *UserRepository) GetFullData(userId int64, lastReqTime *time.Time) (*model.UserFullData, error) {
	baseQuery := `
    SELECT id, name, coin, last_login_at 
    FROM user AS u
         LEFT JOIN (SELECT user_id, sum(size) AS coin FROM user_coin WHERE user_id = ? GROUP BY user_id) AS uc 
            ON u.id = uc.user_id
         LEFT JOIN (SELECT user_id, MAX(event_at) AS last_login_at  FROM user_login_event where user_id = ? GROUP BY user_id) AS ul
            ON u.id = ul.user_id
    where id = ? ;`

	var bd = model.UserBaseData{}
	err := ur.db.Get(&bd, baseQuery, userId, userId, userId)
	if err != nil {
		ur.logger.Error(err)
		return nil, err
	}

	itemQuery := `
SELECT ui.item_id, item.name, ui.size 
FROM
	(SELECT user_id, item_id, count(*) as size from user_item where user_id = ? group by item_id, user_id) AS ui
	LEFT JOIN item ON item.id = ui.item_id
`
	var ui = make([]*model.UserItem, 0)
	err = ur.db.Select(&ui, itemQuery, userId)
	if err != nil {
		ur.logger.Error(err)
		return nil, err
	}

	monsterQuery := `
SELECT um.id as monster_id, m.name AS name, IFNULL(ml.level, 0) + 1 as level  
FROM user_monster um
    LEFT JOIN monster AS m ON m.id = um.monster_id
    LEFT JOIN (SELECT user_monster_id, SUM(level_up_size) AS level
           FROM level_up_event group by user_monster_id) AS ml ON um.id = ml.user_monster_id
WHERE um.user_id = ?
`
	var um = make([]*model.UserMonster, 0)
	err = ur.db.Select(&um, monsterQuery, userId)
	if err != nil {
		ur.logger.Error(err)
		return nil, err
	}

	achieveQuery := `
SELECT ua.id AS achieve_id, ua.achieved_at AS achieved_at, m.name AS name 
from user_achieve_mission as ua
    LEFT JOIN mission m on ua.mission_id = m.id
WHERE ua.user_id = ?
`

	var ua = make([]*model.UserAchieve, 0)
	err = ur.db.Select(&ua, achieveQuery, userId)
	if err != nil {
		ur.logger.Error(err)
		return nil, err
	}

	return &model.UserFullData{
		Id:          bd.Id,
		Name:        bd.Name,
		LastLoginAt: bd.LastLoginAt,
		Coin:        bd.Coin,
		Items:       ui,
		Monsters:    um,
		Achieves:    ua,
	}, nil
}
