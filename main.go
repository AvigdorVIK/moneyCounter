package main

import (
	"fmt"
)

var food int
var cloth int
var thinks int

func main() {

	fmt.Println("Введіть кошти витрачені на їжу, одяг та речі через ентер")
	fmt.Scan(&food)
	fmt.Scan(&cloth)
	fmt.Scan(&thinks)
	fmt.Println("Ви витратили ", food+cloth+thinks, "Гривень за сьогодні")

}
