package model

type Feedback struct {
	ID   int64  `db:"id"`
	Rate int64  `db:"rate"`
	Text string `db:"text"`
}
