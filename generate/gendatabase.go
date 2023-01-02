package generate

import (
	. "sql/generate/utility/generate"
)

func Gendatabase(p string) {
	//make folder
	Makefolder(p)

	//make file/database
	Makefile(p + "/database.db")

}
