package main

import (
	//"fmt"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func submitHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из формы
	food := r.FormValue("food")
	cloth := r.FormValue("cloth")
	thinks := r.FormValue("thinks")

	db, err := sql.Open("mysql", "vik:000(localhost:3306)(127.0.0.1:3306)/moneyCount")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO moneyCount (foodInteger, clotherInteger, thinksInteger) VALUES (?, ?, ?)", food, cloth, thinks)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Данные успешно получены на сервере!")
}

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
	db, err := sql.Open("mysql", "vik:000(localhost:3306)/moneyCount")
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
		port = "5501"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/submit", submitHandler)
	http.ListenAndServe(":"+port, mux)
	dbconect()

}
