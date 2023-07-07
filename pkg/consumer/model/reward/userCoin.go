package reward

type UserCoin struct {
	UserId    int64 `db:"user_id" json:"userId"`
	TotalCoin int   `db:"total_coin" json:"totalCoin"`
}
