package mysql

import (
	"eee/model"
	"strconv"
)

type Fp struct {
	Uid         int    `db:"u_id"`
	Uname       string `db:"u_name"`
	WorkType    string `db:"work_type"`
	Onboarding  string `db:"onboarding"`
	Pros        int    `db:"pros"`
	WorkingTime int
}

type FpList struct {
	FreeList []*Fp
	Cur      string
	Total    string
}

func UserAdd(uid, uname, workType, onboarding, mobile, email, password string, level float64) error {
	str := "insert into user(u_id,u_name,mobile,email,password) values(?,?,?,?,?)"
	if _, err := db.Exec(str, uid, uname, mobile, email, password); err != nil {
		return err
	}
	str1 := "insert into extra_user(e_id,work_type,onboarding,level) values(?,?,?,?)"
	if _, err := db.Exec(str1, uid, workType, onboarding, level); err != nil {
		return err
	}
	return nil
}

func GetAll(pagenum, pagesize string) (userList []*model.UserExtra, err error) {
	num, _ := strconv.Atoi(pagenum)
	size, _ := strconv.Atoi(pagesize)
	str := "select us.e_id 'e_id',us.u_name 'u_name',us.work_type 'work_type',us.onboarding 'onboarding',us.level 'level',us.mobile 'mobile',us.email 'email',ps.pros 'pros' from (select e.e_id 'e_id',e.level 'level',u.u_name 'u_name',e.work_type 'work_type',e.onboarding 'onboarding',u.mobile 'mobile',u.email 'email' from user u inner join extra_user e on e.e_id=u.u_id order by e_id) us inner join (select u_id,u_name,count(p_id) 'pros' from (select u.u_id 'u_id',s.pr_id 'p_id',u.u_name  from support s right join user u on s.us_id=u.u_id) su group by u_id) ps on us.e_id=ps.u_id order by us.e_id limit ?,?"
	db.Select(&userList, str, (num-1)*size, pagesize)
	return
}

func UserList(pagenum, pagesize, pro string) (userList []*model.UserExtra, err error) {
	num, _ := strconv.Atoi(pagenum)
	size, _ := strconv.Atoi(pagesize)
	str := "select d.e_id,d.u_name,d.work_type,d.level,d.onboarding,d.email,d.mobile,p.count 'pros' from (select e.e_id 'e_id',u.u_name 'u_name',e.work_type 'work_type',e.level 'level',e.onboarding 'onboarding',u.email 'email',u.mobile 'mobile' from user u inner join extra_user e on e.e_id=u.u_id where u.u_id in (select us_id from support where pr_id=1) order by e_id) d inner join (select us_id,count(*) 'count' from support where us_id in (select us_id from support where pr_id=?) group by us_id) p on d.e_id=p.us_id limit ?,?"
	err = db.Select(&userList, str, pro, (num-1)*size, size)
	return
}

func ChangePassword(uname, pwd string) error {
	str := "update user set password=? where u_name=?"
	_, err := db.Exec(str, pwd, uname)
	return err
}

func CountH() (num *string, err error) {
	str := "select count(*) from user"
	err = db.Get(&num, str)
	return
}

func GetEmailByUname(uname string) (email string, err error) {
	str := "select email from user where u_name=?"
	err = db.Get(&email, str, uname)
	return
}

func GetFreeWorker(pro string) (list []*Fp, err error) {
	str := "select u.u_id 'u_id',u.u_name 'u_name',u.work_type 'work_type',u.onboarding 'onboarding',c.pros 'pros' from  (select u.u_id 'u_id', u.u_name 'u_name',e.work_type 'work_type',e.onboarding 'onboarding'     from user u              inner join extra_user e on u.u_id=e.e_id) u        right join    (select u.u_id ,count(s.pr_id) pros     from user u              left join support s on u.u_id=s.us_id  group by u.u_id having count(s.pr_id)<3) c  on u.u_id=c.u_id where u.u_id not in (select us_id from support where pr_id=?)"
	err = db.Select(&list, str, pro)
	return
}
