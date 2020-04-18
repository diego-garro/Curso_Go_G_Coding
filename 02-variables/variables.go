package main

import "fmt"

func main() {
	var numero int
	numero = 25
	fmt.Println(numero)

	numero = 40
	fmt.Println(numero)

	nombre := "Samantha"
	fmt.Println(nombre)

	nombre, numero = "Diego", 100
	fmt.Println(nombre, numero)

	nombre2 := "Antonio"
	nombre, nombre2 = nombre2, nombre
	fmt.Println(nombre2, nombre)

	nombre3, nombre := "Melina", "Diego"
	fmt.Println(nombre, nombre3)

	var nombre4 string
	fmt.Println(nombre4)
}
