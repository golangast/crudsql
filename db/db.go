package db

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"time"

	"database/sql"

	_ "modernc.org/sqlite"
)

const file string = "db/database.db"

func DbConnection() (*sql.DB, error) {

	//db urls   conn to db      database used
	db, err := sql.Open("sqlite", file)
	if err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	}

	db.Driver()
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	//check if it pings
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Connected to DB %s successfully\n", file)
	return db, nil
} //end of connect
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

type Names struct {
	Id    int    `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id"`
	Fname string `param:"fname" query:"fname" header:"fname" form:"fname" json:"fname" xml:"fname"`
	Lname string `param:"lname" query:"lname" header:"lname" form:"lname" json:"lname" xml:"lname"`
}

func getAllFilenamesdb(efs *embed.FS) (file string, err error) {
	var pathfile string
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		if path == "db/database.db" {
			pathfile = path
		}

		return nil
	}); err != nil {
		return "", err
	}

	return pathfile, nil
}
func Createtable() {
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	_, err = data.Exec("CREATE TABLE IF NOT EXISTS database(id integer PRIMARY KEY AUTOINCREMENT NOT NULL, fname text, lname text)")
	if err != nil {
		panic(err)
	}
}
