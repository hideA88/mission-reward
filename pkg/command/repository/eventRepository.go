package repository

import (
	"context"
	"github.com/hideA88/mission-reward/pkg/command/model/request"
	"github.com/hideA88/mission-reward/pkg/consumer/model/event"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type EventRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewEventRepository(db *sqlx.DB, logger *zap.SugaredLogger) *EventRepository {
	return &EventRepository{db: db, logger: logger}
}

func (er *EventRepository) SaveLoginEvent(ctx context.Context, req *request.LoginReq) (*event.Login, error) {
	er.logger.Info("insert data login :", req)
	re, err := er.db.NamedExecContext(ctx, `INSERT INTO user_login_event (user_id, event_at) VALUES (:user_id, :event_at)`, req)
	if err != nil {
		return nil, err
	}

	lastId, _ := re.LastInsertId()
	return &event.Login{Id: lastId, UserId: req.UserId, EventAt: req.EventAt}, nil
}

func (er *EventRepository) SaveKillMonsterEvent(ctx context.Context, req *request.KillMonsterReq) (*event.KillMonster, error) {
	er.logger.Info("insert data kill monster :", req)
	re, err := er.db.NamedExecContext(ctx,
		`INSERT INTO kill_monster_event (user_id, kill_monster_id, event_at) VALUES (:user_id, :kill_monster_id, :event_at)`,
		req)
	if err != nil {
		return nil, err
	}

	lastId, _ := re.LastInsertId()
	return &event.KillMonster{Id: lastId, UserId: req.UserId, KillMonsterId: req.KillMonsterId, EventAt: req.EventAt}, nil
}

func (er *EventRepository) SaveLevelUpEvent(ctx context.Context, req *request.LevelUpReq) (*event.LevelUp, error) {
	er.logger.Info("insert data level up :", req)
	re, err := er.db.NamedExecContext(ctx,
		`INSERT INTO level_up_event (user_id, user_monster_id, level_up_size, event_at) VALUES (:user_id, :user_monster_id, :level_up_size, :event_at)`,
		req)
	if err != nil {
		return nil, err
	}

	lastId, _ := re.LastInsertId()
	return &event.LevelUp{Id: lastId, UserId: req.UserId, EventAt: req.EventAt}, nil
}
