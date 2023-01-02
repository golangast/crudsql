package db

import (
	"fmt"
)

func Getdb() []Names {
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		id     int
		fname  string
		lname  string
		Namess []Names //used to store all users
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from database")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &fname, &lname)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i, id, fname, lname)

		//store into memory
		u := Names{Id: id, Fname: fname, Lname: lname}
		Namess = append(Namess, u)

	}
	//close everything
	defer rows.Close()
	defer data.Close()

	//Createfieldtable(DBFieldss)

	return Namess

}
