package main

import (
	//"fmt"
	"net/http"
	"os"
)

//var foodInteger int
//var clothInteger int
//var thinksInteger int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

	//fmt.Println("Введіть кошти витрачені на їжу, одяг та речі через ентер")
	//fmt.Scan(&foodInteger)
	//fmt.Scan(&clothInteger)
	//fmt.Scan(&thinksInteger)
	//fmt.Println("Ви витратили ", foodInteger+clothInteger+thinksInteger, "Гривень за сьогодні")

}
