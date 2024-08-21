package main

import (
	"fmt"
	"os"
)


func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("Hostname: %s \n", hostname)
}
