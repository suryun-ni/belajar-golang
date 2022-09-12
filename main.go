package main

import (
	"fmt"
)

// masih awal, A.7-seterusnya
func main() {
	//variabel beserta type data
	// menuliskan type data nya
	var firstname string = "tengku"

	//tanpa type data
	lastname := "surya"
	//atau
	var lastNasme = "furqan"

	fmt.Println("hello world")
	//sama kayak println cuman %s digunakan untuk string
	fmt.Printf("hello %s %s %s!\n", firstname, lastname, lastNasme)
	// maksud diatas adalah seperti berikut
	fmt.Println("halo", firstname, lastname, lastNasme+"!")
}
