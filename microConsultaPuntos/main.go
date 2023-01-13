package main

import (
	"fmt"
	"microConsultaPuntos/infraestructura"
)

func main() {
	if err := infraestructura.Run(); err != nil {
		fmt.Println("Error in stating server", err)
	}
}
