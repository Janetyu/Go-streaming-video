package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES(?, ?)")
	if err != nil {
		log.Printf("AddUser Error: %v", err)
		return err
	}

	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOuts, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("GetUser Error: %v", err)
		return "", err
	}

	var pwd string
	stmtOuts.QueryRow(loginName).Scan(&pwd)
	stmtOuts.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("DelUser Error: %v", err)
		return err
	}

	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}
