package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// server
	ln, err := net.Listen("tcp", "localhost:9080")

	if err != nil {
		fmt.Printf("Error en la resolución de la dirección de la red!!!", err.Error())
		os.Exit(1)
	}

	defer ln.Close()

	con, erra := ln.Accept()

	if erra != nil {
		fmt.Printf("Error en la conexion!!!", erra.Error())
	}

	defer con.Close()

	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')

	// fmt.Println(msg)
	vav1 := msg
	fmt.Println(msg)

	if msg != vav1 {
		fmt.Println(msg)
		fmt.Fprintln(con, "Negativo")
	}

	if msg == "1.0, 1.0, 1" {
		fmt.Println(msg)
		fmt.Fprintln(con, "Positivo")
	}

	if msg == "0.0, 0.0, 0" {
		fmt.Println(msg)
		fmt.Fprintln(con, "Negativo")
	}

	if msg == "1.0, 1.0, 0" {
		fmt.Println(msg)
		fmt.Fprintln(con, "Positivo")
	}

}
