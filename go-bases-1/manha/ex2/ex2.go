package main

import "fmt"

func main() {
	var (
		temp    float32 //Celsius
		umid    float32 //Porcentagem
		pressao float32 //hPa
	)

	temp, umid, pressao = 19.8, 10.4, 1013.25
	fmt.Println("Temperatura ", temp, "ºC\nUmidade ", umid, "%\nPressão ", pressao, " hPa")
}
