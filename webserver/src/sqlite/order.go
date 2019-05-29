package sqlite

import "database/sql"

func AddOrder(db *sql.DB, args ...interface{}) {
	sql := "INSERT INTO orders(orderID,ownerID,sum,timeCreated) values(?,?,?,?)"
	Insert(db, sql, args...)
}

func DeleteOrder(db *sql.DB, orderID string) {
	sql := "DELETE FROM orders WHERE userID=?"
	Delete(db, sql, orderID)
}

func UpdateOrder(db *sql.DB, orderID string, info string, value string) {
	sql := ""
	switch info {
	case "ownerID":
		sql = "UPDATE orders SET ownerID=? WHERE userID=?"
	case "sum":
		sql = "UPDATE orders SET sum=? WHERE userID=?"
	case "timeCreated":
		sql = "UPDATE orders SET timeCreated=? WHERE userID=?"
	}
	Update(db, sql, value, orderID)
}
