package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var file_name = "Matrice.txt"
	file, err := os.Open(file_name) // For read access.
	fmt.Println("\nOuverture du fichier", file_name, "...")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Réussie !!!")
	file.Close()
}
