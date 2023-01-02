package db

import (
	"fmt"
)

func UpdateNames(fname, lname string) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//prepare statement so that no sql injection
	stmt, err := data.Prepare("UPDATE database set fname=?, lname=? where fname=?")
	ErrorCheck(err)

	//execute qeury
	//id := "1"
	res, err := stmt.Exec(fname, lname, fname)
	ErrorCheck(err)

	//used to print rows
	a, err := res.RowsAffected()
	ErrorCheck(err)
	fmt.Println(a, fname)

}
