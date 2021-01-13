package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
				name STRING,
				age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	/* INSERT*/
	//cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	//_, err = DbConnection.Exec(cmd, "Nancy", 20)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	/* UPDATE*/
	//cmd = "UPDATE person SET age = ? WHERE name = ?"
	//_, err = DbConnection.Exec(cmd, 25, "Mike")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	/* Multiple Select*/
	//cmd = "SELECT * FROM PERSON"
	//rows, _ := DbConnection.Query(cmd)
	//defer rows.Close()
	//var pp []Person
	//for rows.Next() {
	//	var p Person
	//	err := rows.Scan(&p.Name, &p.Age)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	pp = append(pp, p)
	//}
	//for _, p := range pp {
	//	fmt.Println(p.Name, p.Age)
	//}

	/*Single Select*/
	//cmd = "SELECT * FROM person where age = ?"
	//row := DbConnection.QueryRow(cmd, 20)
	//var p Person
	//err = row.Scan(&p.Name, &p.Age)
	//if err != nil {
	//	log.Println("No row")
	//} else {
	//	log.Println(err)
	//}
	//fmt.Println(p.Name, p.Age)

	/*DELETE*/
	//cmd = "DELETE FROM person WHERE name = ?"
	//_, err = DbConnection.Exec(cmd, "Nancy")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	/*SQL Injection*/
	//tableName := "person"
	tableName := "person; INSERT INTO person (name, age) VALUES ('Mr.X', 100)" //SQLが実行されてしまう
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName) //tableNameは?による代入ができない, ?は通常文字列にエスケープされる
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person //Personの構造体のみ入る配列を作成
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}