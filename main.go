package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func rollHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_list.html")
		if err != nil {
			log.Fatal(err)
		}
		cars, err := dbGetCars()
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, cars)
	}
}

func getCarsByBrandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_form_brand.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)
	} else {
		t, err := template.ParseFiles("simple_form_brand.html")
		r.ParseForm()
		brand := r.Form.Get("brand")
		cars, err := dbGetCarByBrand(brand)
		//cars, err := dbGetCars()
		println(brand)
		if err != nil {
			log.Fatal(err)
		}
		err = t.ExecuteTemplate(w, "simple_form_brand.html", cars)
		fmt.Print(err)
		//fmt.Print(error)
	}
}

func addCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("simple_form.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, nil)

	} else {
		r.ParseForm()
		brand := r.Form.Get("brand")
		year := r.Form.Get("year")
		country := r.Form.Get("country")
		price := r.Form.Get("price")
		err := dbAddCar(brand, country, year, price)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println(port)
	}
	return ":" + port
}

func main() {
	err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", rollHandler)
	http.HandleFunc("/add", addCarHandler)
	http.HandleFunc("/brand", getCarsByBrandHandler)
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}
