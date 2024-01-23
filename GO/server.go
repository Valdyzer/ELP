package main

import (
	"fmt"
	"net"
	"log"
	"strconv"
	"strings"
	"math"
	"sync"
	"time"
)

type Server struct {
	listenAddress string 
	ln 			  net.Listener 
	quitch 		  chan struct{}
}


type WorkItem struct {
	Row   int
	MatrixLine []int
}

type ResultItem struct {
	Row    int
	Result string
}


var wg sync.WaitGroup


func worker(id int, jobs <-chan WorkItem, results chan<- ResultItem, size int, mat []int) {
	for job := range jobs {
		//fmt.Printf("Worker %d processing job at row %d\n", id, job.Row)
		result := produit(mat,size,job.Row)
		results <- ResultItem{Row: job.Row, Result: result}
	}
}


func CreateServer(listenAddress string ) *Server {
	serverInstance := &Server {
		listenAddress : listenAddress,
		quitch : make(chan struct{}),
	}
	return serverInstance
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	
	go s.acceptConnection()

	<- s.quitch
	
	return nil
}

func (s *Server) acceptConnection() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error while trying to accept connection: ", err)
			continue
		} else {
			fmt.Println("\nConnection established with client. ")
		}
		
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	var globalBuffer []byte
	var mat []int
	var size int
	
	for {
		buffer := make([]byte, 2048)
		n,err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error printing data: ",err)
			continue
		}
		tempMatString := ""
		matriceString := ""
		msg := buffer[:n]
		if len(msg) > 0 {
			globalBuffer = append(globalBuffer, msg...)
		}
		
		//on teste si on a reçu la totalité de la matrice, pour signaler la fin d-une matrice on met la lettre z à la fin
		if 	globalBuffer[len(globalBuffer)-1] =='z'{

			timerStart := time.Now()
			//on retire le terme qui signale la fin de la matrice puis on commence le traitement
			globalBuffer = globalBuffer[:len(globalBuffer)-1] 
			fmt.Println("GlobalBuffer: "+ string(globalBuffer[:8]) + "..." + string(globalBuffer[(len(globalBuffer)-9):])+"\n")

			tempMatString = string(globalBuffer)
			multMatrice := ParseString(tempMatString)
			mat = String_To_Int(multMatrice)
			size = int(math.Sqrt(float64(len(mat))))



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
					worker(workerID, jobs, results, size, mat)
				}(i)
			}
		
			// Decompose matrix in lines and add them to the jobs channel
			//fmt.Println("Size of mat: ",len(mat),"\n",mat[size*0:size*(0+1)],"\n",mat[size*(size):size*(2999+1)])
		
			for i := 0; i < size; i++ {
				//fmt.Println("Matrice line ",i,"=",mat[size*i:size*(i+1)])
				jobs <- WorkItem{Row: i, MatrixLine: mat[size*i:size*(i+1)]}
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
				if  resultMatrix[i] == "" {	
					print(i)
				}else{
					finalRes += resultMatrix[i] + "\n"
				}
			}

			matriceString = finalRes
			timerEnd := time.Now()
			matriceString += "\nCalcul exécuté par le serveur en : " + timerEnd.Sub(timerStart).String() + "z"
			fmt.Println("Message sent back: ", string(matriceString[:16]) + "..." + string(matriceString[(len(matriceString)-49):])+"\n")
			_, err = conn.Write([]byte(matriceString))
			conn.Close()
			globalBuffer = nil
			return
		}
	}
}

func ParseString(CONTENT string) []string {
	dataByLine := strings.Split(CONTENT, "\n")
	var line []string
	var res []string

	for i := 0; i < len(dataByLine); i++ {
		line = strings.Split(dataByLine[i], " ")
		for j := 0; j < len(line); j++ {
			if line[j] != "" {	
				res = append(res, line[j])
			}
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

func produit(tab []int, size int, i int) string{
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
	server := CreateServer(":3000")
	log.Fatal(server.Start())
}
