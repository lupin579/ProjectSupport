package model

type Mission struct {
	MissionId   uint16 ` db:"m_id"`
	MissionName string ` db:"m_name"`
	StartTime   string ` db:"startTime"`
	EndTime     string ` db:"endTime"`
	Tag         string ` db:"tag"`
	Finish      string ` db:"finish"`
	Leader      string ` db:"leader"`
	ProjectId   uint16 ` db:"project"`
	Description string ` db:"description"`
	WorkTime    uint16 ` db:"worktime"`
	Worker      string ` db:"worker"`
}

type MF struct {
	Uname     string `json:"u_name" db:"u_name"`
	Mid       int    `json:"m_id" db:"m_id"`
	Mname     string `json:"m_name" db:"m_name"`
	StartTime string `json:"startTime" db:"start_time"`
	EndTime   string `json:"endTime" db:"end_time"`
	Tag       string `json:"tag" db:"tag"`
	Finish    string `json:"finish" db:"finish"`
}
