package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mrsheepuk/magicalroleapi/internal/magicalroleapi"
)

func main() {
	port := 8080
	if len(os.Args[1:]) > 0 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil {
			port = n
		} else {
			fmt.Printf("Invalid port provided: %v", os.Args[1])
			os.Exit(1)
		}
	}
	handler := magicalroleapi.NewHandler()
	handler.Run(port)
}
