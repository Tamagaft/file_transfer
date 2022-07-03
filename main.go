package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("choose mode")
	fmt.Println("1 - Send file")
	fmt.Println("2 - Get  file")

	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print(">")
	text, _ := consoleReader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if text == "1" {
		sendFile()
	} else if text == "2" {
		getFile()
	} else {
		fmt.Println(false)
	}
	os.Exit(0)
}

func sendFile() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Input file name\n>")
	fileName, _ := consoleReader.ReadString('\n')
	fileName = strings.Replace(fileName, "\n", "", -1)

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open file: %s \n", err)
	}
	defer f.Close()

	fmt.Print(">")
	dstIP, _ := consoleReader.ReadString('\n')
	dstIP = strings.Replace(dstIP, "\n", "", -1)

	conn, err := net.Dial("tcp", dstIP+":3333")
	if err != nil {
		log.Fatalf("Cannot connect to host: %s \n")
	}
	defer conn.Close()
	io.Copy(conn, f)

}

func getFile() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Input file name\n>")
	fileName, _ := consoleReader.ReadString('\n')
	fileName = strings.Replace(fileName, "\n", "", -1)

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Cannot create file %s \n", err)
	}
	defer f.Close()

	listener, err := net.Listen("tcp4", ":3333")
	if err != nil {
		log.Fatalf("Cannot start listener: %s \n", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Cannot accept connection: %s \n", err)
		}
		_, err = io.Copy(f, conn)
		if err != nil {
			log.Fatalf("Cannot write data to file: %s \n", err)
		}
		conn.Close()
	}

}
