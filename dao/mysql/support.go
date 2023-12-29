package mysql

import "eee/model"

func Distribute(su *model.Support) (err error) {
	str := "insert into support(us_id,pr_id,startTime,endTime) values(?,?,?,?)"
	_, err = db.Exec(str, su.UserID, su.ProID, su.StartTime, su.EndTime)
	return
}

func Cancel(userName, proID string) (err error) {
	str := "delete from support where pr_id =? and us_id in (select u_id from user where u_name=?)"
	_, err = db.Exec(str, proID, userName)
	return
}

func TimeChart() (cList []*model.TimeChart, err error) {
	str := "select u.u_id,u.u_name,s.startTime,s.pr_id,s.endTime from user u left join support s on u.u_id=s.us_id order by u.u_id"
	err = db.Select(&cList, str)
	return
}

func GetUserName() (nameList []string, err error) {
	str := "select u_name from user order by u_id"
	err = db.Select(&nameList, str)
	return
}

func CountT() (num int, err error) {
	str := "select count(*) from (select t.u_id from (select u.u_id,u.u_name,s.startTime,s.pr_id,s.endTime from user u left join support s on u.u_id=s.us_id order by u.u_id) t group by t.u_id) c"
	err = db.Get(&num, str)
	return
}
