package main

import (
	"fmt"
	"microMutationPuntos/infraestructura"
)

func main() {
	if err := infraestructura.Run(); err != nil {
		fmt.Println("Error in stating server", err)
	}
}
