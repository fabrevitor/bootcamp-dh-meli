package main

import (
	"fmt"
	"time"
)

func main() {
	var meses = map[int]string{
		1:  "Jan",
		2:  "Fev",
		3:  "Mar",
		4:  "Apr",
		5:  "May",
		6:  "Jun",
		7:  "Jul",
		8:  "Aug",
		9:  "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dec",
	}

	fmt.Println(meses[1])

	fmt.Println(time.Now().Month())
}
