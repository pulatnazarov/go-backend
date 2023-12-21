package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-backend/postgres/inventory"
)

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	inv := inventory.New(db)

	//c := country.Country{
	//	ID:       uuid.New(),
	//	Name:     "China",
	//	Currency: "Yang",
	//}
	//
	//if err = inv.InsertCountry(c); err != nil {
	//	panic(err)
	//}

	countries, err := inv.GetAllCountries()
	if err != nil {
		panic(err)
	}

	fmt.Println("all the countries: ", countries)
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=person password=12345 database=country sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}
