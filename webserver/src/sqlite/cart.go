package sqlite

import (
	"database/sql"
	"structure"
)

func AddCart(db *sql.DB, args ...interface{}) bool {
	sql := "INSERT INTO carts(cartID,userID,artworkID) values(?,?,?)"
	Insert(db, sql, args...)
	return true
}

func UpdateChartInfo(db *sql.DB, cartID string, info string, value string) {
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

func MaxCart(db *sql.DB) int {
	sql := "SELECT MAX(cartID) FROM carts"
	var cartID int
	row := db.QueryRow(sql)
	row.Scan(&cartID)
	return cartID
}

func SearchCart(db *sql.DB, judge string) ([]structure.Cart, bool) {

	sql := "SELECT * FROM carts " + judge
	rows, err := db.Query(sql)
	checkErr(err)

	//map_columns := make(map[string]string)
	var carts []structure.Cart

	for rows.Next() {
		var cart structure.Cart

		err = rows.Scan(&cart.CartID, &cart.UserID, &cart.ArtworkID)

		carts = append(carts, cart)
		//println(artwork.ArtworkID)
	}
	if len(carts) != 0 {
		return carts, true
	} else {
		return carts, false
	}

}
