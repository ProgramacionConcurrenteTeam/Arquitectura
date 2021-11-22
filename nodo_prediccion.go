package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Rol cliente
	con, _ := net.Dial("tcp", "localhost:9080")

	defer con.Close()
	fmt.Fprintln(con, "0.0, 0.0, 1")

	//leer respuesta
	buffIn := bufio.NewReader(con)
	msg, _ := buffIn.ReadString('\n')

	fmt.Println(msg)
}
