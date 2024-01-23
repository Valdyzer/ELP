package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var globalBuffer []byte
var matrixSize = []string{"1000", "1500", "2000", "3000"}
var srcFile int
var connectFlag bool
var pathToFile string

func readLoop(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Connection terminated by server.")
				connectFlag = false
			} else {
				fmt.Println("Error printing data:", err)
			}
			return
		}

		msg := buffer[:n]

		if len(msg) > 0 {
			globalBuffer = append(globalBuffer, msg...)
		}

		if globalBuffer[len(globalBuffer)-1] == 'z' {
			globalBuffer = globalBuffer[:len(globalBuffer)-1]
			fileTime := strconv.FormatInt(time.Now().Unix(), 10)
			outputFile := "ProduitMat" + matrixSize[srcFile-1] + "x" + matrixSize[srcFile-1] + fileTime + ".txt"
			f, err := os.Create(pathToFile + outputFile)
			if err != nil {
				fmt.Println("Error creating file: ", err)
			}

			fmt.Println("Received result matrix: " + string(globalBuffer[:16]) + "..." + string(globalBuffer[(len(globalBuffer)-47):]))
			f.WriteString(string(globalBuffer))
			fmt.Println("Output in file ", outputFile)
			globalBuffer = nil
		}
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

func startup() int {
	var userInput int
	fmt.Println("What matrix size do you want to send to the server?\n1. 1000x1000\n2. 1500x1500\n3. 2000x2000\n4. 3000x3000 (not recommended)\nPlease enter the number of your choice.\n")

	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return -1
	}

	if userInput >= 1 && userInput <= 4 {
		fmt.Println("Starting...")
	} else {
		fmt.Println("Incorrect format. Please enter a single-digit number between 1 and 4 corresponding to your choice.\n")
		fmt.Scanln(&userInput)
	}

	return userInput
}

func main() {
	connectFlag = true
	pathToFile = "../GO" //to run, uncomment this line and replace /pathtomatrixfiles with the path to the directory containing the matrix files

	for connectFlag {
		srcFile = startup()

		if srcFile == -1 {
			fmt.Println("Startup function returned an error. Let's try again.")
			continue
		}

		file, err := os.Open("matrice" + matrixSize[srcFile-1] + ".txt")
		fmt.Println("\nOuverture du fichier", file.Name(), "...")

		if err != nil {
			fmt.Println("Erreur lors de l'ouverture du fichier :", err)
			return
		}
		fmt.Println("Réussie !\n")

		// Lire le fichier et retourner le contenu en []byte
		content, err := os.ReadFile("matrice500.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			fmt.Println("Error connecting to server: ", err)
		}

		defer conn.Close()
		fmt.Println("Connected to server !")

		_, err = conn.Write(content)
		if err != nil {
			fmt.Println("Erreur d'envoi:", err)
			return
		}
		readLoop(conn)
	}
}
