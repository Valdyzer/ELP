package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Byte_To_String(CONTENT []byte) []string {

	dataByLine := strings.Split(string(CONTENT), "\n")
	var line []string
	var res []string

	for i := 0; i < len(dataByLine); i++ {
		line = strings.Split(dataByLine[i], " ")
		for j := 0; j < len(line); j++ {
			res = append(res, line[j])
		}
	}
	return res
}

func String_To_Int(DATA []string) []int {
	var tab []int
	for i := 0; i < len(DATA); i++ {
		n, err := strconv.Atoi(DATA[i])
		if err != nil {
			log.Fatal(err)
		}
		tab = append(tab, n)
	}
	return tab
}

func main() {

	file_name := "Matrice.txt"

	file, err := os.Open(file_name) // Permet l'accès au fichier
	fmt.Println("\nOuverture du fichier", file_name, "...")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Réussie !!!")

	content, err := os.ReadFile(file_name) // Lecture du fichier...
	if err != nil {
		log.Fatal(err)
	}

	data := Byte_To_String(content)
	mat := String_To_Int(data)

	fmt.Println(mat[0] + mat[4])

	file.Close()
}
