package db

import (
	"fmt"
)

func DeleteNames(fname, lname string) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//prepare statement so that no sql injection
	stmt, err := data.Prepare("DELETE FROM database WHERE fname=?")
	ErrorCheck(err)

	//execute qeury
	res, err := stmt.Exec(fname)
	ErrorCheck(err)

	//used to print rows
	a, err := res.RowsAffected()
	ErrorCheck(err)
	fmt.Println(a, fname)

}
