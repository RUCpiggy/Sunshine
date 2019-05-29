package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Insert(db *sql.DB, sql string, args ...interface{}) {
	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err := stmt.Exec(args...)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func Delete(db *sql.DB, sql string, args ...interface{}) {
	stmt, err := db.Prepare(sql)
	checkErr(err)

	res, err := stmt.Exec(args...)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

}

func Update(db *sql.DB, sql string, args ...interface{}) {
	stmt, err := db.Prepare(sql)
	checkErr(err)
	res, err := stmt.Exec(args...)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
