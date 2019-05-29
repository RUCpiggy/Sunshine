package sqlite

import "database/sql"

func AddCart(db *sql.DB, args ...interface{}) {
	sql := "INSERT INTO carts(cartID,userID,artworkID) values(?,?,?)"
	Insert(db, sql, args...)
}

func UpdateCartInfo(db *sql.DB, cartID string, info string, value string) {
	sql := ""
	switch info {
	case "userID":
		sql = "UPDATE carts SET userID=? WHERE cartID=?"
	case "artworkID":
		sql = "UPDATE carts SET artworkID=? WHERE cartID=?"
	}
	Update(db, sql, value, cartID)
}

func DeleteCart(db *sql.DB, cartID string) {
	sql := "DELETE FROM carts WHERE cartID=?"
	Delete(db, sql, cartID)
}
