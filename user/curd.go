package user

import (
	"budgetBook/dbops"
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

func AddUser(u *user) error {
	stmtIns, err := dbops.DBConn.Prepare("insert into user (username, pwd, lastLogin, budget) values (?,?,?,?)")
	if err != nil {
		panic(err)
	}
	u.pwd = crypt(u.pwd)
	_, err = stmtIns.Exec(u.username, u.pwd, u.lastLogin, u.budget)
	if err != nil {
		panic(err)
	}
	defer stmtIns.Close()
	return nil
}

func DeleteUser(username, pwd string) error {
	pwd = crypt(pwd)
	stmtDel, err := dbops.DBConn.Prepare("DELETE FROM user WHERE username=? AND pwd=?")
	if err != nil {
		panic(err)
		return err
	}
	_, err = stmtDel.Exec(username, pwd)
	if err != nil {
		panic(err)
		return err
	}
	defer stmtDel.Close()
	return nil
}

func SelectUser(username, pwd string) (*user, error) {
	stmtSel, err := dbops.DBConn.Prepare("select id,budget from user where username=? and pwd=?")
	stmtUpdate, err := dbops.DBConn.Prepare("update user set lastLogin=? where username=?")
	if err != nil {
		panic(err)
		return nil, err
	}
	var id int
	var budget string
	pwd = crypt(pwd)
	err = stmtSel.QueryRow(username, pwd).Scan(&id, &budget)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		panic(err)
		return nil, err
	}
	rightNow := time.Now()
	ctime := rightNow.Format("Jan 02 2006, 15:04:05")
	budgetDecimal, _ := decimal.NewFromString(budget)
	res := &user{username: username, pwd: pwd, budget: budgetDecimal, lastLogin: ctime}
	stmtUpdate.Exec(ctime, username)
	defer stmtUpdate.Close()
	defer stmtSel.Close()
	return res, nil
}

