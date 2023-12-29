package mysql

import (
	"eee/model"
	"strconv"
)

type WorkTime struct {
	User string `json:"name" db:"worker"`
	Time int    `json:"value" db:"sum_time"`
}

func MissionList(uname, pagenum, pagesize string) (missionList []*model.Mission, err error) {
	str := "select * from mission where worker=? limit ?,?"
	num, _ := strconv.Atoi(pagenum)
	size, _ := strconv.Atoi(pagesize)
	err = db.Select(&missionList, str, uname, num*size, pagesize)
	return
}

func SetFinish(id string, finish string) (err error) {
	str := "update mission set finish=? where m_id=?"
	_, err = db.Exec(str, finish, id)
	return
}

func UpdateMission(mission *model.Mission) (err error) {
	str := "update mission set m_name=?,startTime=?,endTime=?,tag=?,leader=?,project=?,description=? where m_id=?"
	_, err = db.Exec(str, mission.MissionName, mission.StartTime,
		mission.EndTime, mission.Tag, mission.Leader,
		mission.ProjectId, mission.Description, mission.MissionId)
	return
}

func AddMission(mission *model.Mission) (err error) {
	str := "insert into mission(m_id,m_name,startTime,endTime,tag,leader,project,description,worker) values(?,?,?,?,?,?,?,?,?)"
	_, err = db.Exec(str, mission.MissionId, mission.MissionName,
		mission.StartTime, mission.EndTime, mission.Tag, mission.Leader,
		mission.ProjectId, mission.Description, mission.Worker)
	return
}

func Count(uname string) (num *string, err error) {
	str := "select count(*) from mission where worker=?"
	err = db.Get(&num, str, uname)
	return
}

func GetWorkTime(m_id string) (worktime string, err error) {
	str := "select worktime from mission where m_id=?"
	err = db.Get(&worktime, str, m_id)
	return
}

func PostWorkTime(realTime int, m_id string) error {
	str := "update mission set worktime=? where m_id=?"
	_, err := db.Exec(str, realTime, m_id)
	return err
}

func TimePercent(pro string) (workTime []*WorkTime, err error) {
	str := "select sum(worktime) as sum_time,worker from mission where project=? group by worker"
	err = db.Select(&workTime, str, pro)
	return
}

func GetMP(pro string) (mfList []*model.MF, err error) {
	str := "select u_name,m_id,m_name,tag,finish,start_time,end_time from ((select u_id 'u_id',u_name 'u_name' from user where u_id in (select us_id from support where pr_id = ?)) u inner join (select m_id 'm_id',m_name 'm_name',startTime 'start_time',endTime 'end_time',tag,worker,finish from mission where project=?) m on u.u_name=m.worker)"
	err = db.Select(&mfList, str, pro, pro)
	return
}
