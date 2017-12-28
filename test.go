package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"encoding/csv"
	"strings"
	"io/ioutil"
	"fmt"
	"io"
)

func check(e error) {
	if e != nil {
		log.Printf("%q: %s\n", e)
		panic(e)
	}
}

func main() {
	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)


	//sqlStmt := "create table foo (id integer not null primary key, name text);"
	sqlStmt := "CREATE TABLE t (tnr TEXT NOT NULL PRIMARY KEY, tname TEXT, farbe TEXT, gewicht INTEGER);"
	_, err = db.Exec(sqlStmt)
	check(err)

	dat, err := ioutil.ReadFile("./csv/T.csv")
	check(err)
	//fmt.Print(string(dat))

	r := csv.NewReader(strings.NewReader(string(dat)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		//fmt.Println(reflect.TypeOf(record))
		//fmt.Println("RecordRow ", record)

		stmt, err := db.Prepare("INSERT INTO t(tnr) VALUES(?)")
		check(err)

		stmt.Exec(record[0])

	}







	/*
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		for i := 0; i < rand.Intn(99); i++ {
			_, err = stmt.Exec(i, fmt.Sprintf("Zahl%03d", i))
			if err != nil {
				log.Fatal(err)
			}
		}
		tx.Commit()

		/*
			rows, err := db.Query("select id, name from foo")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				var id int
				var name string
				err = rows.Scan(&id, &name)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(id, name)
			}
			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}

			stmt, err = db.Prepare("select name from foo where id = ?")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()
			var name string
			err = stmt.QueryRow("3").Scan(&name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(name)

			_, err = db.Exec("delete from foo")
			if err != nil {
				log.Fatal(err)
			}

			_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
			if err != nil {
				log.Fatal(err)
			}

			rows, err = db.Query("select id, name from foo")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				var id int
				var name string
				err = rows.Scan(&id, &name)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(id, name)
			}
			err = rows.Err()
			if err != nil {
				log.Fatal(err)
			}*/
}