package mysql

import "eee/model"

func AddProject(pro model.Project) (err error) {
	str := "insert into project values(?,?,?,?,?,?,?,?,?,?,?)"
	_, err = db.Exec(str, pro.ProjectID, pro.ProjectName, pro.Leader, pro.Finish,
		pro.StartTime, pro.EndTime, pro.Amount, pro.Input,
		pro.Balance, pro.Tag, pro.Description)
	return
}
