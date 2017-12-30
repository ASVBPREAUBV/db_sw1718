package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
)

func check(e error) {
	if e != nil {
		log.Printf("%q: %s\n", e)
		panic(e)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "./foo.db")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)



	//QUERIES
	var (
		pname string
	)

	//(a) Bestimme die Namen aller Projekte in Berlin.
	rows_a, err := db.Query("SELECT pname FROM p WHERE ort = ?", "Berlin")
	check(err)
	defer rows_a.Close()
	for rows_a.Next() {
		err := rows_a.Scan(&pname)
		check(err)
		fmt.Println("(a)", pname)
	}
	err = rows_a.Err()
	check(err)


	//(b) Bestimme fur jedes Projekt in Berlin die Namen aller gelieferten Teile.

	var (
		tnr string
		tname string
		farbe string
		gewicht int
	)

	var query_b = "SELECT * " +
		"FROM t "
	//+"JOIN ltp ON t.tnr=ltp.tnr"

	rows_b, err := db.Query(query_b, "Berlin")
	check(err)
	defer rows_b.Close()
	for rows_b.Next() {
		err := rows_b.Scan(&tnr, &tname, &farbe, &gewicht)
		check(err)
		fmt.Println("(b)", tnr, tname, farbe, gewicht)
	}
	err = rows_b.Err()
	check(err)

	//(c) Finde die Namen und Nummern aller Teile, die Lieferant Schulz liefert.
	//(d) Bestimme die Namen aller Lieferanten, die von Meschede nach Wetter liefern.
	//(e) Bestimme die Nummern und Orte aller Projekte, zu denen ein rotes Teil mit einem Gewicht von mehr als 5 geliefert wird.


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