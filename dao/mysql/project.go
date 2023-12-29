package mysql

import (
	"eee/model"
	"strconv"
)

func GetRightByName(uname string) (right *string, err error) {
	str := "select right_name from u_right where right_id in (select r_id from user where u_name=?)"
	err = db.Get(&right, str, uname)
	return
}

func SetPFinish(id string, finish string) (err error) {
	str := "update project set finish=? where p_id=?"
	_, err = db.Exec(str, finish, id)
	return
}

func GetProList(pagenum, pagesize string) (proList []*model.Project, err error) {
	num, _ := strconv.Atoi(pagenum)
	size, _ := strconv.Atoi(pagesize)
	str := "select * from project limit ?,?"
	err = db.Select(&proList, str, (num-1)*size, pagesize)
	return
}

func GetProByUid(uid string) (proList []*model.MyPro, err error) {
	str := "select p_id,p_name,leader,tag,finish from project where p_id in (select pr_id from support where us_id=?)"
	err = db.Select(&proList, str, uid)
	return
}

func GetProByLeader(leader string) (proList []*model.MyPro, err error) {
	str := "select p_id,p_name,leader,tag,finish from project where leader=?"
	err = db.Select(&proList, str, leader)
	return
}

func CountP() (num *string, err error) {
	str := "select count(*) from project"
	err = db.Get(&num, str)
	return
}

func UpdateP(pro *model.Project) (err error) {
	str := "update project set tag=?,p_name=?,start_time=?,end_time=?,amount=?,input=?,balance=?,description=? where p_id=?"
	_, err = db.Exec(str, pro.Tag, pro.ProjectName, pro.StartTime, pro.EndTime, pro.Amount, pro.Input, pro.Balance, pro.Description, pro.ProjectID)
	return
}

func ProDetails(pid string) (pro model.Project, err error) {
	str := "select * from project where p_id=?"
	err = db.Get(&pro, str, pid)
	return
}

func WorkerList(pid string) (list []*model.WS, err error) {
	str := "select u.u_name,s.startTime,s.endTime from user u inner join support s on u.u_id=s.us_id where s.pr_id=?"
	err = db.Select(&list, str, pid)
	return
}

func RecentOutput() (list []*model.ProM, err error) {
	str := "select amount,input,balance from project order by end_time desc limit 0,7"
	err = db.Select(&list, str)
	return
}

func RecentPros() (list []*model.Project, err error) {
	str := "select * from project order by end_time desc limit 0,7"
	err = db.Select(&list, str)
	return
}

func ProMoney(p_id string) (money model.Money, err error) {
	str := "select p_id,p_name,leader,start_time,end_time,amount,balance,input from project where p_id=?"
	err = db.Get(&money, str, p_id)
	return
}
