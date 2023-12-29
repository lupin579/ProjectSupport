package model

type User struct {
	Uid      uint   `json:"uid" db:"u_id"`
	Uname    string `json:"uname" db:"u_name" `
	Mobile   string `json:"mobile" db:"mobile" `
	Email    string `json:"email" db:"email" `
	Password string `json:"password" db:"password" `
	RightId  uint16 `db:"r_id" `
}

type UserExtra struct {
	Uid         uint   `db:"e_id"`
	Uname       string `db:"u_name"`
	WorkType    string `db:"work_type"`
	Onboarding  string `db:"onboarding"`
	Mobile      string `json:"mobile" db:"mobile" `
	Email       string `json:"email" db:"email" `
	Level       string `json:"level" db:"level"`
	WorkingTime int
	Projects    int `db:"pros"`
}

type UserAll struct {
	WorkType   string  `db:"work_type"`
	Onboarding string  `db:"onboarding"`
	Level      float64 `db:"level"`
	User
}
