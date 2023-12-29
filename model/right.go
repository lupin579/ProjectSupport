package model

type Right struct {
	RightId   uint16 `db:"right_id"`
	RightName string `db:"right_name"`
	RightDesc string `db:"right_desc"`
}
