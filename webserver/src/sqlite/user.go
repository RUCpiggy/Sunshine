package sqlite

import (
	"database/sql"
	"structure"
)

func AddUser(db *sql.DB, args ...interface{}) {
	sql := "INSERT INTO users(userID,name,email,password,tel,address,balance) values(?,?,?,?,?,?,?)"
	Insert(db, sql, args...)
}

func UpdateUserInfo(db *sql.DB, userID string, info string, value string) {
	sql := ""
	switch info {
	case "name":
		sql = "UPDATE users SET name=? WHERE userID=?"
	case "email":
		sql = "UPDATE users SET email=? WHERE userID=?"
	case "tel":
		sql = "UPDATE users SET tel=? WHERE userID=?"
	case "address":
		sql = "UPDATE users SET address=? WHERE userID=?"
	case "password":
		sql = "UPDATE users SET password=? WHERE userID=?"
	case "balance":
		sql = "UPDATE users SET balance=? WHERE userID=?"
	}
	Update(db, sql, value, userID)
}

func DeleteUser(db *sql.DB, userID string) {
	sql := "DELETE FROM users WHERE userID=?"
	Delete(db, sql, userID)
}

func SearchUser(db *sql.DB, name string) (structure.User, bool) {
	sql := "SELECT userID,name,password,email,tel,address,balance FROM users WHERE name=" + "\"" + name + "\""
	var user structure.User
	row := db.QueryRow(sql)
	row.Scan(&user.UserID, &user.Name, &user.Password, &user.Email, &user.Tel, &user.Address, &user.Balance)
	if user.UserID != 0 {
		return user, true
	} else {
		return user, false
	}

}

func MaxUserID(db *sql.DB) int {
	sql := "SELECT MAX(userID) FROM users"
	var userID int
	row := db.QueryRow(sql)
	row.Scan(&userID)
	return userID

}

func SearchUserByUserID(db *sql.DB, userID string) (structure.User, bool) {
	sql := "SELECT userID,name,password,email,tel,address,balance FROM users WHERE userID=" + "\"" + userID + "\""
	var user structure.User
	row := db.QueryRow(sql)
	row.Scan(&user.UserID, &user.Name, &user.Password, &user.Email, &user.Tel, &user.Address, &user.Balance)
	if user.UserID != 0 {
		return user, true
	} else {
		return user, false
	}

}
