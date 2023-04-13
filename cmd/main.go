package main

import (
	//"fmt"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//var foodInteger int
//var clothInteger int
//var thinksInteger int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../web/template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title         string
		FoodInteger   int
		ClothInteger  int
		ThinksInteger int
	}{
		Title:         "storinka",
		FoodInteger:   5,
		ClothInteger:  8,
		ThinksInteger: 7,
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func dbconect() {
	db, err := sql.Open("mysql", "viktor:***(localhost:5500)/moneyCounter")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, foodInteger, clotherInteger, thinksInteger FROM moneyCount")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var foodInteger int
		var clotherInteger int
		var thinksInteger int
		err = rows.Scan(&id, &foodInteger, &clotherInteger, &thinksInteger)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, foodInteger, clotherInteger, thinksInteger)
	}

}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5500"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

	//fmt.Println("Введіть кошти витрачені на їжу, одяг та речі через ентер")
	//fmt.Scan(&foodInteger)
	//fmt.Scan(&clothInteger)
	//fmt.Scan(&thinksInteger)
	//fmt.Println("Ви витратили ", foodInteger+clothInteger+thinksInteger, "Гривень за сьогодні")
	dbconect()

}
