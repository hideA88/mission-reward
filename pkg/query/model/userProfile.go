package model

import "time"

type UserFullData struct {
	Id          int64          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Coin        int32          `db:"coin" json:"coin"`
	Items       []*UserItem    `db:"items" json:"items"`
	Monsters    []*UserMonster `db:"monsters" json:"monsters"`
	Achieves    []*UserAchieve `db:"achieves" json:"achieves"`
	LastLoginAt *time.Time     `db:"last_login_at" json:"lastLoginAt"`
}

type UserBaseData struct {
	Id          int64      `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Coin        int32      `db:"coin" json:"coin"`
	LastLoginAt *time.Time `db:"last_login_at" json:"lastLoginAt"`
}

type UserItem struct {
	ItemId int64  `db:"item_id"   json:"itemId"`
	Name   string `db:"name" json:"name"`
	Size   int32  `db:"size" json:"size"`
}

type UserMonster struct {
	MonsterId int64  `db:"monster_id"   json:"monsterId"`
	Name      string `db:"name" json:"name"`
	Level     int32  `db:"level" json:"level"`
}

type UserAchieve struct {
	AchieveId  int64      `db:"achieve_id"  json:"achieveId"`
	Name       string     `db:"name"        json:"name"`
	AchievedAt *time.Time `db:"achieved_at" json:"achievedAt"`
}
