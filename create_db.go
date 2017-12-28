package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"encoding/csv"
	"strings"
	"io/ioutil"
	"io"
	"fmt"
)


func main() {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)


	//sqlStmt := "create table foo (id integer not null primary key, name text);"
	sqlStmt := "CREATE TABLE T (tnr TEXT NOT NULL PRIMARY KEY, tname TEXT, farbe TEXT, gewicht INTEGER);" +
		"CREATE TABLE L (lnr TEXT NOT NULL PRIMARY KEY, lname TEXT, status INTEGER, ort TEXT);" +
		"CREATE TABLE P (pnr TEXT NOT NULL PRIMARY KEY, pname TEXT, ort TEXT);" +
		"CREATE TABLE LTP (tnr TEXT, lnr TEXT, pnr TEXT, menge INTEGER, primary key (tnr,lnr,pnr));"

	_, err = db.Exec(sqlStmt)
	check(err)

	data_t, err := ioutil.ReadFile("./csv/T.csv")
	check(err)
	r_t := csv.NewReader(strings.NewReader(string(data_t)))
	for {
		record, err := r_t.Read()
		if err == io.EOF {
			break
		}
		check(err)

		//fmt.Println(reflect.TypeOf(record))
		fmt.Println("RecordRow ", record)

		stmt, err := db.Prepare("INSERT INTO t(tnr,tname,farbe,gewicht) VALUES(?,?,?,?)")
		check(err)

		stmt.Exec(record[0], record[1], record[2], record[3])

	}

	data_l, err := ioutil.ReadFile("./csv/L.csv")
	check(err)
	//fmt.Print(string(dat))

	l_t := csv.NewReader(strings.NewReader(string(data_l)))

	for {
		record, err := l_t.Read()
		if err == io.EOF {
			break
		}
		check(err)

		//fmt.Println(reflect.TypeOf(record))
		fmt.Println("RecordRow ", record)

		stmt, err := db.Prepare("INSERT INTO l(lnr,lname,status,ort) VALUES(?,?,?,?)")
		check(err)

		stmt.Exec(record[0], record[1], record[2], record[3])

	}

	data_p, err := ioutil.ReadFile("./csv/P.csv")
	check(err)
	//fmt.Print(string(dat))

	r_p := csv.NewReader(strings.NewReader(string(data_p)))

	for {
		record, err := r_p.Read()
		if err == io.EOF {
			break
		}
		check(err)

		//fmt.Println(reflect.TypeOf(record))
		//fmt.Println("RecordRow ", record)

		stmt, err := db.Prepare("INSERT INTO p(pnr,pname,ort) VALUES(?,?,?)")
		check(err)

		stmt.Exec(record[0], record[1], record[2])

	}

	data_ltp, err := ioutil.ReadFile("./csv/LTP.csv")
	check(err)
	//fmt.Print(string(dat))

	r_ltp := csv.NewReader(strings.NewReader(string(data_ltp)))

	for {
		record, err := r_ltp.Read()
		if err == io.EOF {
			break
		}
		check(err)

		//fmt.Println(reflect.TypeOf(record))
		//fmt.Println("RecordRow ", record)

		stmt, err := db.Prepare("INSERT INTO ltp(tnr,lnr,pnr,menge) VALUES(?,?,?,?)")
		check(err)

		stmt.Exec(record[0], record[1], record[2], record[3])

	}
}