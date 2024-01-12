package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

func Produit(tab []int) {
	var matprod string
	var sum int
	size := int(math.Sqrt(float64(len(tab))))
	for i := 0; i < size; i++ {

		for j := 0; j < size; j++ {

			sum = 0
			for k := 0; k < size; k++ {
				elem1 := tab[k+i*size]
				elem2 := tab[j+k*size]
				sum = sum + elem1*elem2
			}
			matprod = matprod + strconv.Itoa(sum) + "\t"
		}
		matprod = matprod + "\n\n"
	}
	//fmt.Println(matprod)
}

func produit(tab []int, size int, i int) {

	matprod := ""
	for j := 0; j < size; j++ {

		sum := 0
		for k := 0; k < size; k++ {
			elem1 := tab[k+i*size]
			elem2 := tab[j+k*size]
			sum = sum + elem1*elem2
		}
		matprod = matprod + strconv.Itoa(sum) + "\t"
	}
	matprod = matprod + "\n\n"
	//fmt.Println(matprod)
}

func main() {
	// Ouvrir le fichier en lecture seulement
	file, err := os.Open("Matrice.txt")
	fmt.Println("\nOuverture du fichier", file.Name(), "...")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	fmt.Println("Réussie !\n")

	// Lire le fichier et retourner le contenu en []byte
	content, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := Byte_To_String(content)
	mat := String_To_Int(data)
	fmt.Println("On effectue maintenant le produit d'une matrice de longueur", math.Sqrt(float64(len(mat))), "avec elle-même.")

	debut := time.Now()
	Produit(mat)
	fin := time.Now()
	fmt.Println(fin.Sub(debut))

	//p := ""
	size := int(math.Sqrt(float64(len(mat))))
	debut2 := time.Now()
	for i := 0; i < size; i++ {
		go produit(mat, size, i)
	}
	fin2 := time.Now()
	fmt.Println(fin2.Sub(debut2))
}
