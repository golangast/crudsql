package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
)

type Data struct {
	Table      string
	Fields     []string
	Types      []string
	Fieldtypes []string
}

// write any template to file
func Writetemplate(temp string, f *os.File, d map[string]string) {
	dbmb := template.Must(template.New("queue").Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
}

// write any template to file
func WritetemplateData(temp string, f *os.File, d Data) {
	funcMap := template.FuncMap{
		"Iterate": func(str []string) []string {
			ls := len(str)
			var i int
			var Items []string
			for i = 0; i < (ls); i++ {
				if i+1 == ls {
					//Items = append(Items, str[i][0:len(str[i])-1])
					Items = append(Items, str[i]+" NOT NULL ")
				} else {
					Items = append(Items, str[i]+" NOT NULL, ")
				}
			}

			return Items
		},
	}
	dbmb := template.Must(template.New("queue").Funcs(funcMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
}

// make any folder
func Makefolder(p string) {
	if err := os.MkdirAll(p, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("~~~~could not create"+p, err)
	} else {
		fmt.Println("Directory " + p + " successfully created with sticky bits and full permissions")
	}
}

// make any file
func Makefile(p string) *os.File {
	file, err := os.Create(p)
	if IsError(err) {
		fmt.Println("error -", err, file)
	}
	return file
}

// error checker
func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

// get all url and title from database
func GetUrlTitle(prop []string) ([]string, []string) {
	var title []string
	var url []string
	for _, ss := range prop {
		title = append(title, TrimColan(ss))
		url = append(url, TrimColanright(ss))

	}
	return title, url
}

func TrimColan(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		return s[:idx]
	}
	return s
}

func TrimColanright(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		id := strings.Replace(s[idx:], ":", "", 1)
		return id
	}
	return s
}

// creates a template file for the app
func Createtemplatefile(f string) {
	mfile, err := os.Create("templates/" + f)
	if IsError(err) {
		fmt.Println("error -", err, mfile)
	}
}

// write to any file
func Writefiles(path, o, n string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	outputs := bytes.Replace(input, []byte(o), []byte(n), -1)
	err = ioutil.WriteFile(path, outputs, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// show data on html template file
func Showdata(path string) {
	o := "<!-- ### -->"

	n := `{{.}}` + "\n" + "<!-- ### -->"

	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(o), []byte(n), -1)
	fmt.Println("file: ", path, " old: ", o, " new: ", n)
	if err = ioutil.WriteFile(path, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func GetTableNameAppData(icb any) Data {
	//get tables
	//tables := Getapptables()

	//get all table's data

	//separate them
	//https://go.dev/play/p/0HrA-jrPZqG
	d := Data{}
	s := reflect.ValueOf(icb).Elem()
	//get table fields, types, and name
	for i := 0; i < s.NumField(); i++ {
		if s.Field(i).Interface() != "" {
			switch s.Type().Field(i).Name[0:1] {
			case "F": //fields
				d.Fields = append(d.Fields, s.Field(i).Interface().(string))
			case "T": //types
				d.Types = append(d.Types, s.Field(i).Interface().(string))
			case "S": //table name
				d.Table = s.Field(i).Interface().(string)
			}
		}
	}

	//add them to frontend
	return d

}
