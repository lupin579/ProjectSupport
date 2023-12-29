package model

type Project struct {
	ProjectID   uint16 `json:"projectID" db:"p_id"`
	ProjectName string `json:"projectName" db:"p_name"`
	Tag         string `json:"tag" db:"tag"`
	Finish      string `json:"finish" db:"finish"`
	Leader      string `json:"leader" db:"leader"`
	StartTime   string `json:"startTime" db:"start_time"`
	EndTime     string `json:"endTime" db:"end_time"`
	Amount      uint32 `json:"ContractAmount" db:"amount"`
	Input       uint32 `json:"investment" db:"input"`
	Balance     uint32 `json:"TailPayment" db:"balance"`
	Description string `json:"description" db:"description"`
}

type ProM struct {
	Amount  int `db:"amount"`
	Input   int `db:"input"`
	Balance int `db:"balance"`
}

type MyPro struct {
	Pid    int    `db:"p_id" json:"pid"`
	Pname  string `db:"p_name" json:"pname"`
	Leader string `db:"leader" json:"leader"`
	Tag    string `db:"tag" json:"tag"`
	Finish string `db:"finish" json:"finish"`
}

type Money struct {
	ProjectID int    `db:"p_id"`
	ProName   string `db:"p_name"`
	Leader    string `db:"leader"`
	StartTime string `db:"start_time"`
	EndTime   string `db:"end_time"`
	Amount    int    `db:"amount"`
	Balance   int    `db:"balance"`
	Input     int    `db:"input"`
}
