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

type WorkItem struct {
	Row        int
	MatrixLine []int
}

type ResultItem struct {
	Row    int
	Result string
}

var mat []int
var size int
var wg sync.WaitGroup

func worker(id int, jobs <-chan WorkItem, results chan<- ResultItem) {
	for job := range jobs {
		//fmt.Printf("Worker %d processing job at row %d\n", id, job.Row)
		result := produit(mat, size, job.Row)
		results <- ResultItem{Row: job.Row, Result: result}
	}
}

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

func produit(tab []int, size int, i int) string {
	matprod := ""
	// fmt.Println("start go routine ",i)
	for j := 0; j < size; j++ {

		sum := 0
		for k := 0; k < size; k++ {
			elem1 := tab[k+i*size]
			elem2 := tab[j+k*size]
			sum = sum + elem1*elem2
		}
		matprod = matprod + strconv.Itoa(sum) + " "
	}
	// fmt.Println("finished go routine ", i)
	return matprod
}

func main() {

	// Ouvrir le fichier en lecture seulement
	file, err := os.Open("matrice500.txt")
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
	mat = String_To_Int(data)
	fmt.Println("On effectue maintenant le produit d'une matrice de longueur", math.Sqrt(float64(len(mat))), "avec elle-même.")

	debut := time.Now()
	Produit(mat)
	fin := time.Now()
	fmt.Println("Délai en concurrentiel :", fin.Sub(debut))
	size = int(math.Sqrt(float64(len(mat))))
	debut2 := time.Now()
	var numJobs = size
	const numWorkers = 7

	// Create channels for jobs and results
	jobs := make(chan WorkItem, numJobs)
	results := make(chan ResultItem, numJobs)

	// Create worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, jobs, results)
		}(i)
	}

	// Decompose matrix in lines and add them to the jobs channel
	//fmt.Println("Size of mat: ",len(mat),"\n",mat[size*0:size*(0+1)],"\n",mat[size*(size):size*(2999+1)])

	for i := 0; i < size; i++ {
		//fmt.Println("Matrice line ",i,"=",mat[size*i:size*(i+1)])
		if mat[size*i:size*(i+1)] == nil {
			fmt.Println(i)
		}
		jobs <- WorkItem{Row: i, MatrixLine: mat[size*i : size*(i+1)]}
	}

	// Close the jobs channel to signal that no more jobs will be added
	close(jobs)

	// Wait for all workers to finish processing
	wg.Wait()

	// Close the results channel since all workers are done
	close(results)

	// Collect and print the results in the correct order
	resultMatrix := make([]string, numJobs)
	for i := 0; i < numJobs; i++ {
		result := <-results
		resultMatrix[result.Row] = result.Result
	}

	finalRes := ""

	for i := 0; i < len(resultMatrix); i++ {
		//fmt.Println(i)
		if resultMatrix[i] == "" {
			print(i)
		} else {
			finalRes += resultMatrix[i] + "\n"
		}
	} // empêche l'exécution des lignes de code suivantes avant que toutes les goroutines se terminent
	fin2 := time.Now()
	fmt.Println("Délai en parallèle :", fin2.Sub(debut2), "\n")

	// Création d'un fichier .txt et écriture du
	f, err := os.Create("ProduitMat.txt")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(finalRes)
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
