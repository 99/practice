package main

import (
	"fmt"
	"practice/port_scanner/port"
)

func main() {
	fmt.Println("Building Port Scanner in Go")
	open := port.ScanPort("tcp", "localhost", 65088 )
	fmt.Printf("Port Open: %t\n", open)

	results := port.InitialScan("localhost")
	fmt.Println(results)
}