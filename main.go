package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

func main() {
	citienames := os.Args[1:]

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	fmt.Println("Connected!")

	for _, name := range citienames {
		city := City{}
		if err := db.Get(&city, "SELECT * FROM city WHERE city.Name = ?", name); err != nil {
			log.Fatalf("Cannot find city %s: %s", name, err)
		} else {
			fmt.Printf("%sの人口は%d人です\n", city.Name, city.Population)
		}
	}
	if _, err := db.Exec("INSERT INTO city (Name, CountryCode, District, Population) VALUES (?, ?, ?, ?)",
		"oookayama", "JPN", "Tokyo", 2147483647); err != nil {
		log.Fatalf("Cannot insert city ooo: %s", err)
	}
	city := City{}
	if err := db.Get(&city, "SELECT * FROM city ORDER BY id DESC LIMIT 1;"); err != nil {
		log.Fatalf("Cannot find city ooo: %s", err)
	} else {
		fmt.Printf("%sの人口は%d人です\n", city.Name, city.Population)
	}
}
