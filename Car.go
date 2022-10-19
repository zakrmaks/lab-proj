package main

import (
	"database/sql"
	"fmt"
)

type Car struct {
	Brand    string
	Country  string
	Price    int
	YearProd int
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgress"
	DB_NAME     = "postgres"
)

func dbConnect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME))
	if err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS cars (brand text,country text,price smallint, year_prod integer)"); err != nil {
		return err
	}
	return nil
}
func dbAddCar(brand, country, price, year string) error {
	sqlstmt := "INSERT INTO cars VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(sqlstmt, brand, country, price, year)
	if err != nil {
		return err
	}
	return nil
}
func dbGetCars() ([]Car, error) {
	var cars []Car
	stmt, err := db.Prepare("SELECT brand, country, price, year_prod FROM cars")
	if err != nil {
		return cars, err
	}
	res, err := stmt.Query()
	if err != nil {
		return cars, err
	}
	var tempCar Car
	for res.Next() {
		err = res.Scan(&tempCar.Brand, &tempCar.Country, &tempCar.Price, &tempCar.YearProd)
		if err != nil {
			return cars, err
		}
		cars = append(cars, tempCar)
	}
	return cars, err
}
func dbGetCarByBrand(brand string) (car []Car, error error) {
	var cars []Car
	fmt.Println(brand)
	//sqlstmt := "SELECT * FROM cars WHERE brand = ?"
	//stmt, err := db.Prepare(sqlstmt)
	//if err != nil {
	//	return cars, err, brand
	//}
	res, err := db.Query("SELECT * FROM cars WHERE brand = $1", brand)
	if err != nil {
		return cars, err
	}
	var tempCar Car
	for res.Next() {
		err = res.Scan(&tempCar.Brand, &tempCar.Country, &tempCar.Price, &tempCar.YearProd)
		if err != nil {
			return cars, err
		}
		cars = append(cars, tempCar)
	}
	return cars, err
}
