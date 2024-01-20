package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup // instanciation de notre structure WaitGroup

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
			matprod = matprod + strconv.Itoa(sum) + " "
		}
		matprod = matprod + "\n"
	}
	//fmt.Println(matprod)
}

func produit(ch chan string, tab []int, size int, i int) {
	defer wg.Done()
	matprod := ""
	for j := 0; j < size; j++ {

		sum := 0
		for k := 0; k < size; k++ {
			elem1 := tab[k+i*size]
			elem2 := tab[j+k*size]
			sum = sum + elem1*elem2
		}
		matprod = matprod + strconv.Itoa(sum) + " "
	}
	ch <- matprod
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
	fmt.Println("Délai en concurrentiel :", fin.Sub(debut))

	p := make(chan string)
	size := int(math.Sqrt(float64(len(mat))))
	prod := ""
	debut2 := time.Now()
	for i := 0; i < size; i++ {
		wg.Add(1) // ajoute 1 goroutine à attendre
		go produit(p, mat, size, i)
		prod = prod + "\n" + <-p
	}
	wg.Wait() // empêche l'exécution des lignes de code suivantes avant que toutes les goroutines se terminent
	fin2 := time.Now()
	fmt.Println("Délai en parallèle :", fin2.Sub(debut2), "\n")

	// Création d'un fichier .txt et écriture du
	f, err := os.Create("ProduitMat.txt")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(prod)
	defer f.Close()

	fmt.Println("Observez le résultat de matrice dans :", f.Name())
	fmt.Println("Est-il juste ?")
	var rep string
	for (rep != "OUI") || (rep != "NON") {
		fmt.Scanln(&rep)
		rep = strings.ToUpper(rep)
		if rep == "OUI" {
			os.Remove(f.Name())
			break
		} else if rep == "NON" {
			fmt.Println("AH SH*T, HERE WE GO AGAIN \n")
			os.Remove(f.Name())
			break
		} else {
			fmt.Println("Répondez par 'oui' ou par 'non'")
		}
	}

}
