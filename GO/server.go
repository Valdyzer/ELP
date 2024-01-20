package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Server struct {
	listenAddress string
	ln            net.Listener
	quitch        chan struct{}
}

var wg sync.WaitGroup // instantiation of our WaitGroup structure

func CreateServer(listenAddress string) *Server {
	serverInstance := &Server{
		listenAddress: listenAddress,
		quitch:        make(chan struct{}),
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

	<-s.quitch

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

	for {
		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error printing data: ", err)
			continue
		}
		tempMatString := ""
		matriceString := ""
		msg := buffer[:n]
		if len(msg) > 0 {
			globalBuffer = append(globalBuffer, msg...)
		}

		// Check if we received the entire matrix; to signal the end of a matrix, we append the letter 'z' at the end
		if globalBuffer[len(globalBuffer)-1] == 'z' {
			timerStart := time.Now()
			// Remove the term signaling the end of the matrix and start processing
			globalBuffer = globalBuffer[:len(globalBuffer)-1]
			fmt.Println("GlobalBuffer: " + string(globalBuffer[:8]) + "..." + string(globalBuffer[(len(globalBuffer)-9):]) + "\n")

			tempMatString = string(globalBuffer)
			multMatrice := ParseString(tempMatString)
			mat := String_To_Int(multMatrice)
			size := int(math.Sqrt(float64(len(mat))))
			p := make(chan string)

			for i := 0; i < size; i++ {
				wg.Add(1) // add 1 goroutine to wait
				go produit(p, mat, size, i)
				matriceString = matriceString + "\n" + <-p
			}

			wg.Wait()
			if matriceString[:1] == "\n" {
				matriceString = matriceString[1:]
			}
			timerEnd := time.Now()
			matriceString += "\nCalculation executed by the server in: " + timerEnd.Sub(timerStart).String() + "z"
			fmt.Println("Message sent back: ", string(matriceString[:16])+"..."+string(matriceString[(len(matriceString)-51):])+"\n")
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
	server := CreateServer(":3000")
	log.Fatal(server.Start())
}
