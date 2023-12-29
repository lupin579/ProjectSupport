package model

type Support struct {
	UserID       uint16 `json:"userID" db:"us_id"`
	ProID        uint16 `json:"proID" db:"pr_id"`
	StartTime    string `json:"startTime" db:"startTime"`
	EndTime      string `json:"endTime" db:"endTime"`
	SupportState string
}

type WS struct {
	Uname     string `db:"u_name"`
	StartTime string `db:"startTime"`
	EndTime   string `db:"endTime"`
	IsSupport int
}

type TimeChart struct {
	Uid       int     `db:"u_id"`
	Uname     string  `db:"u_name"`
	ProID     *int    `db:"pr_id"`
	StartTime *string `db:"startTime"`
	EndTime   *string `db:"endTime"`
}
