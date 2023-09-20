package main

import (
	"fmt"

	"github.com/reddevil31/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvertiraTexto(1456)

	fmt.Println(texto)
	fmt.Println(estado)

}
