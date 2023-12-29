package mysql

func GetRightByUserName(uname string) (right *string, err error) {
	str := "select right_name from u_right where right_id in (select r_id from user where u_name=?)"
	err = db.Get(&right, str, uname)
	return
}

func GetPassword(uname string) (password *string, err error) {
	str := "select password from user where u_name=?"
	err = db.Get(&password, str, uname)
	return
}

func IsExist(uname string) (exist bool, err error) {
	var dest *string
	str := "select u_id from user where u_name=?"
	err = db.Get(&dest, str, uname)
	if err != nil {
		return false, err
	}
	return true, nil
}
